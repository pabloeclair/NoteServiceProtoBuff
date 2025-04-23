package tests

import (
	"context"
	"fmt"
	"log"
	"project11/internal/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestService(addrs string) error {
	conn, err := grpc.Dial(fmt.Sprintf("localhost%s", addrs), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := protos.NewNoteServiceClient(conn)

	noteString := &protos.NoteString{Name: "Ютуб каналы", Content: "MrLololoshka, Slimecicle"}
	resCreate, err := client.CreateNote(context.Background(), noteString)
	if err != nil {
		return fmt.Errorf("CreateNote: %w", err)
	}
	if resCreate.GetId() != 1 {
		return fmt.Errorf("CreateNote: expected: id = 1; actual: id = %d", resCreate.GetId())
	}

	resGet, err := client.GetNote(context.Background(), resCreate)
	if err != nil {
		return fmt.Errorf("GetNote: %w", err)
	}
	if resGet.GetName() != "Ютуб каналы" || resGet.GetContent() != "MrLololoshka, Slimecicle" {
		return fmt.Errorf(
			`GetNote: expected: Name = "Ютуб каналы", Content = "MrLololoshka, Slimecicle"; actual: Name = "%s", Content = "%s"`,
			resGet.GetName(), resGet.GetContent(),
		)
	}

	return nil
}
