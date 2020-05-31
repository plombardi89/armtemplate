package armtemplate_test

import (
	"testing"

	"github.com/kinbiko/jsonassert"
	"github.com/plombardi89/armtemplate"
	"github.com/stretchr/testify/assert"
)

func TestTemplate_Render(t *testing.T) {
	testCases := []struct {
		template func() armtemplate.Template
		expected string
	}{
		{
			expected: testdata(t, "template_new.json"),
			template: armtemplate.New,
		},
		{
			expected: testdata(t, "template_vars_0.json"),
			template: templateVars0,
		},
	}

	ja := jsonassert.New(t)

	for _, tc := range testCases {
		rendered, err := tc.template().Render()
		if assert.NoError(t, err) {
			//fmt.Println(rendered)
			ja.Assertf(rendered, tc.expected)
		}
	}
}

func templateVars0() armtemplate.Template {
	tmpl := armtemplate.New()
	vars := tmpl.Variables()

	vars.String("var1", "myVariable")
	vars.IntArray("var2", []int{1, 2, 3, 4})
	vars.Reference("var3", "var1")
	vars.Map("var4", map[string]interface{}{
		"property1": "value1",
		"property2": "value2",
	})
	vars.Copy("var5", []armtemplate.Copy{
		{
			Name:  "disks",
			Count: 5,
			Input: map[string]interface{}{
				"name":       "[concat('myDataDisk', copyIndex('disks', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('disks')]",
			},
		},
		{
			Name:  "diskNames",
			Count: 5,
			Input: "[concat('myDataDisk', copyIndex('diskNames', 1))]",
		},
	})

	vars.Copy("copy", []armtemplate.Copy{
		{
			Name:  "var6",
			Count: 5,
			Input: map[string]interface{}{
				"name":       "[concat('oneDataDisk', copyIndex('var6', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('var6')]",
			},
		},
		{
			Name:  "var7",
			Count: 3,
			Input: map[string]interface{}{
				"name":       "[concat('twoDataDisk', copyIndex('var7', 1))]",
				"diskSizeGB": "1",
				"diskIndex":  "[copyIndex('var7')]",
			},
		},
		{
			Name:  "var8",
			Count: 4,
			Input: "[concat('stringValue', copyIndex('var8'))]",
		},
		{
			Name:  "var9",
			Count: 4,
			Input: "[copyIndex('var9')]",
		},
	})

	return tmpl
}
