default:
	GOOS=js GOARCH=wasm go build -o public/lib.wasm app/main.go