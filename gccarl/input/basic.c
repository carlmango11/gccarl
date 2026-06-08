void read_str(char buf[]) {
    do_syscall(0, 0, buf, 20);
}

int main() {
    char msg[] = "hello ";

    char c[20];
    read_str(c);

    int i = 0;

    while (i < 3) {
        print(msg, 6);
        print(c, 20);
        i = i + 1;
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

