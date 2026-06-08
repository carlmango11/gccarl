char read_char() {
    char c[1];
    do_syscall(0, 0, c, 1);

    return c[0];
}

int main() {
    char msg[2] = {'Y', 'E'};
    char ans = read_char();

    if (ans == 'X') {
        print(msg, 2);
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

