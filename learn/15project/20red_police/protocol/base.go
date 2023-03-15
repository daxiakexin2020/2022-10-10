package protocol

type Header struct {
	Token string `json:"token" mapstructure:"token"  validate:"required"`
	BName string `json:"bname"  mapstructure:"bname"   validate:"required"`
}
