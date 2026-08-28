struct Person {
    char age;
}

int main() {
    struct Person p = {'X'};

   struct Person *y;

   y = &p;

    char m[] = {y->age};

    print(m, 1);
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

