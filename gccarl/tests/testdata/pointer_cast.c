void main() {
    char c = 'C';
    char *args = &c;

    int x = ( * (int  *)((args += sizeof(int)) - sizeof(int)));
}