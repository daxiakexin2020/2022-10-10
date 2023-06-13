package router

import (
	"35all_tools/internal/handlers"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	engine               *gin.Engine
	jsonHandler          *handlers.JsonHanler
	endeHandler          *handlers.EnDeHandler
	symmetryEndeHandler  *handlers.SymmetryEnDeHandler
	ComprehensiveHandler *handlers.ComprehensiveHandler
}

func NewApiRouter(engine *gin.Engine, jh *handlers.JsonHanler, edh *handlers.EnDeHandler, sedh *handlers.SymmetryEnDeHandler, ch *handlers.ComprehensiveHandler) *ApiRouter {
	return &ApiRouter{
		engine:               engine,
		jsonHandler:          jh,
		endeHandler:          edh,
		symmetryEndeHandler:  sedh,
		ComprehensiveHandler: ch,
	}
}

func (ar *ApiRouter) RegisterHandlers() {

	jsonGroup := ar.engine.Group("/json/")
	{
		jsonGroup.POST("check_json", ar.jsonHandler.CheckJson)
		jsonGroup.POST("json_to_golang_struct", ar.jsonHandler.JsonToGolangStruct)
	}

	endeGroup := ar.engine.Group("/ende/")
	{
		endeGroup.POST("md5_encode", ar.endeHandler.MD5Encode)
		endeGroup.POST("url_16_encode", ar.endeHandler.Url16Encode)
		endeGroup.POST("base64_encode", ar.endeHandler.Base64Encode)
		endeGroup.POST("base64_decode", ar.endeHandler.Base64Decode)
		endeGroup.POST("escape", ar.endeHandler.Escape)
		endeGroup.POST("deescape", ar.endeHandler.DeEscape)
	}

	symmetryEndeGroup := ar.engine.Group("/symmetry_ende/")
	{
		symmetryEndeGroup.POST("encode", ar.symmetryEndeHandler.Encode)
		symmetryEndeGroup.POST("decode", ar.symmetryEndeHandler.Decode)
	}

	comprehensiveGroup := ar.engine.Group("/comprehensive/")
	{
		comprehensiveGroup.POST("ip_info", ar.ComprehensiveHandler.IpInfo)
		comprehensiveGroup.POST("domain_map_ip", ar.ComprehensiveHandler.DomainMapIp)
		comprehensiveGroup.POST("image_base64", ar.ComprehensiveHandler.ImageBase64)
	}
}
