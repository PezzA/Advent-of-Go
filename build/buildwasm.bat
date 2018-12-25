set GOARCH=wasm
set GOOS=js
go build -o static\main.wasm 2018\Test\main.go

