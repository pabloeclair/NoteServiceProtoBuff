package tests

import (
	"context"
	"errors"
	"fmt"
	"log"
	"project11/internal/protos"
	"slices"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestService(addrs string) error {
	conn, err := grpc.Dial(fmt.Sprintf("localhost%s", addrs), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := protos.NewNoteServiceClient(conn)

	// Test Create
	noteString := &protos.NoteString{Name: "Любимые ютуб каналы", Content: "MrLololoshka, Slimecicle"}
	resCreate, err := client.CreateNote(context.Background(), noteString)
	if err != nil {
		return fmt.Errorf("CreateNote: error: %w", err)
	}
	if resCreate.GetId() != 1 {
		return fmt.Errorf("CreateNote: expected: id = 1; actual: id = %d", resCreate.GetId())
	}

	noteString = &protos.NoteString{Name: "Книги", Content: "Горменгаст"}
	resCreate, err = client.CreateNote(context.Background(), noteString)
	if err != nil {
		return fmt.Errorf("CreateNote: error: %w", err)
	}
	if resCreate.GetId() != 2 {
		return fmt.Errorf("CreateNote: expected: id = 2; actual: id = %d", resCreate.GetId())
	}

	noteString = &protos.NoteString{Name: "ЯП", Content: "Python, ЛЮБИМЫЙ Go, Java"}
	resCreate, err = client.CreateNote(context.Background(), noteString)
	if err != nil {
		return fmt.Errorf("CreateNote: error: %w", err)
	}
	if resCreate.GetId() != 3 {
		return fmt.Errorf("CreateNote: expected: id = 3; actual: id = %d", resCreate.GetId())
	}

	noteStringWrong := &protos.NoteString{Name: "", Content: ""}
	resCreate, err = client.CreateNote(context.Background(), noteStringWrong)
	if err == nil {
		return fmt.Errorf(
			`CreateNote (wrong): expected: err = invalid argument; actual: err = nil, Id = %d"`, resCreate.GetId(),
		)
	}
	if !errors.Is(err, status.Error(codes.InvalidArgument, "fields should not be empty")) {
		return fmt.Errorf(
			`CreateNote (wrong): expected: err = not found; actual: err = %v"`, err,
		)
	}
	log.Println("TestCreateNote pass")

	// Test Get
	resGet, err := client.GetNote(context.Background(), &protos.NoteId{Id: 1})
	if err != nil {
		return fmt.Errorf("GetNote: %w", err)
	}
	if resGet.GetName() != "Любимые ютуб каналы" || resGet.GetContent() != "MrLololoshka, Slimecicle" {
		return fmt.Errorf(
			`GetNote: expected: Name = "Любимые ютуб каналы", Content = "MrLololoshka, Slimecicle"; actual: Name = "%s", Content = "%s"`,
			resGet.GetName(), resGet.GetContent(),
		)
	}

	resGetWrong, err := client.GetNote(context.Background(), &protos.NoteId{Id: 999})
	if err == nil {
		return fmt.Errorf(
			`GetNote (wrong): expected: err = not found; actual: err = nil, Name = "%s", Content = "%s"`,
			resGetWrong.GetName(), resGetWrong.GetContent(),
		)
	}
	if !errors.Is(err, status.Error(codes.NotFound, "note with id = 999 not exists")) {
		return fmt.Errorf(
			`GetNote (wrong): expected: err = not found; actual: err = %v"`, err,
		)
	}
	log.Println("TestGetNote pass")

	// Test Update
	noteUpdate := &protos.UpdateNoteRequest{
		Id:      1,
		Name:    "Любимые ютуб каналы",
		Content: "MrLololoshka, Slimecicle, Kyngstom Myles",
	}

	_, err = client.UpdateNote(context.Background(), noteUpdate)
	if err != nil {
		return fmt.Errorf("UpdateNote (update): error: %w", err)
	}

	check, err := client.GetNote(context.Background(), &protos.NoteId{Id: 1})
	if err != nil {
		return fmt.Errorf("UpdateNote (get): error: %w", err)
	}

	if check.GetName() != "Любимые ютуб каналы" || check.GetContent() != "MrLololoshka, Slimecicle, Kyngstom Myles" {
		return fmt.Errorf(
			`UpdateNote (get): Name = "Любимые ютуб каналы", Content = "MrLololoshka, Slimecicle, Kyngstom Myles"; actual: Name = "%s", Content = "%s"`,
			check.GetName(), check.GetContent(),
		)
	}
	log.Println("TestUpdateNote pass")

	// Test Search
	listId, err := client.SearchNotes(context.Background(), &protos.SearchNotesRequest{Pattern: "любим"})
	if err != nil {
		return fmt.Errorf("SearchNotes: error: %w", err)
	}
	if !slices.Equal([]int32{1, 3}, listId.GetId()) {
		return fmt.Errorf(
			`SerchNotes: expected: %v; actual: %v`, []int32{1, 3}, listId.GetId(),
		)
	}
	log.Println("TestSearchNotes pass")

	return nil
}
