package storage

type Storage interface {
	AddFile(string) (string, error)
	AddDir(string) (string, error)
}
