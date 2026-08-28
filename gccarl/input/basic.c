struct Person {
    char age;
}

int main() {
    struct Person p = {'X'};
    p.age = 'Y';
    char m[] = {p.age};

    print(m, 1);
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

