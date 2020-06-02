package armtemplate_test

import (
	"testing"

	"github.com/plombardi89/armtemplate"
	"github.com/stretchr/testify/assert"
)

func TestResources_Add(t *testing.T) {
	r := armtemplate.NewResources()
	r.Add(
		armtemplate.Resource{Name: "foo"},
		armtemplate.Resource{Name: "bar"})

	assert.Equal(t, 2, r.Len())
}
