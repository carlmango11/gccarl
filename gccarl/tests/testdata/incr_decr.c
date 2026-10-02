void main() {
    int res = 0;
    res++;

    assert(1, res);

    res--;

    assert(0, res);
}