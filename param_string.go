package armtemplate

type StringParameter interface {
	Default(s string) StringParameter
	Description(s string) StringParameter
	Allowed(values []string) StringParameter
	Min(min uint) StringParameter
	Max(max uint) StringParameter
}

type stringParameter struct {
	Type          string    `json:"type"`
	DefaultValue  string    `json:"defaultValue,omitempty"`
	AllowedValues []string  `json:"allowedValues,omitempty"`
	MinLength     uint      `json:"minLength,omitempty"`
	MaxLength     uint      `json:"maxLength,omitempty"`
	Metadata      *Metadata `json:"metadata,omitempty"`
}

func (p *stringParameter) Description(s string) StringParameter {
	if p.Metadata == nil {
		p.Metadata = &Metadata{}
	}

	p.Metadata.Description = s

	return p
}

func (p *stringParameter) Default(s string) StringParameter {
	p.DefaultValue = s
	return p
}

func (p *stringParameter) Allowed(values []string) StringParameter {
	p.AllowedValues = values
	return p
}

func (p *stringParameter) Min(min uint) StringParameter {
	p.MinLength = min
	return p
}

func (p *stringParameter) Max(max uint) StringParameter {
	p.MaxLength = max
	return p
}
