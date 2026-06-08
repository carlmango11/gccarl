int main() {
    char msg[2] = {'Y', 'E'};
    print(msg, 2);
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

