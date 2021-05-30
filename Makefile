generate:
	go run transformation/gen/main.go
	go run transformation/gen/generator.go
	rm -rf transformation/gen/generator.go
