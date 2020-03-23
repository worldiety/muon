package doc

type Document struct {
	Title    string
	Elements Chapters
}

type PageBreak struct{}
type Br struct{}
type It struct {
	Element interface{}
}

type Element interface {
}

type Chapters []*Chapter

type Chapter struct {
	Title    string
	Elements []interface{}
}


type Text string
type Italic string
type Fat string

type NewLine struct{}

func testdsl() {
	mentor := "Torben"
	_ = Document{
		Title: "A Test document",
		Elements: Chapters{
			{
				Title: "my first chapter",
				Elements: []interface{}{
					Chapter{
						"a sub chapter", []interface{}{
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

func (c *Document) NewChapter(s string) *Chapter {
	chap := &Chapter{
		Title: s,
	}
	c.Elements = append(c.Elements, chap)
	return chap
}

func (c *Chapter) Add(e ...interface{}) {
	c.Elements = append(c.Elements, e...)
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
