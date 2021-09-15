build:
	go build -o ./bin/scanport cmd/port-scanner/main.go
run: build
	./bin/scanport