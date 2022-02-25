build:
	go build -o ./cmd/run_client ./cmd/run_client
	go build -o ./cmd/run_server ./cmd/run_server

clientDefault:
	./cmd/run_client/run_client -f ./input.txt

serverDefault:
	./cmd/run_server/run_server
