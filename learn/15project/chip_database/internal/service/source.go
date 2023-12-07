package service

import (
	cerr "chip_database/error"
	"chip_database/internal/data/db"
	"chip_database/internal/model"
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

type SourceService struct {
	sourceDB *db.Source
}

func NewSourceService(sourceDB *db.Source) *SourceService {
	return &SourceService{sourceDB: sourceDB}
}

func (ss *SourceService) Create(m *model.Source) error {
	return ss.sourceDB.Create(m)
}

func (ss *SourceService) Delete(id int64) error {
	return ss.sourceDB.Delete(id)
}

func (ss *SourceService) Update(m *model.Source) error {
	return ss.sourceDB.Update(m)
}

func (ss *SourceService) GenernalFullDir(filename string, size int64) (string, error) {

	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return "", cerr.NO_SUPPORT_FILE_TYPE
	}
	sourceType := filename[index:]
	prefix := filename[0:index]

	if !model.IsAllow(sourceType) {
		return "", cerr.NO_SUPPORT_FILE_TYPE
	}

	if model.IsOverflowMaxUploadSize(size) {
		return "", cerr.OVERFLOW_UPLOAD_MAX_SIZE
	}

	return model.GenernalFullDir(prefix, sourceType)
}

func (ss *SourceService) SourceMd5(r io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
