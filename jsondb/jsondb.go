package jsondb

import "os"

type JSONDB struct {
	FilePath string
}

func (j *JSONDB) Read() ([]byte, error) {
	data, err := os.ReadFile(j.FilePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (j *JSONDB) Write(data []byte) error {
	err := os.WriteFile(j.FilePath, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
