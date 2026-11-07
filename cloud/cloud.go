package cloud

type CloudDB struct {
	URL string
}

func NewCloudDB(url string) (*CloudDB, error) {
	return &CloudDB{url}, nil
}

func (db *CloudDB) Read() ([]byte, error) {
	return nil, nil
}
func (db *CloudDB) Write(data []byte) error {
	return nil
}
