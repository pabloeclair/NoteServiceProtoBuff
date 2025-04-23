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

	// Test Create
	noteString := &protos.NoteString{Name: "Ютуб каналы", Content: "MrLololoshka, Slimecicle"}
	resCreate, err := client.CreateNote(context.Background(), noteString)
	if err != nil {
		return fmt.Errorf("CreateNote: %w", err)
	}
	if resCreate.GetId() != 1 {
		return fmt.Errorf("CreateNote: expected: id = 1; actual: id = %d", resCreate.GetId())
	}

	// Test Get
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

	// Test Update
	noteUpdate := &protos.UpdateNoteRequest{
		Id:      1,
		Name:    "Ютуб каналы",
		Content: "MrLololoshka, Slimecicle, Kyngstom Myles",
	}
	_, err = client.UpdateNote(context.Background(), noteUpdate)
	if err != nil {
		return fmt.Errorf("UpdateNote (update): %w", err)
	}
	check, err := client.GetNote(context.Background(), &protos.NoteId{Id: 1})
	if err != nil {
		return fmt.Errorf("UpdateNote (get): %w", err)
	}
	if check.GetName() != "Ютуб каналы" || check.GetContent() != "MrLololoshka, Slimecicle, Kyngstom Myles" {
		return fmt.Errorf(
			`UpdateNote (get): Name = "Ютуб каналы", Content = "MrLololoshka, Slimecicle, Kyngstom Myles"; actual: Name = "%s", Content = "%s"`,
			check.GetName(), check.GetContent(),
		)
	}

	return nil
}
