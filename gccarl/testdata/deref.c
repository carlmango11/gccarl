int main() {
    char c = 'Y';
    char *p = &c;
    c = *p;

    test_set(1, c);
}