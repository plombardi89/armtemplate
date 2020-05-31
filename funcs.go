package armtemplate

import (
	"fmt"
)

//func fmtExpression(v ...string) (string, error) {
//	if len(v) == 0 {
//		return "", errors.New("empty expression")
//	}
//
//	return "", nil
//}

func VarRef(name string) string {
	return fmt.Sprintf("variables('%s')", name)
}

func ParameterRef(name string) string {
	return fmt.Sprintf("parameters('%s')", name)
}

func Concat(first string, other ...string) string {
	return fmt.Sprintf("concat(%s, %s)", first, other)
}
