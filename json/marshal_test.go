package json

import (
	"fmt"
	"reflect"
	"testing"
)

type Doc struct {
	Title    string
	Subtitle string
	Chapters []*Chapter
	Opts     *Options `json:"options"`
}

type Options struct {
	Something string
}

type Chapter struct {
	Name string
	Body string
}

func TestContext_Marshal(t *testing.T) {
	doc := &Doc{
		Title:    "my doc",
		Subtitle: "a good book",
		Chapters: []*Chapter{
			{Name: "my chapter", Body: "lorem ipsum"},
		},
	}
	ctx := NewContext()
	ctx.AddType("doc", reflect.TypeOf(&Doc{}))
	ctx.AddType("chap", reflect.TypeOf(&Chapter{}))

	b, err := ctx.Marshal(doc)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
