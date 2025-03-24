package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	first := strings.ToUpper(string(names[0][0]))
	second := strings.ToUpper(string(names[1][0]))

	return fmt.Sprintf("%s.%s", first, second)
}
