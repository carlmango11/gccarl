struct Person {
    Inner i;
}

struct Inner {
    int y;
}

int main() {
    struct Person p = (Person){
        .i = 4,
        45,
    };
}

int print(char msg[], int len) {
    do_syscall(1, 1, msg, len);
}

