run:
	go run cmd/main/main.go

run-bin:
	bin/test_release

run-bin-o:
	bin/test_release_optimized

build:
	go build -o bin/test_release cmd/main/main.go

build-o:
	go build -tags=ebitenginesinglethread -o bin/test_release_optimized cmd/main/main.go
