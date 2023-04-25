package jiuqi

import "30thrid_parth/v2/model"

const token_path = "https://ug.baidu.com/mcp/pc/pcsearch"

type token struct {
	base
	decoding *model.Token
}

func NewToken(decoding *model.Token) *token {
	t := &token{decoding: decoding}
	t.SetUrl(token_path)
	return t
}

func (t *token) Send(params interface{}) (*model.Token, error) {
	if err := t.send(params, &t.decoding); err != nil {
		return nil, err
	}
	return t.decoding, nil
}
