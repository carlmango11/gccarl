int main() {
    int x = 4;
    int *p = &x;
    int addr = (int)p;

    assert(addr, p);

    p = p + 2;

    assert(addr + 8, p);
}