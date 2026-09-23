struct Point {
    int x;
    int y;
}

struct Coord {
    char name;
    struct Point point;
}

void main() {
    struct Coord c = {
        'X',
        { 2, 5 },
    };

    assert('X', c.name);
    assert(2, c.point.x);
    assert(5, c.point.y);
}