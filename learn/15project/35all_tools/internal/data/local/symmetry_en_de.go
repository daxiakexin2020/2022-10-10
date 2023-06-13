package local

import (
	"35all_tools/internal/model"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"sync"
)

type SymmetryEnDeRepository struct {
	strategy map[string]cipher.Block
	ende     model.EnDeRepo
	mu       sync.RWMutex
}

type strategys map[model.SYM_TYPE]func(key []byte) (cipher.Block, error)

var (
	gstrategys strategys
	_          (model.SymmetryEnDeRepo) = (*SymmetryEnDeRepository)(nil)
)

func NewSymmetryEnDeRepository(ende model.EnDeRepo) model.SymmetryEnDeRepo {

	sedr := &SymmetryEnDeRepository{ende: ende, strategy: map[string]cipher.Block{}}
	gstrategys = strategys{
		model.SYM_AES: sedr.generateAES,
		model.SYM_DES: sedr.generateDES,
	}

	for k, f := range gstrategys {
		cipher, err := f(model.DefalutSecretKey)
		if err == nil {
			generatekey, err := sedr.generatekey(k, model.DefalutSecretKey)
			if err == nil {
				sedr.strategy[generatekey] = cipher
			}
		}
	}
	return sedr
}

func (sedr *SymmetryEnDeRepository) generatekey(typ model.SYM_TYPE, key []byte) (string, error) {
	s := string(typ) + "_" + string(key)
	return sedr.ende.MD5Encode(s)
}

func (sedr *SymmetryEnDeRepository) generateAES(key []byte) (cipher.Block, error) {
	sk := model.DefalutSecretKey
	if len(key) > 0 {
		sk = key
	}
	return aes.NewCipher(sk)
}

func (sedr *SymmetryEnDeRepository) generateDES(key []byte) (cipher.Block, error) {
	sk := model.DefalutSecretKey
	if len(key) > 0 {
		sk = key
	}
	return des.NewCipher(sk)
}

func (sedr *SymmetryEnDeRepository) factory(typ model.SYM_TYPE, key []byte) (cipher.Block, error) {

	sedr.mu.Lock()
	defer sedr.mu.Unlock()

	sk := model.DefalutSecretKey
	if len(sk) > 0 {
		sk = key
	}

	generatekey, err := sedr.generatekey(typ, sk)
	if err != nil {
		return nil, err
	}

	cipher, ok := sedr.strategy[generatekey]
	if ok {
		return cipher, nil
	}

	f, ok := gstrategys[typ]
	if !ok {
		return nil, errors.New("暂不支持此种策略：" + string(typ))
	}

	cipher, err = f(sk)
	if err != nil {
		return nil, err
	}
	sedr.strategy[generatekey] = cipher
	return cipher, nil
}

func (sedr *SymmetryEnDeRepository) Encode(str string, typ model.SYM_TYPE) (string, error) {
	factory, err := sedr.factory(typ, nil)
	if err != nil {
		return "", err
	}
	var dest []byte
	factory.Encrypt(dest, []byte(str))
	return string(dest), nil
}

func (sedr *SymmetryEnDeRepository) Decode(str string, typ model.SYM_TYPE) (string, error) {
	return "", nil
}
