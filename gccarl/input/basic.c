int main() {
    char c = 'Y';
    char *p = &c;
    c = *p;
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

