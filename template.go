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
	UsingSchema(schema string)

	Render() (string, error)
	WithContentVersion(cv string)
	Outputs() Outputs
	Resources() Resources
	Parameters() Parameters
	Variables() Variables
	WithOutputs(func(o Outputs))
	WithParameters(func(p Parameters))
	WithResources(func(r Resources))
	WithVariables(func(v Variables))
}

func New() Template {
	return &template{
		Schema:            "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
		ContentVersion:    defaultContentVersion,
		DefinedOutputs:    NewOutputs(),
		DefinedParameters: NewParameters(),
		DefinedResources:  NewResources(),
		DefinedVariables:  NewVariables(),
	}
}

type template struct {
	APIProfile        string     `json:"apiProfile,omitempty"`
	Schema            string     `json:"$schema"`
	ContentVersion    string     `json:"contentVersion"`
	DefinedResources  Resources  `json:"resources"`
	DefinedOutputs    Outputs    `json:"outputs"`
	DefinedParameters Parameters `json:"parameters,omitempty"`
	DefinedVariables  Variables  `json:"variables"`
}

func (t *template) Outputs() Outputs {
	return t.DefinedOutputs
}

func (t *template) Parameters() Parameters {
	return t.DefinedParameters
}

func (t *template) Resources() Resources {
	return t.DefinedResources
}

func (t *template) Variables() Variables {
	return t.DefinedVariables
}

func (t *template) UsingAPIProfile(v string) {
	t.APIProfile = v
}

func (t *template) UsingSchema(schema string) {
	t.Schema = schema
}

func (t *template) WithContentVersion(cv string) {
	t.ContentVersion = cv
}

func (t *template) WithOutputs(f func(o Outputs)) {
	f(t.DefinedOutputs)
}

func (t *template) WithParameters(f func(p Parameters)) {
	f(t.DefinedParameters)
}

func (t *template) WithResources(f func(r Resources)) {
	f(t.DefinedResources)
}

func (t *template) WithVariables(f func(v Variables)) {
	f(t.DefinedVariables)
}

func (t *template) Render() (string, error) {
	b, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
