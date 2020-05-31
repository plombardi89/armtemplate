package armtemplate_test

import (
	"fmt"
	"testing"

	"github.com/plombardi89/armtemplate"
)

func TestNewStringParameter(t *testing.T) {
	tmpl := armtemplate.New()
	params := tmpl.Parameters()

	sp := params.String("foo")

	params.Int("bar")

	sp.Default("bar").Allowed([]string{"baz", "bot"}).Min(3).Max(6)

	r, _ := tmpl.Render()
	fmt.Println(r)
}
