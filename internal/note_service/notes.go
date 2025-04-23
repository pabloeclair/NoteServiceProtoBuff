package noteservice

import (
	"context"
	"fmt"
	"project11/internal/protos"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NoteServer struct {
	protos.UnimplementedNoteServiceServer
	mu    sync.RWMutex
	notes map[int32]note
}

type note struct {
	name    string
	content string
}

func (s *NoteServer) CreateNote(ctx context.Context, req *protos.NoteString) (*protos.NoteId, error) {

	if req.GetName() == "" && req.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "fields should not be empty")
	}
	id := int32(len(s.notes)) + 1

	s.mu.Lock()
	if s.notes == nil {
		s.notes = make(map[int32]note)
	}
	s.notes[id] = note{name: req.GetName(), content: req.GetContent()}
	s.mu.Unlock()
	return &protos.NoteId{Id: id}, nil
}

func (s *NoteServer) GetNote(ctx context.Context, req *protos.NoteId) (*protos.NoteString, error) {
	s.mu.RLock()
	res := s.notes[req.GetId()]
	if res.name == "" && res.content == "" {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("note with if = %d not exists", req.GetId()))
	}
	s.mu.RUnlock()
	noteRes := &protos.NoteString{Name: res.name, Content: res.content}
	return noteRes, nil
}
