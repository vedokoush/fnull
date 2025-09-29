#include <bits/stdc++.h>

using namespace std;

int main() {
    // khai bao
    int n;
    int a[n];
    int s = 0;

    // nhap du lieu
    cin >> n;
    for (int i = 1; i <= n; i++) {
        // cout << i << ' ';

        cin >> a[i];
        s = s + a[i];

    }

    for (int i = 1; i <= n; ++i) {
        cout << a[i] << ' ';
    }
    cout << endl << s;




    
}

/*

{3, 4, _, _, 2}
 1  2  3  4  5


*/