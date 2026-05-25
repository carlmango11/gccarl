int main() {
    int x = 3;
    char msg[1] = {'Y'};

    if (x == 3) {
        print(msg, 1);
    }
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}
