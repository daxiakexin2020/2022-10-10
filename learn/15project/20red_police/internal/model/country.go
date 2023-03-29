package model

type country struct {
	id                 string
	name               string
	architecture_names []string
}

const (
	SovietUnion   = "苏联"
	Cuba          = "古巴"
	Iraq          = "伊拉克"
	Syria         = "叙利亚"
	Libya         = "利比亚"
	United_States = "美国"
	Britain       = "英国"
	France        = "法国"
	Germany       = "德国"
	South_Korea   = "韩国"
)

var gSovietUnion = &country{
	id:                 "1",
	name:               SovietUnion,
	architecture_names: []string{sa_barrack},
}
