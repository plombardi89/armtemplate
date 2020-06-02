package armtemplate

import (
	"fmt"
)

func VarRef(name string) string {
	return fmt.Sprintf("variables('%s')", name)
}

func ParameterRef(name string) string {
	return fmt.Sprintf("parameters('%s')", name)
}
