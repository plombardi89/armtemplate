package armtemplate

type IntParameter interface {
	Default(n int) IntParameter
}

type intParameter struct {
	Type         string `json:"int"`
	DefaultValue int    `json:"defaultValue"`
	MinValue     int    `json:"minValue"`
	MaxValue     int    `json:"maxValue"`
}

func (p *intParameter) Default(n int) IntParameter {
	p.DefaultValue = n
	return p
}

func (p *intParameter) Min(n int) IntParameter {
	p.MinValue = n
	return p
}

func (p *intParameter) Max(n int) IntParameter {
	p.MaxValue = n
	return p
}
