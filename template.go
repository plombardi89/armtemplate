package armtemplate

import (
	"encoding/json"
	"fmt"
)

const (
	// Null is an explicit null value for ARM templates. Generally it should be possible to use the Go nil to achieve
	// the same result as Null. However it is useful when a JSON struct has an 'omitempty' option on a json tag and
	// therefore would exclude the null value.
	Null = "[json('null')]"

	defaultContentVersion = "1.0.0.0"
)

// RawExpression constructs an ARM template expression by wrapping the input with expression marking brackets.
func RawExpression(expression string) string {
	return fmt.Sprintf("[%s]", expression)
}

type Template interface {
	UsingAPIProfile(v string)
	Render() (string, error)
	WithContentVersion(cv string)

	Parameters() Parameters

	// WithVariables replaces the currently defined variables with a different set of variables.
	WithVariables(v Variables)

	// Variables returns the defined variables for the template.
	Variables() Variables
}

func New() Template {
	resources := make([]interface{}, 0)

	return &template{
		Schema:            "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		ContentVersion:    defaultContentVersion,
		DefinedVariables:  NewVariables(),
		DefinedParameters: NewParameters(),
		Resources:         resources,
	}
}

type template struct {
	APIProfile        string        `json:"apiProfile,omitempty"`
	Schema            string        `json:"$schema"`
	ContentVersion    string        `json:"contentVersion"`
	Resources         []interface{} `json:"resources"`
	DefinedParameters Parameters    `json:"parameters,omitempty"`
	DefinedVariables  Variables     `json:"variables"`
}

func (t *template) Parameters() Parameters {
	return t.DefinedParameters
}

func (t *template) Variables() Variables {
	return t.DefinedVariables
}

func (t *template) UsingAPIProfile(v string) {
	t.APIProfile = v
}

func (t *template) WithContentVersion(cv string) {
	t.ContentVersion = cv
}

func (t *template) WithVariables(v Variables) {
	t.DefinedVariables = v
}

func (t *template) Render() (string, error) {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
