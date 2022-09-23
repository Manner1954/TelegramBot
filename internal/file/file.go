package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	Id   uuid.UUID
	Name string
	Data []byte
}

func NewFile(fileName string, blob []byte) (*File, error) {
	fileId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &File{
		Id:   fileId,
		Name: fileName,
		Data: blob,
	}, err
}

func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v", f.Name, f.Id)
}
