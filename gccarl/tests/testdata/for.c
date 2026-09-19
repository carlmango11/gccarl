int main() {
    int res = 0;

    for (int i = 0; i < 10; i = i + 1) {
        res = res + 1;
    }

    assert(10, res);
}