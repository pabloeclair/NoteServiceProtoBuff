package generate

//go:generate protoc -I=../api --go_out=.. --go-grpc_out=.. ../api/note.proto
