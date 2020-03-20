package doc

type Document struct {
	Title    string
	Chapters Chapters
}

type PageBreak struct{}
type Br struct{}
type It struct {
	Element interface{}
}

type Element interface {
}

type Chapters []Chapter

type Chapter struct {
	Title    string
	Elements Elements
}

type Elements []Element

type Text string
type Italic string
type Fat string

type NewLine struct{}

func testdsl() {
	mentor := "Torben"
	_ = Document{
		Title: "A Test document",
		Chapters: Chapters{
			{
				Title: "my first chapter",
				Elements: Elements{
					Chapter{
						"a sub chapter", Elements{
							`
									The inventory system consists of a login server, an inventory service and a web application.
									Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
									ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
									dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
									Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
									labore et dolore magna aliquyam erat, sed diam voluptua.
									
									Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
									are discarded.
						
							   `,
							"hel", Fat("lo"), " ", Italic("world"),
							"your mentor is ", Fat(mentor),
						},
					},
				},
			},
		},
	}
}

func (c *Document) NewChapter(string) *Chapter {
	return nil
}

func (c *Chapter) Add(...interface{}) {

}

func Span(str string) string {
	return ""
}

func Em(str string) string {
	return ""
}

func Title(str string) interface{} {
	return nil
}

func NewChapter(string) interface{} {
	return nil
}

func Create(...interface{}) {}

func testDsl2() {
	Create(
		Title("my book"),
		It("hal"), Em("llo"), " world",
		`
		The inventory system consists of a login server, an inventory service and a web application.
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt 
		ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo 
		dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. 
		Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut 
		labore et dolore magna aliquyam erat, sed diam voluptua.
		
		Text is aligned to chars at the left side and interpreted as markdown. Leading and trailing empty lines
		are discarded.
		`,
		NewChapter("my first chapter"),


	)

}

func testDsl3() {
	doc := &Document{}
	doc.Title = "my book"
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

}
