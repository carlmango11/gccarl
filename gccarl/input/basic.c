struct Person {
    struct Address addr;
    int age;
}

struct Address {
    char street;
}

int main() {
    struct Person p = {
        .addr.street = 'X',
        17,
    };
}

