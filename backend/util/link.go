package util

import (
	"net/url"
	"path"
)

func ToLink(base string, elem ...string) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	if len(elem) > 0 {
		elem = append([]string{b.Path}, elem...)
		b.Path = path.Join(elem...)
	}
	return b.String(), nil
}
