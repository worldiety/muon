package document

type Model struct {
	title    string
	subtitle string
	preface  string
	chapters []*ChapterLevel1
}

func New() *Model {
	return &Model{}
}

func (m *Model) Title() string {
	return m.title
}

func (m *Model) Subtitle() string {
	return m.subtitle
}

func (m *Model) SetTitle(text string) *Model {
	m.title = text
	return m
}
func (m *Model) SetSubtitle(text string) *Model {
	m.subtitle = text
	return m
}

func (m *Model) SetPreface(text string) *Model {
	m.preface = text
	return m
}

func (m *Model) Preface() string {
	return m.preface
}

func (m *Model) AddChapter(title string) *ChapterLevel1 {
	c := &ChapterLevel1{
		parent: m,
		title:  title,
	}
	m.chapters = append(m.chapters, c)
	return c
}

func (m *Model) Chapters() []*ChapterLevel1 {
	return m.chapters
}

type ChapterLevel1 struct {
	parent   *Model
	title    string
	body     []string
	chapters []*ChapterLevel2
}

func (c *ChapterLevel1) Body() []string {
	return c.body
}

func (c *ChapterLevel1) Title() string {
	return c.title
}

func (c *ChapterLevel1) AddText(text string) *ChapterLevel1 {
	c.body = append(c.body, text)
	return c
}

func (c *ChapterLevel1) AddChapter(title string) *ChapterLevel2 {
	n := &ChapterLevel2{
		parent: c,
		title:  title,
	}
	c.chapters = append(c.chapters, n)
	return n
}

func (c *ChapterLevel1) Chapters() []*ChapterLevel2 {
	return c.chapters
}

func (c *ChapterLevel1) Parent() *Model {
	return c.parent
}

type ChapterLevel2 struct {
	parent *ChapterLevel1
	title  string
	body   []string
}

func (c *ChapterLevel2) Title() string {
	return c.title
}

func (c *ChapterLevel2) AddText(text string) *ChapterLevel2 {
	c.body = append(c.body, text)
	return c
}
func (c *ChapterLevel2) Parent() *ChapterLevel1 {
	return c.parent
}

func (c *ChapterLevel2) Body() []string {
	return c.body
}
