package items

import (
	"22go_redis/server/construct"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Gredis struct {
	dbIndex    dbindex
	maxSize    int64
	maxSizeStr string
	usedSize   int64
	*construct.CGORedis
}

type dbindex uint

type Option func(g *Gredis)

const (
	DBIndex_0 dbindex = iota
	DBIndex_1
	DBIndex_2
	DBIndex_3
	DBIndex_4
	DBIndex_5
	DBIndex_6
	DBIndex_7
	DBIndex_8
	DBIndex_9
	DBIndex_10
	DBIndex_11
	DBIndex_12
	DBIndex_13
	DBIndex_14
	DBIndex_15
	DBIndex_16
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

const (
	UNIT_B  = "B"
	UNIT_KB = "KB"
	UNIT_MB = "MB"
	UNIT_GB = "GB"
	UNIT_TB = "TB"
	UNIT_PB = "PB"
)

const (
	defaultDBIndex    = DBIndex_0
	defaultMaxSizeStr = "500MB"
)

var GORedis *Gredis

func WithDBIndex(dbx dbindex) Option {
	return func(g *Gredis) {
		if dbx >= DBIndex_0 && dbx <= DBIndex_16 {
			g.dbIndex = dbx
		}
	}
}

func (gr *Gredis) apply(option ...Option) {
	for _, opt := range option {
		opt(gr)
	}
}

func init() {
	gr, err := NewGredis(defaultMaxSizeStr)
	if err == nil {
		GORedis = gr
	}
}

func NewGredis(maxSizeStr string, option ...Option) (*Gredis, error) {
	maxSize, maxSizeUnit, err := Parse(maxSizeStr)
	if err != nil {
		return nil, err
	}
	gr := &Gredis{
		dbIndex:    defaultDBIndex,
		maxSize:    maxSize,
		maxSizeStr: maxSizeUnit,
		CGORedis:   construct.NewCGORedis(),
	}
	gr.apply(option...)
	return gr, nil
}

func (gr *Gredis) isData(key string) (construct.VInterface, bool) {
	return gr.CGOIsData(key)
}

func Parse(str string) (int64, string, error) {
	numRe := regexp.MustCompile("[0-9]+")
	unitRe := regexp.MustCompile("[a-zA-Z]+")
	num, err := strconv.ParseInt(numRe.FindString(str), 10, 64)
	var size int64
	var unitStr string
	if err != nil {
		return size, unitStr, err
	}
	unit := strings.ToUpper(unitRe.FindString(str))
	switch unit {
	case UNIT_B:
		size = num * B
	case UNIT_KB:
		size = num * KB
	case UNIT_MB:
		size = num * MB
	case UNIT_GB:
		size = num * GB
	case UNIT_TB:
		size = num * TB
	case UNIT_PB:
		size = num * PB
	default:
		return size, unitStr, errors.New("support unit：" + UNIT_B + "," + UNIT_KB + "," + UNIT_MB + "," + UNIT_GB + "," + UNIT_TB + "," + UNIT_PB)
	}
	return size, strconv.FormatInt(num, 10) + unit, nil
}
