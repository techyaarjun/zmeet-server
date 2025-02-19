package store

import (
	"github.com/google/uuid"
	"sync"
	"zmeet/pkg/user"
)

type Store struct {
	mu         sync.Mutex
	ZMeetUsers map[uuid.UUID]*user.ZMeetUser
}

func NewStore() *Store {
	return &Store{
		ZMeetUsers: make(map[uuid.UUID]*user.ZMeetUser),
	}
}

func (s *Store) AddZMeetUser(z *user.ZMeetUser) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ZMeetUsers[z.ID()] = z
}

func (s *Store) GetZMeetUser(id uuid.UUID) *user.ZMeetUser {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.ZMeetUsers[id]
}

func (s *Store) RemoveZMeetUser(id uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.ZMeetUsers, id)
}

func (s *Store) RemoveAllZMeetUser() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ZMeetUsers = make(map[uuid.UUID]*user.ZMeetUser)
}

func (s *Store) GetAllZMeetUsers() []*user.ZMeetUser {
	s.mu.Lock()
	defer s.mu.Unlock()

	var res []*user.ZMeetUser
	for _, z := range s.ZMeetUsers {
		if z.Connected() {
			res = append(res, z)
		}
	}
	return res
}
