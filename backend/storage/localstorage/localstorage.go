package localstorage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"orbit_nft/storage"
	"os"
	"path/filepath"
)

type Local struct {
	baseDir string
}

var _ storage.Storage = (*Local)(nil)

func New(baseDir string) (*Local, error) {
	storage := &Local{
		baseDir: baseDir,
	}
	err := safeMkdir(baseDir)
	if err != nil {
		return nil, err
	}
	return storage, nil
}

func (l *Local) AddFile(path string) (string, error) {
	file, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if file.IsDir() {
		return "", errors.New("path is not a file")
	}
	baseDir, err := filepath.Abs(l.baseDir)
	if err != nil {
		return "", err
	}

	hash, err := hashFile(path)
	if err != nil {
		return "", err
	}
	ext := filepath.Ext(path)
	dir := filepath.Dir(file.Name())
	newName := hash + ext
	newPath := filepath.Join(baseDir, dir, newName)
	if isFileExists(newPath) {
		return newName, nil
	}

	oldFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer oldFile.Close()
	newFile, err := os.Create(newPath)
	if err != nil {
		return "", err
	}
	defer newFile.Close()
	if _, err = io.Copy(newFile, oldFile); err != nil {
		return "", err
	}
	if err = newFile.Sync(); err != nil {
		return "", err
	}

	return newName, nil
}

// TODO
func (l *Local) AddDir(path string) (string, error) {
	return "", errors.New("not yet implemented")
}

func safeMkdir(dirPath string) error {
	_, err := os.Stat(dirPath)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func hashFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	hash := sha256.New()
	_, err = io.Copy(hash, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func isFileExists(pathFile string) bool {
	_, err := os.Stat(pathFile)
	return err == nil
}
