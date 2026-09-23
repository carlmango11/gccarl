void main() {
    int res = 0;

    switch(5) {
    case 2:
        res = res + 2;
    case 7:
        res = res + 7;
    case 5:
        res = res + 5;
    case 3:
        res = res + 3;
    }

    assert(8, res);
}