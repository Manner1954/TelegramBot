package storage

import (
	"fmt"

	"github.com/Manner1954/TelegramBot/storage/internal/file"

	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error) {
	file, err := file.NewFile(fileName, blob)

	if err != nil {
		return nil, err
	}

	s.files[file.Id] = file
	return file, err
}

func (s *Storage) GetById(fileId uuid.UUID) (*file.File, error) {
	file, ok := s.files[fileId]

	if !ok {
		return nil, fmt.Errorf("File %v not found", fileId)
	}

	return file, nil
}
