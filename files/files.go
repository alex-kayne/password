package files

import "os"

func ReadFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(data string, fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(data)

	if err != nil {
		return err
	}

	return nil
}
