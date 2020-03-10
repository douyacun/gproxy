main:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0  go build -ldflags '-extldflags "-static"'  -a -o gproxy main.go
