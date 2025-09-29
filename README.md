# fnull

[![build passing](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/vedokoush/fnull/action)
[![license MIT](https://img.shields.io/badge/license-MIT-blue)](LICENSE)
[![docs](https://img.shields.io/badge/docs-online-lightgrey)](https://fnull.shouko.site)

Get things from one computer to another — simple and fast.

`fnull` is a small, focused file-transfer CLI and library written in Go. It lets you send and receive files and directories through a relay server using short links. Designed for quick ad-hoc transfers and easy automation.

## Design
- Small, single-binary CLI (no heavy deps)
- Short human-friendly transfer codes / links
- Relay server + client model (future: P2P / E2EE)
- Cross-platform: Linux, macOS, Windows

## Quickstart

### Send
On the sending machine:
```shellbash
fnull send path/to/file-or-folder
# -> prints a short link (or code) for the receiver
```

### Receive
On the receiving machine:
```shellbash
fnull receive <fnull-link-or-code>
# -> downloads file(s) to current directory
```

## Packaging & Distribution
Binaries are produced for Linux, macOS, and Windows.
Detailed packaging instructions (Homebrew, apt, scoop, winget, .deb/.rpm, etc.) and signed releases will be published on the project website:

https://... (installation & package guide)

## License
MIT — see the LICENSE file.
