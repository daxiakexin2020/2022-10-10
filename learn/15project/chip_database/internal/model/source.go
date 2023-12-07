package model

import (
	"chip_database/util"
	"os"
	"path/filepath"
	"strings"
)

type Source struct {
	Id         int64  `json:"id" form:"id"`
	TestId     int    `json:"test_id" form:"test_id"`
	Type       string `json:"type" form:"type" validate:"required"`
	Path       string `json:"path" form:"path"`
	Md5        string `json:"md5" form:"md5"`
	IsActivate int    `json:"is_activate" form:"is_activate"`
}

type DeleteSource struct {
	Id int64 `json:"id" form:"id" validate:"required"`
}

func (s *Source) IsActivated() bool {
	return s.IsActivate == 2
}

func (s *Source) SetActivated() {
	s.IsActivate = 2
}

func (s *Source) SetNotActivated() {
	s.IsActivate = 1
}

const (
	defaultInnerPath = "/default"
	saveDir          = "/source"
)

const maxUploadSize = int64(10 << 20)

const (
	pdf   = "/pdf/"
	img   = "/img/"
	excel = "/excel/"
)

const (
	spdf  = ".pdf"
	spng  = ".png"
	sjpeg = ".jpeg"
	sjpg  = ".sjpg"
	sbmp  = ".bmp"
	sgif  = ".gif"
	scsv  = ".csv"
	sxlsx = ".xlsx"
	sxls  = ".xls"
)

var pathTypeMap = map[string]string{
	spdf:  pdf,
	spng:  img,
	sjpeg: img,
	sjpg:  img,
	sbmp:  img,
	sgif:  img,
	scsv:  excel,
	sxlsx: excel,
	sxls:  excel,
}

var allowType = map[string]struct{}{
	spdf:  struct{}{},
	spng:  struct{}{},
	sjpeg: struct{}{},
	sjpg:  struct{}{},
	sbmp:  struct{}{},
	sgif:  struct{}{},
	scsv:  struct{}{},
	sxlsx: struct{}{},
	sxls:  struct{}{},
}

const (
	thinning_process      = "thinning_process"
	encapsulation_process = "encapsulation_process"
	specification         = "specification"
)

var typeMap = map[string]string{
	thinning_process:      "薄化工艺流程",
	encapsulation_process: "封装工艺流程",
	specification:         "规格书",
}

func GetTypeTitle(typ string) string {
	return typeMap[typ]
}

func getInnerPath(typ string) string {
	if path, has := pathTypeMap[strings.ToLower(typ)]; has {
		return path
	}
	return defaultInnerPath
}

func GenernalFullDir(prefix, sourceType string) (string, error) {

	innerPath := getInnerPath(sourceType)

	newName := prefix + "_" + util.UniqId() + sourceType

	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	absPath, err := filepath.Abs(currentDir)
	if err != nil {
		return "", err
	}
	full := filepath.Join(absPath, "../") + saveDir + innerPath + newName

	return full, nil
}

func IsAllow(tye string) bool {
	if _, has := allowType[strings.ToLower(tye)]; has {
		return true
	}
	return false
}

func IsOverflowMaxUploadSize(size int64) bool {
	return size > maxUploadSize
}
