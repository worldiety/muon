package template

import (
	"fmt"
	"strings"
)

type File struct {
	parent *Project
	b      *strings.Builder
}

func (f *File) Printf(format string, a ...interface{}) *File {
	f.b.WriteString(fmt.Sprintf(format, a))
	return f
}

func (f *File) Println(a ...interface{}) *File {
	f.b.WriteString(fmt.Sprintln(a))
	return f
}
