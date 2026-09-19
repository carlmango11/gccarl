void put_char(char c) {
    write(1, &c, 1);
}

void put_string(const char *s) {
    if (s == NULL) {
        s = "(null)";
    }

    while (*s) {
        put_char(*s++);
    }
}

void put_unsigned(unsigned int n, unsigned int base) {
    const char digits[] = "0123456789abcdef";

    if (n >= base) {
        put_unsigned(n / base, base);
    }

    put_char(digits[n % base]);
}

void put_int(int n) {
    if (n < 0) {
        put_char('-');

        // Avoid overflow for INT_MIN.
        put_unsigned(0u - (unsigned int)n, 10);
        return;
    }

    put_unsigned((unsigned int)n, 10);
}

void my_printf(const char *format, ...) {
    va_list args;
    va_start(args, format);

    while (*format) {
        if (*format != '%') {
            put_char(*format++);
            continue;
        }

        format++; // Skip '%'

        switch (*format) {
        case 'd':
        case 'i':
            put_int(va_arg(args, int));
            break;

        case 'u':
            put_unsigned(va_arg(args, unsigned int), 10);
            break;

        case 'x':
            put_unsigned(va_arg(args, unsigned int), 16);
            break;

        case 'c':
            put_char((char)va_arg(args, int));
            break;

        case 's':
            put_string(va_arg(args, const char *));
            break;

        case '%':
            put_char('%');
            break;

        case '\0':
            // Trailing '%' — stop rather than reading past the string.
            va_end(args);
            return;

        default:
            // Preserve unknown format specifiers visibly.
            put_char('%');
            put_char(*format);
            break;
        }

        format++;
    }

    va_end(args);
}

int main(void) {
    my_printf("Hello %s\n", "world");
    my_printf("n=%d, hex=%x, char=%c, percent=%%\n", -42, 255, 'A');
}