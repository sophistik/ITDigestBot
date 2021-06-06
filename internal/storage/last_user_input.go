package storage

import (
	"errors"

	"github.com/sophistik/ITDigestBot/internal/repos"
)

var NotFound error = errors.New("Not found")

// todo добавить мютексы
type LastUserInputStorage struct {
	Input map[int64]string
}

func NewLastUserInputStorage() (repos.LastUserInput, error) {
	return &LastUserInputStorage{
		Input: make(map[int64]string, 0),
	}, nil
}

func (s *LastUserInputStorage) Create(id int64, input string) error {
	s.Input[id] = input

	return nil
}

func (s *LastUserInputStorage) Delete(id int64) error {
	delete(s.Input, id)

	return nil
}

func (s *LastUserInputStorage) Find(id int64) (string, error) {
	str, ok := s.Input[id]
	if !ok {
		return "", NotFound
	}

	return str, nil
}
