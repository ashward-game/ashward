package util

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func readFileJson(file string) (map[string]string, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return data, nil
}

// WriteToJSONFile overwrites if key already exists.
// It also creates file with empty content if file does not exist.
func WriteToJSONFile(file string, key, value string) error {
	if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
		if err := ioutil.WriteFile(file, []byte("{}"), 0600); err != nil {
			return err
		}
	}
	address, err := readFileJson(file)
	if err != nil {
		return err
	}

	address[key] = value

	data, err := json.Marshal(address)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, data, 0600)
}

func GetContractAddress(file string, name string) (string, error) {
	obj, err := readFileJson(file)
	if err != nil {
		return "", err
	}
	if obj[name] == "" {
		return "", errors.New("object key not found")
	}
	return obj[name], nil
}

func GetSecrets(file string, key string) (string, error) {
	obj, err := readFileJson(file)
	if err != nil {
		return "", err
	}
	if obj[key] == "" {
		return "", errors.New("object key not found")
	}
	return obj[key], nil
}

func ReadFileCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func HexToBigInt(s string) *big.Int {
	var a big.Int
	a.SetBytes(common.FromHex(s))
	return &a
}

func XorBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("xor two slices: length mismatch")
	}

	buf := make([]byte, len(a))

	for i := range a {
		buf[i] = a[i] ^ b[i]
	}

	return buf, nil
}

func GetFilesRecursively(baseDir string, fileName string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(baseDir, func(filePath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.EqualFold(d.Name(), fileName) {
			files = append(files, filePath)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func RandomMinToMax(min, max string) (int, error) {
	iMin, err := strconv.Atoi(min)
	if err != nil {
		return 0, err
	}
	iMax, err := strconv.Atoi(max)
	if err != nil {
		return 0, err
	}
	return rand.Intn(iMax-iMin) + iMin, nil
}

func StrSliceContains(arr []string, s string) bool {
	for _, ss := range arr {
		if s == ss {
			return true
		}
	}
	return false
}
