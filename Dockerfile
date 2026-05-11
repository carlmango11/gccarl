FROM --platform=linux/amd64 debian:latest

WORKDIR /app

COPY build/basic /app/basic
RUN chmod +x /app/basic

CMD ["./basic"]

