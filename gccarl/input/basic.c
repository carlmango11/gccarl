char read_char() {
    char c[1];
    do_syscall(0, 0, c, 1);

    return c[0];
}

int main() {
    char ans = read_char();

    if (ans == '1') {
        print("was 1", 5);
    } else if (ans == '2'){
        print("was 2", 5);
    } else {
        print("neither", 5);
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

