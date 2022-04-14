package metadata

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"orbit_nft/storage"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/gosimple/slug"
)

type Metadata struct {
	baseURI string // NFT's base URI (defined in contract NFT)
	storage storage.Storage
}

func NewMetadata(baseURI string, storage storage.Storage) *Metadata {
	return &Metadata{
		baseURI: baseURI,
		storage: storage,
	}
}

func (m *Metadata) GenerateMetadata(rarityPath string, row []string) (string, error) {
	var record []string
	for _, item := range row {
		record = append(record, strings.ReplaceAll(item, "\n", " "))
	}
	tokenTypePath := filepath.Dir(rarityPath)
	tokenType := filepath.Base(tokenTypePath)

	imagePath := filepath.Join(rarityPath, "images")
	imageCid, err := m.getImageCid(m.storage, imagePath, record[2])
	if err != nil {
		return "", err
	}

	metadataRecord, err := generateToken(tokenType, m.baseURI, imageCid, record)
	if err != nil {
		return "", err
	}
	contentTemplate, err := os.ReadFile(filepath.Join(rarityPath, "template.stub"))
	if err != nil {
		return "", err
	}
	temp, err := template.New(record[0]).Parse(string(contentTemplate))
	if err != nil {
		return "", err
	}

	// save file metadata
	var buffer bytes.Buffer
	temp.Execute(&buffer, metadataRecord)
	metadataPath := filepath.Join(rarityPath, "metadata")
	if err = os.MkdirAll(metadataPath, 0755); err != nil {
		return "", err
	}
	files, err := ioutil.ReadDir(metadataPath)
	if err != nil {
		return "", err
	}
	metadataName := slug.Make(fmt.Sprintf("%s_%s_%s", tokenType, record[0], strconv.Itoa(len(files)))) + ".json"
	filePath := filepath.Join(metadataPath, metadataName)

	if err = ioutil.WriteFile(filePath, buffer.Bytes(), 0777); err != nil {
		return "", err
	}
	metadataCid, err := m.storage.AddFile(filePath)
	if err != nil {
		return "", err
	}
	return metadataCid, nil
}

func (m *Metadata) getImageCid(storage storage.Storage, imagePath, image string) (string, error) {
	imageFullPath := filepath.Join(imagePath, image)
	_, err := os.Stat(imageFullPath)
	if err != nil {
		return "", err
	}

	imageCid, err := storage.AddFile(imageFullPath)
	if err != nil {
		return "", err
	}
	return imageCid, nil
}
