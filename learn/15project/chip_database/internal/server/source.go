package server

import (
	cerr "chip_database/error"
	"chip_database/internal/model"
	"chip_database/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

type Source struct {
	*Base
	sourceSrv *service.SourceService
}

func NewSource(sourceSrv *service.SourceService) *Source {
	return &Source{sourceSrv: sourceSrv}
}

func (s *Source) Upload(ctx *gin.Context) {
	condition := &model.Source{}
	if err := s.Check(ctx, condition); err != nil {
		s.ParamErr(ctx, err.Error(), condition)
		return
	}

	file, mfile, err := ctx.Request.FormFile("file")

	if err != nil {
		s.ParamErr(ctx, err.Error(), condition)
		return
	}

	remote, err := s.sourceSrv.FormatFromRemote(file)
	fmt.Println("source data len", mfile.Size)
	fmt.Println("format result", remote, err)

	fullDirInfo, err := s.sourceSrv.GenernalFullDir(mfile.Filename, mfile.Size)

	if err != nil {
		s.Err(ctx, cerr.UPLOAD_ERR_CODE, err.Error(), condition)
		return
	}

	md5, err := s.sourceSrv.SourceMd5(file)

	if err != nil {
		s.Err(ctx, cerr.UPLOAD_ERR_CODE, err.Error(), condition)
		return
	}

	if err := ctx.SaveUploadedFile(mfile, fullDirInfo); err != nil {
		s.Err(ctx, cerr.UPLOAD_ERR_CODE, err.Error(), condition)
		return
	}

	condition.Md5 = md5
	condition.Path = fullDirInfo
	condition.OriginalFileName = mfile.Filename
	if condition.TestId > 0 && condition.Id > 0 {
		condition.SetActivated()
		err = s.sourceSrv.Update(condition)
	} else {
		condition.SetNotActivated()
		err = s.sourceSrv.Create(condition)
	}

	if err != nil {
		s.Err(ctx, cerr.UPLOAD_ERR_CODE, err.Error(), condition)
		return
	}

	s.Success(ctx, nil)
}

func (s *Source) Delete(ctx *gin.Context) {
	condition := &model.DeleteSource{}
	if err := s.CheckJson(ctx, condition); err != nil {
		s.ParamErr(ctx, err.Error(), condition)
		return
	}
	if err := s.sourceSrv.Delete(condition.Id); err != nil {
		s.DBErr(ctx, err.Error(), condition)
		return
	}
	s.Success(ctx, condition)
}

func (s *Source) Update(ctx *gin.Context) {

}

func (s *Source) Find(ctx *gin.Context) {

}

func (s *Source) FetchAllByTestId() {

}

func (s *Source) TmpTest(ctx *gin.Context) {

	r := ctx.Request.Body
	b := make([]byte, 0)
	for {
		tb := make([]byte, 1024)
		n, err := r.Read(tb)
		tb = tb[0:n]
		b = append(b, tb...)
		if err == io.EOF {
			fmt.Println("client read over,total count byte :............", len(b))
			break
		}
		if err != nil {
			fmt.Println("client read err")
			break
		}
	}
	s.Success(ctx, len(b))
}
