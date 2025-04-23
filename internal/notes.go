package noteservice

import (
	"context"
	"fmt"
	"project11/internal/protos"
	"strings"
	"sync"

	"slices"

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
	var id int32
	if len(s.notes) == 0 {
		id = 1
	} else {
		for k := range s.notes {
			id = k
		}
		id++
	}

	s.mu.Lock()
	if s.notes == nil {
		s.notes = make(map[int32]note)
	}
	s.notes[id] = note{name: req.GetName(), content: req.GetContent()}
	s.mu.Unlock()
	return &protos.NoteId{Id: id}, nil
}

func (s *NoteServer) GetNote(ctx context.Context, req *protos.NoteId) (*protos.NoteString, error) {

	id := req.GetId()

	s.mu.RLock()
	res := s.notes[id]
	s.mu.RUnlock()
	if res.name == "" && res.content == "" {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("note with id = %d not exists", id))
	}
	noteRes := &protos.NoteString{Name: res.name, Content: res.content}
	return noteRes, nil
}

func (s *NoteServer) UpdateNote(ctx context.Context, req *protos.UpdateNoteRequest) (*protos.Empty, error) {

	if req.GetName() == "" && req.GetContent() == "" {
		return nil, status.Error(codes.InvalidArgument, "fields should not be empty")
	}

	s.mu.RLock()
	res := s.notes[req.GetId()]
	s.mu.RUnlock()

	if res.name == "" && res.content == "" {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("note with id = %d not exists", req.GetId()))
	}

	s.mu.Lock()
	s.notes[req.GetId()] = note{name: req.GetName(), content: req.GetContent()}
	s.mu.Unlock()
	return &protos.Empty{}, nil
}

func (s *NoteServer) SearchNotes(ctx context.Context, req *protos.SearchNotesRequest) (*protos.NoteIdRepeated, error) {

	res := make([]int32, 0)

	pattern := req.GetPattern()
	s.mu.RLock()
	for id, note := range s.notes {
		if strings.Contains(strings.ToLower(note.name), pattern) || strings.Contains(strings.ToLower(note.content), pattern) {
			res = append(res, id)
		}
	}
	s.mu.RUnlock()
	slices.Sort(res)

	return &protos.NoteIdRepeated{Id: res}, nil
}

func (s *NoteServer) DeleteNote(ctx context.Context, req *protos.NoteId) (*protos.Empty, error) {

	id := req.GetId()

	s.mu.RLock()
	note := s.notes[id]
	s.mu.RUnlock()
	if note.name == "" && note.content == "" {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("note with id = %d not exists", id))
	}

	s.mu.Lock()
	delete(s.notes, id)
	s.mu.Unlock()
	return &protos.Empty{}, nil
}
