package model

type BaseDataMapping struct {
	Base
	Data struct {
		Id         string `json:"id"`
		TableName  string `json:"tableName"`
		Pagenation struct {
			HasNext  bool `json:"has_next"`
			Offset   int  `json:"offset"`
			Pagesize int  `json:"pagesize"`
		} `json:"pagenation"`
	} `json:"data"`
}

func NewBaseDataMapping() *BaseDataMapping {
	return &BaseDataMapping{}
}

func (bdm *BaseDataMapping) ToPb() {}
