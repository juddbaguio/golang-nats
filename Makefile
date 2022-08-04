server:
	go run sub.go

client = 1
run-client:
	go run ./pub --client $(client)