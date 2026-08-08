int main() {
    char c = 'Y';
    char *p = &c;
    c = *p;

    char msg[1] = {c};
   print(msg, 1);
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

