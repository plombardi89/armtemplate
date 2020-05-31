package armtemplate

type BoolParameter interface {
	Default(b bool) BoolParameter
}

type boolParameter struct {
	Type         string `json:"type"`
	DefaultValue bool   `json:"defaultValue"`
}

func (p *boolParameter) Default(b bool) BoolParameter {
	p.DefaultValue = b
	return p
}
