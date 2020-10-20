package tools

import (
	"crypto"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

// FileCrypto method.
func FileCrypto(filename, method string) string {
	file, err := os.Open("storage/" + filename)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer func() { _ = file.Close() }()

	fileHash := cryptoFactory(method)
	h := fileHash.New()
	_, err = io.Copy(h, file)
	if err != nil {
		log.Println(err)
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))

}

func cryptoFactory(method string) crypto.Hash {
	switch strings.ToLower(method) {
	case "sha1":
		return crypto.SHA1
	case "sha256":
		return crypto.SHA256
	case "md5":
		return crypto.MD5
	default:
		return 0
	}
}
