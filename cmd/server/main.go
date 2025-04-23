package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	service "project11/internal"
	"project11/protos"
	"project11/tests"
	"syscall"

	"google.golang.org/grpc"
)

func main() {

	if len(os.Args) != 2 {
		panic("invalid argument: принимается только один аргумент – строка для прослушивания")
	}

	addrs := os.Args[1]

	serverCtx, serverCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer serverCancel()

	go func() {
		lis, err := net.Listen("tcp", addrs)
		if err != nil {
			log.Fatalf("Ошибка запуска слушателя: %v", err)
		}
		grpcServer := grpc.NewServer()

		protos.RegisterNoteServiceServer(grpcServer, &service.NoteServer{})
		log.Printf("gRPC-сервер запущен на порту %s...", os.Args[1])

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Ошибка работы gRPC-сервера: %v", err)
		}
	}()

	go func() {
		if err := tests.TestService(addrs); err != nil {
			log.Fatalf("Ошибка тестов: %v", err)
			serverCancel()
		}
		log.Println("Сервер готов к работе!")
	}()

	<-serverCtx.Done()

}
