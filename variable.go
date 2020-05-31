package armtemplate

type Copy struct {
	Name  string      `json:"name"`
	Count uint        `json:"count"`
	Input interface{} `json:"input"`
}

type Variables interface {
	// Any accepts any valid type.
	Any(name string, value interface{})

	// Copy is designed to handle "copy" loop variables in Azure templates. The Copy function supports two modes of
	// operation.
	//
	// Reference: https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/copy-variables
	Copy(name string, values []Copy)

	Map(name string, value map[string]interface{})

	// Reference creates a variable that references another variable.
	Reference(name, other string)

	// String creates a string variable.
	String(name, value string)

	// StringArray creates an array of strings as a variable.
	StringArray(name string, values []string)

	// IntArray creates an array of ints as a variable.
	IntArray(name string, values []int)
}

func NewVariables() Variables {
	return make(variables)
}

type variables map[string]interface{}

func (v variables) set(name string, value interface{}) {
	v[name] = value
}

func (v variables) Any(name string, value interface{}) {
	v.set(name, value)
}

func (v variables) Copy(name string, values []Copy) {
	if name == "copy" {
		v.set(name, values)
	} else {
		wrapper := struct {
			Copy []Copy `json:"copy"`
		}{Copy: values}

		v.set(name, wrapper)
	}
}

func (v variables) Map(name string, value map[string]interface{}) {
	v.set(name, value)
}

func (v variables) Reference(name string, other string) {
	v.set(name, RawExpression(VarRef(other)))
}

func (v variables) String(name, value string) {
	v.set(name, value)
}

func (v variables) StringArray(name string, values []string) {
	v.set(name, values)
}

func (v variables) IntArray(name string, values []int) {
	v.set(name, values)
}
