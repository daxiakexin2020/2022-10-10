package error

import "errors"

const (
	SUCCESS_CODE    = 0
	SERVER_ERR_CODE = 500
	PARAM_ERR_CODE  = 10001

	DB_ERR_CODE      = 20001
	BUSSINE_ERR_CODE = 30001

	UPLOAD_ERR_CODE = 40001
)

var (
	NO_SUPPORT_FILE_TYPE     = errors.New("no support this file type")
	OVERFLOW_UPLOAD_MAX_SIZE = errors.New("overflow uplaod max size")
)
