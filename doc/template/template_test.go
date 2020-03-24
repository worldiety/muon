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
	ws := &Workspace{}
	doc := ws.NewDocument()
	doc.Add(TitlePage(Text("my technical book"), Text("a subtitle")), TOC())
	chap := doc.NewChapter("my first chapter")

	chap.Text(
		`
		The inventory system consists of a login server, an inventory service and a web application.
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
		ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
		dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
		labore et dolore magna aliquyam erat, sed diam voluptua.
		
		Span is aligned to chars at the left side. Empty lines are ignored.`)
	chap.Text("wtf")
	chap.Add(Newline())
	chap.Add(Text("hello "), Italic(Text("worl"), Bold(Underline(Text("d")))), Newline())
	chap.Add(Bold(Italic(Text(`ugly chars: & % $ # _ { } ~ ^ \`))), Newline())

	sub := chap.NewChapter("a section")
	sub.Text("This is a section within a chapter.")

	subsub := sub.NewChapter("a subsection")
	subsub.Text("This is another text but in a subsubsection.")

	chap = doc.NewChapter("another main chapter")
	chap.Text("typesetting test.")

	return doc
}
