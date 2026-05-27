char read_char() {
    char c[1];
    do_syscall(0, 0, c, 1);

    return c[0];
}

int main() {
    char msg[3] = {'Y','E','S'};
    char read = read_char();

    if (read == 'x') {
        print(msg, 3);
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

