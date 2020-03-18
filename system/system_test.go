package system

import "testing"

func TestGenerate(t *testing.T) {
	System(
		Name("Inventory System"),
		Text(`
				The inventory system consists of a login server, an inventory service and a web application.
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
				ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
				dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
				Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
				labore et dolore magna aliquyam erat, sed diam voluptua.
				
				Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
				are discarded.

		   `),
	)
}
