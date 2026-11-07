package db

type DB interface {
	Read() ([]byte, error)
	Write([]byte) error
}
