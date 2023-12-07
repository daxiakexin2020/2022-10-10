package server

import (
	cerr "chip_database/error"
	"chip_database/internal/model"
	"chip_database/internal/service"
	"github.com/gin-gonic/gin"
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

	s.Success(ctx, condition)
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