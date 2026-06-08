void read_str(char buf[]) {
    do_syscall(0, 0, buf, 20);
}

int main() {
    char msg[] = "hello";

    char c[20];
    read_str(c);

    print(msg, 5);
    print(c, 5);
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

