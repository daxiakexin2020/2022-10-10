package route

import (
	"chip_database/internal/server"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	engine   *gin.Engine
	baseInfo *server.BaseInfo
	testItem *server.TestItem
	source   *server.Source
}

func NewApiRouter(engine *gin.Engine, baseInfo *server.BaseInfo, testItem *server.TestItem, source *server.Source) *ApiRouter {
	return &ApiRouter{engine: engine, baseInfo: baseInfo, testItem: testItem, source: source}
}

func (r *ApiRouter) RegisterHandlers() {

	api := r.engine.Group("/api/")

	baseInfo := api.Group("base_info/")
	{
		baseInfo.POST("create", r.baseInfo.Create)
		baseInfo.DELETE("delete", r.baseInfo.Delete)
		baseInfo.PUT("update", r.baseInfo.Update)
		baseInfo.GET("list", r.baseInfo.List)
		baseInfo.GET("tree", r.baseInfo.Tree)
		baseInfo.GET("find", r.baseInfo.Find)
	}
	testItem := api.Group("test_item/")
	{
		testItem.POST("create", r.testItem.Create)
		testItem.DELETE("delete", r.testItem.Delete)
		testItem.PUT("update", r.testItem.Update)
		testItem.POST("list", r.testItem.FetchAllByBaseId)
	}

	source := api.Group("source/")
	{
		source.POST("upload", r.source.Upload)
		source.POST("delete", r.source.Delete)
		source.POST("tmp_test", r.source.TmpTest)
	}
}
