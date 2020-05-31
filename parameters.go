package armtemplate

const (
	ParamTypeArray        = "array"
	ParamTypeBool         = "bool"
	ParamTypeInt          = "int"
	ParamTypeObject       = "object"
	ParamTypeString       = "string"
	ParamTypeSecureString = "securestring"
)

type Parameters interface {
	Array(name string) ArrayParameter
	Bool(name string) BoolParameter
	Int(name string) IntParameter
	Object(name string) ObjectParameter
	String(name string) StringParameter
}

func NewParameters() Parameters {
	return make(parameters)
}

type parameters map[string]interface{}

func (p parameters) Array(name string) ArrayParameter {
	bp := &arrayParameter{Type: ParamTypeArray}
	p[name] = bp

	return bp
}

func (p parameters) Bool(name string) BoolParameter {
	bp := &boolParameter{Type: ParamTypeBool}
	p[name] = bp

	return bp
}

func (p parameters) Int(name string) IntParameter {
	ip := &intParameter{Type: ParamTypeInt}
	p[name] = ip

	return ip
}

func (p parameters) Object(name string) ObjectParameter {
	bp := &objectParameter{Type: ParamTypeObject}
	p[name] = bp

	return bp
}

func (p parameters) String(name string) StringParameter {
	sp := &stringParameter{Type: ParamTypeString}
	p[name] = sp

	return sp
}

func (p parameters) SecureString(name string) StringParameter {
	sp := &stringParameter{Type: ParamTypeSecureString}
	p[name] = sp

	return sp
}
