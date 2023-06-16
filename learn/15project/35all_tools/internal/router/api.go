package router

import (
	"35all_tools/internal/handler"
	"35all_tools/internal/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	engine               *gin.Engine
	jsonhandler          *handler.JsonHanler
	endehandler          *handler.EnDeHandler
	symmetryEndehandler  *handler.SymmetryEnDeHandler
	Comprehensivehandler *handler.ComprehensiveHandler
}

func NewApiRouter(engine *gin.Engine, jh *handler.JsonHanler, edh *handler.EnDeHandler, sedh *handler.SymmetryEnDeHandler, ch *handler.ComprehensiveHandler) *ApiRouter {
	return &ApiRouter{
		engine:               engine,
		jsonhandler:          jh,
		endehandler:          edh,
		symmetryEndehandler:  sedh,
		Comprehensivehandler: ch,
	}
}

func (ar *ApiRouter) RegisterHandlers() {

	ar.engine.Use(middleware.LogMiddleware())

	jsonGroup := ar.engine.Group("/json/")
	{
		jsonGroup.POST("check_json", ar.jsonhandler.CheckJson)
		jsonGroup.POST("json_to_golang_struct", ar.jsonhandler.JsonToGolangStruct)
	}

	endeGroup := ar.engine.Group("/ende/")
	{
		endeGroup.POST("md5_encode", ar.endehandler.MD5Encode)
		endeGroup.POST("url_16_encode", ar.endehandler.Url16Encode)
		endeGroup.POST("base64_encode", ar.endehandler.Base64Encode)
		endeGroup.POST("base64_decode", ar.endehandler.Base64Decode)
		endeGroup.POST("escape", ar.endehandler.Escape)
		endeGroup.POST("deescape", ar.endehandler.DeEscape)
	}

	symmetryEndeGroup := ar.engine.Group("/symmetry_ende/")
	{
		symmetryEndeGroup.POST("encode", ar.symmetryEndehandler.Encode)
		symmetryEndeGroup.POST("decode", ar.symmetryEndehandler.Decode)
	}

	comprehensiveGroup := ar.engine.Group("/comprehensive/")
	{
		comprehensiveGroup.POST("ip_info", ar.Comprehensivehandler.IpInfo)
		comprehensiveGroup.POST("domain_map_ip", ar.Comprehensivehandler.DomainMapIp)
		comprehensiveGroup.POST("image_base64", ar.Comprehensivehandler.ImageBase64)
	}
}
