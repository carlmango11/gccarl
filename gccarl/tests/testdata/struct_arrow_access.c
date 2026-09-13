struct Person {
    int age;
}

int main() {
    struct Person p = {12};

    struct Person *y;
    y = &p;

    assert(12, y->age);
}