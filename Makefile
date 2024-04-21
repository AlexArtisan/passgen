build:
	go build -o passgen cmd/passgen.go
run:
	go run cmd/passgen.go
clean:
	rm passgen