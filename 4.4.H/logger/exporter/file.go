package exporter

import (
	"os"
)

type File struct {
	*Base
}

func NewFile(fileName string) *File {
	f := &File{&Base{}}
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	f.SetOutput(file)
	return f
}
