package db

import "os"

type JSONDB struct {
	FilePath string
}

func (f *JSONDB) Read() ([]byte, error) {
	data, err := os.ReadFile(f.FilePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *JSONDB) Write(data []byte) error {
	err := os.WriteFile(f.FilePath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
