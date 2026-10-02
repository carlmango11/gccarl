void main() {
    int res = 0;

    if (1 <= 2) {
        res = 1;
    }

    assert(1, res);

    if (2 <= 2) {
        res = 2;
    }

    assert(2, res);

    if (3 <= 2) {
        res = 3;
    }

    assert(2, res);
}