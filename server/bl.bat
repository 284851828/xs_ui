swag init
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o xs -tags="sonic avx" main.go 

