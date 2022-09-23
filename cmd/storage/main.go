package main

import (
	"fmt"
	"log"

	"github.com/arlatypov/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("Hello"))

	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.Id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("It's restored", restoredFile)
}
