# brew install FiloSottile/musl-cross/musl-cross first
build-mac:
	go build -o bin/tx-decoder
build-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" go build -o bin/tx-decoder_linux
build-windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CGO_LDFLAGS="-static" go build -o bin/tx-decoder_windows
build-all: build-mac build-linux build-windows
