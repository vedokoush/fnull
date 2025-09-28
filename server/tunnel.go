package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var (
	relayRegisterURL = ""
	wsConn           *websocket.Conn
)

type reqMsg struct {
	Type    string              `json:"type"`
	ID      string              `json:"id"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
}

type respMsg struct {
	Type    string              `json:"type"`
	ID      string              `json:"id"`
	Status  int                 `json:"status"`
	Headers map[string][]string `json:"headers"`
}

type endMsg struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func startTunnel() {
	if relayRegisterURL == "" {
		log.Println("relayRegisterURL empty: not starting relay tunnel")
		return
	}

	u := relayRegisterURL
	if _, err := url.Parse(u); err != nil {
		log.Printf("invalid relayRegisterURL: %v", err)
		return
	}
	if !hasQuery(u) {
		u = fmt.Sprintf("%s?token=%s", u, token)
	} else {
		u = fmt.Sprintf("%s&token=%s", u, token)
	}

	log.Printf("Connecting to relay %s ...", u)
	dialer := websocket.DefaultDialer
	conn, resp, err := dialer.Dial(u, nil)
	if err != nil {
		if resp != nil {
			body := new(bytes.Buffer)
			_, _ = io.Copy(body, resp.Body)
			log.Fatalf("dial error: %v, resp: %s", err, body.String())
		}
		log.Fatalf("dial error: %v", err)
	}
	wsConn = conn
	log.Printf("Connected to relay. token=%s", token)

	go func() {
		for {
			mt, b, err := wsConn.ReadMessage()
			if err != nil {
				log.Printf("relay read error: %v", err)
				return
			}
			if mt == websocket.TextMessage {
				var rm reqMsg
				if err := json.Unmarshal(b, &rm); err != nil {
					log.Printf("invalid request json: %v", err)
					continue
				}
				if rm.Type != "request" {
					continue
				}
				go handleRelayRequest(rm)
			}
		}
	}()
}

func hasQuery(s string) bool {
	return len(s) > 0 && (bytes.IndexByte([]byte(s), '?') >= 0)
}

func handleRelayRequest(rm reqMsg) {
	id := rm.ID
	localURL := fmt.Sprintf("http://localhost:8000%s", rm.Path)
	req, err := http.NewRequest(rm.Method, localURL, nil)
	if err != nil {
		sendErrorResponse(id, 500, map[string][]string{"Content-Type": {"text/plain"}})
		sendEnd(id)
		return
	}
	for k, v := range rm.Headers {
		req.Header[k] = v
	}

	client := &http.Client{
		Timeout: 0,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("local request failed: %v", err)
		sendErrorResponse(id, 502, map[string][]string{"Content-Type": {"text/plain"}})
		sendEnd(id)
		return
	}
	defer resp.Body.Close()

	rh := respMsg{
		Type:    "response",
		ID:      id,
		Status:  resp.StatusCode,
		Headers: map[string][]string{},
	}
	for k, v := range resp.Header {
		rh.Headers[k] = v
	}
	wsConn.SetWriteDeadline(time.Now().Add(15 * time.Second))
	if err := wsConn.WriteJSON(rh); err != nil {
		log.Printf("write response header to relay failed: %v", err)
		return
	}

	buf := make([]byte, 32*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			wsConn.SetWriteDeadline(time.Now().Add(30 * time.Second))
			if err := wsConn.WriteMessage(websocket.BinaryMessage, buf[:n]); err != nil {
				log.Printf("write binary to relay failed: %v", err)
				return
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read local body error: %v", err)
			break
		}
	}
	sendEnd(id)
}

func sendErrorResponse(id string, status int, headers map[string][]string) {
	rh := respMsg{
		Type:    "response",
		ID:      id,
		Status:  status,
		Headers: headers,
	}
	wsConn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_ = wsConn.WriteJSON(rh)
}

func sendEnd(id string) {
	e := endMsg{Type: "end", ID: id}
	wsConn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	_ = wsConn.WriteJSON(e)
}
