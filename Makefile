build:
	go build -o ./build ./src

run:
	go run ./src

test:
	go test ./src/... -v

build-docker:
	docker build -t go-service .

run-docker:
	docker run -it --rm -p 8080:8080 -e HOST:"0.0.0.0" go-service

generate-swag:
	${GOBIN}/swag init -d ./src/
