go run main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cmdb main.go
