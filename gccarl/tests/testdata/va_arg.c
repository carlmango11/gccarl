#define va_arg(ap, Type) (*(Type *)((ap += sizeof(Type)) - sizeof(Type)))

typedef char *va_list;

void main() {
    va_list args;
    int x = va_arg(args, int);
}