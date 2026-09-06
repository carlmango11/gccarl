struct Point {
    int x;
    int y;
}

int main() {
    struct Point p = {1, 2};

    assert(1, p.x);
    assert(2, p.y);
}