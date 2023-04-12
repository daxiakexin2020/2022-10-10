package error

import (
	"22go_redis/server/construct"
	"errors"
)

var (
	HASH_TYPE_ERROR   = errors.New("type is error，not：" + construct.HASH)
	LIST_TYPE_ERROR   = errors.New("type is error，not：" + construct.LIST)
	SET_TYPE_ERROR    = errors.New("type is error，not：" + construct.SET)
	ZSET_TYPE_ERROR   = errors.New("type is error，not：" + construct.ZSET)
	STRING_TYPE_ERROR = errors.New("type is error，not：" + construct.STRING)
)
