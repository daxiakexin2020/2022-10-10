package jiuqi

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

const base_data_mapping_path = "basedata/anon/data/getBaseDataMapping"

func NewBaseDataMapping() *BaseDataMapping {
	return &BaseDataMapping{}
}

func (bdm *BaseDataMapping) Send(param interface{}) (*BaseDataMapping, error) {
	if err := bdm.send(bdm.generateUrl(""), param, &bdm); err != nil {
		return nil, err
	}
	return bdm, nil
}
