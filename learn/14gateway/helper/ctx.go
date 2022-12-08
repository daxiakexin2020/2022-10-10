package helper

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// query
func GetQueryParams(c *gin.Context) map[string]interface{} {
	query := c.Request.URL.Query()
	var queryParams = make(map[string]interface{}, len(query))
	for k := range query {
		queryParams[k] = c.Query(k)
	}
	return queryParams
}

// path  /test:name/:age
func GetPathParams(c *gin.Context) map[string]interface{} {
	params := c.Params
	var pathParams = make(map[string]interface{}, len(params))
	for _, v := range params {
		pathParams[v.Key] = v.Value
	}
	return pathParams
}

// post body->fomr-data 、application/x-www-form-urlencoded
func GetPostFormParams(c *gin.Context) (map[string]interface{}, error) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		if !errors.Is(err, http.ErrNotMultipart) {
			return nil, err
		}
	}
	var postParams = make(map[string]interface{}, len(c.Request.PostForm))
	for k, v := range c.Request.PostForm {
		if len(v) > 1 {
			postParams[k] = v
		} else if len(v) == 1 {
			postParams[k] = v[0]
		}
	}
	return postParams, nil
}

// post body->json格式、text/html...
func GetBody(c *gin.Context) []byte {
	// 读取body数据
	body, err := c.GetRawData()
	if err != nil {
		return nil
	}
	//把读过的字节流重新放到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body
}

func GetCMethod(ctx *gin.Context) (string, bool) {
	cv, cexists := get(ctx, CMETHOD)
	if !cexists {
		return "", false
	}
	str, ok := cv.(string)
	if !ok {
		return "", ok
	}
	return str, true
}

func GetCPath(ctx *gin.Context) (string, bool) {
	cv, cexists := get(ctx, CPATH)
	if !cexists {
		return "", false
	}
	str, ok := cv.(string)
	if !ok {
		return "", ok
	}
	return str, true
}

func GetCApi(ctx *gin.Context) (string, bool) {
	cv, cexists := get(ctx, CAPI)
	if !cexists {
		return "", false
	}
	str, ok := cv.(string)
	if !ok {
		return "", ok
	}
	return str, true
}

func GetCReqData(ctx *gin.Context) (interface{}, bool) {
	return get(ctx, CReqData)
}

func GetVal(ctx *gin.Context, key string) (val any, exists bool) {
	return get(ctx, key)
}

func get(ctx *gin.Context, key string) (val any, exists bool) {
	return ctx.Get(key)
}

func SetCMethod(ctx *gin.Context, val any) {
	set(ctx, CMETHOD, val)
}

func SetCPath(ctx *gin.Context, val any) {
	set(ctx, CPATH, val)
}

func SetCApi(ctx *gin.Context, val any) {
	set(ctx, CAPI, val)
}

func SetCReqData(ctx *gin.Context, val any) {
	set(ctx, CReqData, val)
}

func SetKV(ctx *gin.Context, key string, val any) {
	set(ctx, key, val)
}

func set(ctx *gin.Context, key string, val any) {
	ctx.Set(key, val)
}
