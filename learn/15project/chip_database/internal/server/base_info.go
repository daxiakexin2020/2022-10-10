package server

import (
	cerr "chip_database/error"
	"chip_database/internal/model"
	"chip_database/internal/service"
	"github.com/gin-gonic/gin"
)

type BaseInfo struct {
	*Base
	baseInfoSrv *service.BaseInfoService
	testItemSrv *service.TestItemService
}

func NewBaseInfo(baseInfoSrv *service.BaseInfoService, testItemSrv *service.TestItemService) *BaseInfo {
	return &BaseInfo{baseInfoSrv: baseInfoSrv, testItemSrv: testItemSrv}
}

func (bi *BaseInfo) Create(ctx *gin.Context) {
	condition := &model.BaseInfo{}
	if err := bi.CheckJson(ctx, condition); err != nil {
		bi.ParamErr(ctx, err.Error(), condition)
		return
	}
	if err := bi.baseInfoSrv.Create(condition); err != nil {
		bi.DBErr(ctx, err.Error(), condition)
		return
	}
	bi.Success(ctx, condition)
}

func (bi *BaseInfo) Delete(ctx *gin.Context) {

	condition := &model.DeleteBaseInfo{}
	if err := bi.CheckJson(ctx, condition); err != nil {
		bi.ParamErr(ctx, err.Error(), condition)
		return
	}
	testItem, err := bi.testItemSrv.GetAllByBaseId(condition.Id)
	if err != nil {
		bi.DBErr(ctx, err.Error(), condition)
		return
	}
	if len(testItem) > 0 {
		bi.Err(ctx, cerr.BUSSINE_ERR_CODE, "this base info of test item aleady has data,please delete", condition)
		return
	}
	if err := bi.baseInfoSrv.Delete(condition.Id); err != nil {
		bi.DBErr(ctx, err.Error(), condition)
		return
	}
	bi.Success(ctx, nil)
}

func (bi *BaseInfo) Update(ctx *gin.Context) {}

func (bi *BaseInfo) List(ctx *gin.Context) {
	list, err := bi.baseInfoSrv.List()
	if err != nil {
		bi.DBErr(ctx, err.Error(), nil)
		return
	}
	bi.Success(ctx, list)
}

func (bi *BaseInfo) Find(ctx *gin.Context) {
	condition := &model.FindBaseInfo{}
	if err := bi.CheckJson(ctx, condition); err != nil {
		bi.ParamErr(ctx, err.Error(), condition)
		return
	}
	baseInfo, err := bi.baseInfoSrv.Find(condition.Id)
	if err != nil {
		bi.DBErr(ctx, err.Error(), condition)
		return
	}

	tree := baseInfo.GenerateTree()
	items, err := bi.testItemSrv.GetAllByBaseId(condition.Id)
	if err != nil {
		bi.DBErr(ctx, err.Error(), condition)
		return
	}
	tree.Children = items
	bi.Success(ctx, tree)
}

func (bi *BaseInfo) Tree(ctx *gin.Context) {
	models, err := bi.baseInfoSrv.List()
	if err != nil {
		bi.DBErr(ctx, err.Error(), nil)
		return
	}

	items, err := bi.testItemSrv.List()
	if err != nil {
		bi.DBErr(ctx, err.Error(), nil)
		return
	}

	tmpTrees := make(map[int]*model.Tree)
	for _, model := range models {
		tmpTrees[model.ID] = model.GenerateTree()
	}

	var trees []*model.Tree
	for _, item := range items {
		baseId := item.BaseId
		if _, has := tmpTrees[baseId]; has {
			tmpTrees[baseId].Children = append(tmpTrees[baseId].Children, item)
		}
	}

	for _, tree := range tmpTrees {
		trees = append(trees, tree)
	}
	bi.Success(ctx, trees)
}
