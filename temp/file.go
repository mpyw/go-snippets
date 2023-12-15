package temp

import (
	"errors"
	"os"
	"runtime"
)

func NewFile() (*File, error) {
	f, err := os.CreateTemp("", "*")
	if err != nil {
		return nil, err
	}
	tf := &File{f}
	runtime.SetFinalizer(tf, func(tf *File) error {
		return tf.Close()
	})
	return tf, nil
}

type File struct {
	*os.File
}

func (f *File) Close() error {
	defer runtime.SetFinalizer(f, nil)
	return errors.Join(f.File.Close(), os.Remove(f.File.Name()))
}
