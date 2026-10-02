void main() {
    int res = 0;

    if (1 != 1) {
        res = 1;
    }

    assert(0, res);

    if (1 != 2) {
        res = 2;
    }

    assert(2, res);
}