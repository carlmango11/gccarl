int main() {
    char msg[2] = {'Y', 'E'};
    int res = 0;

    if (msg[1] == 'E') {
        res = 1;
    }

    assert(1, res);
}