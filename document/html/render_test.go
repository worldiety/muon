package html

import (
	"bytes"
	. "github.com/worldiety/muon/document"
	"io/ioutil"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	b := &bytes.Buffer{}

	err := Write(createDoc(), b)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("test.html", b.Bytes(), os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func createDoc() *Model {
	return New().
		SetTitle("Hello World").
		SetSubtitle("A system documentation").
		SetPreface("Dedicated to all who just wants documentations fast.").
		AddChapter("my first chapter").
		AddText("The first chapter is about stuff.").
		AddText(`
				The inventory system consists of a login server, an inventory service and a web application.
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
				ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
				dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
				labore et dolore magna aliquyam erat, sed diam voluptua.
				
				Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
				are discarded.

		   `).
		AddChapter("my sub chapter and heading").
		AddText("my sub chapter is about sub stuff.").Parent().Parent().
		AddChapter("The second chapter").
		AddText("second chapter stuff.").
		AddText(`
				The inventory system consists of a login server, an inventory service and a web application.
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
				ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
				dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
				labore et dolore magna aliquyam erat, sed diam voluptua.
				
				Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
				are discarded.

		   `).
		AddText(`
				The inventory system consists of a login server, an inventory service and a web application.
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
				ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
				dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
				labore et dolore magna aliquyam erat, sed diam voluptua.
				
				Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
				are discarded.

		   `).
		AddText(`
				The inventory system consists of a login server, an inventory service and a web application.
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
				ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
				dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
				labore et dolore magna aliquyam erat, sed diam voluptua.
				
				Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
				are discarded.

		   `).
		Parent()
}
