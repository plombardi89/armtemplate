package armtemplate

type ObjectParameter interface {
	Default(v interface{}) ObjectParameter
}

type objectParameter struct {
	Type         string      `json:"type"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
}

func (p *objectParameter) Default(v interface{}) ObjectParameter {
	p.DefaultValue = v
	return p
}
