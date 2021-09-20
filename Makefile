gen-cal:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:calculator
run-server:
	go run server/server.go
run-client:
	go run client/client.go