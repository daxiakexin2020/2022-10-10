package local

import (
	"35all_tools/internal/model"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
)

type EnDeRepository struct{}

var _ (model.EnDeRepo) = (*EnDeRepository)(nil)

func NewEnDeRepository() model.EnDeRepo {
	return &EnDeRepository{}
}

func (edr *EnDeRepository) MD5Encode(str string) (string, error) {
	hash := md5.New()
	if _, err := hash.Write([]byte(str)); err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

func (edr *EnDeRepository) Url16Encode(str string) (string, error) {
	//todo
	escape := fmt.Sprintf("%x", str)
	return escape, nil
}

func (edr *EnDeRepository) Base64Encode(str string) (string, error) {
	toString := base64.StdEncoding.EncodeToString([]byte(str))
	return toString, nil
}

func (edr *EnDeRepository) Base64Decode(str string) (string, error) {
	decodeString, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(decodeString), nil
}

func (edr *EnDeRepository) Escape(str string) (string, error) {
	escape := url.QueryEscape(str)
	return escape, nil
}

func (edr *EnDeRepository) DeEscape(str string) (string, error) {
	return url.QueryUnescape(str)
}
