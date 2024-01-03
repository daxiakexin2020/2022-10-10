package server

import (
	"chip_database/internal/model"
	"chip_database/internal/service"
	"github.com/gin-gonic/gin"
)

type TestItem struct {
	*Base
	testItemSrv *service.TestItemService
	sourceSrv   *service.SourceService
}

func NewTestItem(testItemSrv *service.TestItemService, sourceSrv *service.SourceService) *TestItem {
	return &TestItem{testItemSrv: testItemSrv, sourceSrv: sourceSrv}
}

func (tt *TestItem) Create(ctx *gin.Context) {
	condition := &model.TestItem{}
	if err := tt.CheckJson(ctx, condition); err != nil {
		tt.ParamErr(ctx, err.Error(), condition)
		return
	}
	if err := tt.testItemSrv.Create(condition); err != nil {
		tt.DBErr(ctx, err.Error(), condition)
		return
	}

	if len(condition.SourceIds) > 0 && condition.Id > 0 {
		if err := tt.sourceSrv.UpdateTestIdByIds(condition.SourceIds, condition.Id); err != nil {
			tt.DBErr(ctx, err.Error(), condition)
			return
		}
	}

	tt.Success(ctx, condition)
}

func (tt *TestItem) Delete(ctx *gin.Context) {
	condition := &model.DeleteBaseInfo{}
	if err := tt.CheckJson(ctx, condition); err != nil {
		tt.ParamErr(ctx, err.Error(), condition)
		return
	}
	if err := tt.testItemSrv.Delete(condition.Id); err != nil {
		tt.DBErr(ctx, err.Error(), condition)
		return
	}
	tt.Success(ctx, condition)
}

func (tt *TestItem) Update(ctx *gin.Context) {

	condition := &model.TestItem{}
	if err := tt.CheckJson(ctx, condition); err != nil {
		tt.ParamErr(ctx, err.Error(), condition)
		return
	}
	if condition.Id == 0 {
		tt.ParamErr(ctx, "update,id is needed", condition)
		return
	}
	if err := tt.testItemSrv.Update(condition); err != nil {
		tt.DBErr(ctx, err.Error(), condition)
		return
	}

	tt.Success(ctx, condition)
}

func (tt *TestItem) FetchAllByBaseId(ctx *gin.Context) {
	condition := &model.ListTestItem{}
	if err := tt.CheckJson(ctx, condition); err != nil {
		tt.ParamErr(ctx, err.Error(), condition)
		return
	}
	items, err := tt.testItemSrv.GetAllByBaseId(condition.BaseId)
	if err != nil {
		tt.DBErr(ctx, err.Error(), condition)
		return
	}
	tt.Success(ctx, items)
}

func (tt *TestItem) Find(ctx *gin.Context) {
	condition := &model.FindTestItem{}
	if err := tt.CheckJson(ctx, condition); err != nil {
		tt.ParamErr(ctx, err.Error(), condition)
		return
	}
	item, err := tt.testItemSrv.Find(condition.Id)
	if err != nil {
		tt.DBErr(ctx, err.Error(), condition)
		return
	}
	tt.Success(ctx, item)
}
