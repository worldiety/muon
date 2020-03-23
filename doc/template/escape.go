package template

import (
	"reflect"
	"strings"
)

// EscapeLatex escapes & % $ # _ { } ~ ^ \
func EscapeLatex(str string) string {
	sb := &strings.Builder{}
	for _, r := range str {
		switch r {
		case '&':
			sb.WriteString(`\&`)
		case '%':
			sb.WriteString(`\%`)
		case '$':
			sb.WriteString(`\$`)
		case '#':
			sb.WriteString(`\#`)
		case '_':
			sb.WriteString(`\_`)
		case '{':
			sb.WriteString(`\{`)
		case '}':
			sb.WriteString(`\}`)
		case '~':
			sb.WriteString(`\textasciitilde`)
		case '^':
			sb.WriteString(`\textasciicircum`)
		case '/':
			sb.WriteString(`\textbackslash`)
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func typeOfName(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func isTypeName(i interface{}, name string) bool {
	return typeOfName(i) == name
}
