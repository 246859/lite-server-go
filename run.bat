go build -o ./bin/run/run.exe -gcflags "all=-N -l" github.com/246859/lite-server-go
start ./bin/run/run.exe