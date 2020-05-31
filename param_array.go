package armtemplate

type ArrayParameter interface {
	Default(v []interface{}) ArrayParameter
	Min(n uint) ArrayParameter
	Max(n uint) ArrayParameter
}

type arrayParameter struct {
	Type         string        `json:"type"`
	DefaultValue []interface{} `json:"defaultValue,omitempty"`
	MinLength    uint          `json:"minLength,omitempty"`
	MaxLength    uint          `json:"maxLength,omitempty"`
}

func (p *arrayParameter) Default(v []interface{}) ArrayParameter {
	p.DefaultValue = v
	return p
}

func (p *arrayParameter) Min(n uint) ArrayParameter {
	p.MinLength = n
	return p
}

func (p *arrayParameter) Max(n uint) ArrayParameter {
	p.MaxLength = n
	return p
}
