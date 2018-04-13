package tools

import (
	"os"
)

type Directory struct {
}

func RuntimeDir() (dir string) {
	dir, _ = os.Getwd()
	return
}

func NewDirectory() *Directory {
	return &Directory{}
}

func (directory *Directory) Create(name string) (err error) {
	if len(name) == 0 {
		return
	}
	err = os.MkdirAll(name, os.ModePerm)
	return
}

func (directory *Directory) Delete(name string) (err error) {
	if len(name) == 0 {
		return
	}
	err = os.RemoveAll(name)
	return
}
