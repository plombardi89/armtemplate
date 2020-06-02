package armtemplate

//func CopyWithIntCount(input string, count int) *CopyOutput {
//	return &CopyOutput{Count: count, Input: input}
//}
//
//func CopyWithExpressionCount(input, expression string) *CopyOutput {
//	return &CopyOutput{Count: expression, Input: input}
//}
//
//type CopyOutput struct {
//	Count interface{} `json:"count"`
//	Input string      `json:"input"`
//}

type Output struct {
	Type      string      `json:"type"`
	Condition string      `json:"condition,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Copy      interface{} `json:"copy,omitempty"`
}

type OutputOption func(o *Output)

type Outputs interface {
	//Array(name string) ArrayParameter
	//Bool(name string) BoolParameter
	//Int(name string) IntParameter
	//Object(name string) ObjectParameter
	String(name string, opts ...OutputOption)
	Remove(name string)
	WithValue(v string) OutputOption
	WithCondition(c string) OutputOption
}

func NewOutputs() Outputs {
	return make(outputs)
}

type outputs map[string]interface{}

func (o outputs) addOutput(name, valueType string, opts ...OutputOption) {
	output := &Output{Type: valueType}

	for _, opt := range opts {
		opt(output)
	}

	o[name] = output
}

func (o outputs) String(name string, opts ...OutputOption) {
	o.addOutput(name, TypeString, opts...)
}

func (o outputs) WithValue(v string) OutputOption {
	return func(o *Output) {
		o.Value = v
	}
}

func (o outputs) WithCondition(c string) OutputOption {
	return func(o *Output) {
		o.Condition = c
	}
}

func (o outputs) Remove(name string) {
	delete(o, name)
}
