int main() {
    char c = 'Y';
    char *p = &c;
    c = *p;

    assert('Y', c);
}