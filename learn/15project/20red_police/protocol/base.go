package protocol

type Base struct {
	BName  string `json:"bname"  mapstructure:"bname"   validate:"required"`
	Cookie string `json:"cookie" mapstructure:"cookie"  validate:"required"`
}
