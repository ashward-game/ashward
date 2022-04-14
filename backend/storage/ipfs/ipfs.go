package ipfs

import (
	"context"
	"errors"
	"orbit_nft/storage"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFS struct {
	sh *shell.Shell
}

var _ storage.Storage = (*IPFS)(nil)

func NewShell(host string) (*IPFS, error) {
	sh := shell.NewShell(host)
	err := sh.
		Request("version").
		Exec(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &IPFS{sh: sh}, nil
}

func (i *IPFS) AddDir(path string) (string, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if !fi.IsDir() {
		return "", errors.New("path is not a directory")
	}
	return i.sh.AddDir(path)
}

func (i *IPFS) AddFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return i.sh.Add(file)
}

func (i *IPFS) List(path string) ([]*shell.LsLink, error) {
	return i.sh.List(path)
}
