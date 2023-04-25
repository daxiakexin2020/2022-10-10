package jiuqi

import "30thrid_parth/v2/model"

// const base_data_mapping_path = "basedata/anon/data/getBaseDataMapping"
const base_data_mapping_path = ""

type baseDataMapping struct {
	base
	decoding *model.BaseDataMapping
}

func NewBaseDataMapping(decoding *model.BaseDataMapping) *baseDataMapping {
	ndm := &baseDataMapping{decoding: decoding}
	ndm.generateUrl(base_data_mapping_path)
	return ndm
}

func (bdm *baseDataMapping) Send(params interface{}) (*model.BaseDataMapping, error) {
	if err := bdm.send(params, &bdm.decoding, WithMethod(GET)); err != nil {
		return nil, err
	}
	return bdm.decoding, nil
}
