#include <iostream>
using namespace std;

int main() {
    int N;
    cin >> N;
    int HFirst, HCurrent;
    cin >> HFirst;
    for (uint i = 1; i < N; i++) {
        cin >> HCurrent;
        if (HCurrent > HFirst) {
            cout << i + 1 << endl;
            return 0;
        }
    }
    cout << -1 << endl;

    return 0;
}
