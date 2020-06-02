package armtemplate

//func Test_outputString(t *testing.T) {
//	testCases := []struct{
//		expected string
//		input    stringOutput
//	}{
//		{
//			expected: testutil.Data(t, "output/string_basic.json"),
//			input:    stringOutput{Type: TypeString, Value: "__string_value__"},
//		},
//		{
//			expected: testutil.Data(t, "output/string_extended.json"),
//			input: stringOutput{
//				Type:      TypeString,
//				Value:     "__string_value__",
//				Condition: "__condition_expression__",
//				Copy:      CopyWithExpressionCount("__input__", "__count_expression__"),
//			},
//		},
//	}
//
//	ja := jsonassert.New(t)
//
//	for _, tc := range testCases {
//		ja.Assertf(testutil.Jsonify(t, tc.input), tc.expected)
//	}
//}
