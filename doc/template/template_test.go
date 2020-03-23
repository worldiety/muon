package template

import (
	. "github.com/worldiety/muon/doc"
	"testing"
)

func TestOpen(t *testing.T) {
	prj, err := Open("/Users/tschinke/tmp/muondoc-wdy-book-01-latex", "/Users/tschinke/tmp/muondoc-wdy-book-01-latex/build")
	if err != nil {
		t.Fatal(err)
	}

	err = prj.Build(createDoc())
	if err != nil {
		t.Fatal(err)
	}
}

func createDoc() *Document {
	doc := &Document{}
	doc.Title = `my book & % $ # _ { } ~ ^ \`
	chap := doc.NewChapter("my first chapter")
	chap.Add(
		`
		The inventory system consists of a login server, an inventory service and a web application.
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
		ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
		dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
		labore et dolore magna aliquyam erat, sed diam voluptua.
		
		Text is aligned to chars at the left side. Empty lines are ignored.
		`,
		Br{},
		It{"hal"}, Em("llo"), " world",
	)

	return doc
}
