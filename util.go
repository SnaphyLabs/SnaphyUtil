package SnaphyUtil

import (
	"encoding/json"
	"fmt"
	"crypto/md5"
)

// Etag computes an etag based on containt of the payload
func GenEtag(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(b)), nil
}


