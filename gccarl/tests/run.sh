export CGO_CFLAGS="-I$(brew --prefix keystone)/include"
export CGO_LDFLAGS="-L$(brew --prefix keystone)/lib"
go run .