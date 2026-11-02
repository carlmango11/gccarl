typedef char * va_list;
void put_char(char c) {
  write(1, & c, 1);
}
void put_string(const char * s) {
  if (s == 0) {
    s = "(null)";
  }
  while ( * s) {
    put_char( * s++);
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
    put_unsigned(0 - (unsigned int) n, 10);
    return;
  }
  put_unsigned((unsigned int) n, 10);
}
void my_printf(const char * format, ...) {
  va_list args;
  va_start(args, format);
  while ( * format) {
    if ( * format != '%') {
      put_char( * format++);
      continue;
    }
    format++;
    switch ( * format) {
    case 'd':
    case 'i':
      put_int(( * (int * )((args + = sizeof(int)) - sizeof(int))));
      break;
    case 'u':
      put_unsigned(( * (unsigned int * )((args + = sizeof(unsigned int)) - sizeof(unsigned int))), 10);
      break;
    case 'x':
      put_unsigned(( * (unsigned int * )((args + = sizeof(unsigned int)) - sizeof(unsigned int))), 16);
      break;
    case 'c':
      put_char((char)( * (int * )((args + = sizeof(int)) - sizeof(int))));
      break;
    case 's':
      put_string(( * (const char * * )((args + = sizeof(const char * )) - sizeof(const char * ))));
      break;
    case '%':
      put_char('%');
      break;
    case '\0':
      va_end(args);
      return;
    default:
      put_char('%');
      put_char( * format);
      break;
    }
    format++;
  }
  va_end(args);
}
int main(void) {
  my_printf("Hello %s\n", "world");
}