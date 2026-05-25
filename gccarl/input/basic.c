
char read_char() {
    char c[1] = {' '};
    do_syscall(0, 0, c, 1);
}

int main() {
    char c[1] = {' '};
    do_syscall(0, 0, c, 1);

    char msg[3] = {'Y','E','S'};

    if (c[0] == 'x') {
        print(msg, 3);
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

