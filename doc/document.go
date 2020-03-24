package doc

// A Discriminator returns a unique type name
type Discriminator interface {
	Type() string
	toJson() map[string]interface{}
	fromJson(map[string]interface{})
}

// A Workspace contains all resources for different projects, groups whatever.
type Workspace struct {
	Version   string
	Title     string
	Resources []Discriminator
}

func (w *Workspace) NewDocument() *Document {
	doc := &Document{}
	w.Resources = append(w.Resources, doc)
	return doc
}

func (w *Workspace) Type() string {
	return WorkspaceType
}

func (w *Workspace) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = w.Type()
	m["title"] = w.Title
	m["version"] = w.Version
	m["resources"] = toJson(w.Resources)
	return m
}

func (w *Workspace) fromJson(m map[string]interface{}) {
	w.Title = m["title"].(string)
	w.Version = m["version"].(string)
	w.Resources = nil
	for _, obj := range assertObjList(m["resources"]) {
		w.Resources = append(w.Resources, fromJson(obj))
	}
}

// A Document contains a markup mixture related to typesetting a book, article or webpage, especially for
// technical content.
type Document struct {
	Title   string
	Authors []*Author
	Body    []Discriminator
}

func (c *Document) NewChapter(s string) *Chapter {
	chap := &Chapter{
		Title: s,
		Level: 0,
	}
	c.Body = append(c.Body, chap)
	return chap
}

func (c *Document) Add(e ...Discriminator) *Document {
	c.Body = append(c.Body, e...)
	return c
}

func (c *Document) Type() string {
	return DocumentType
}

func (c *Document) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = c.Type()
	m["title"] = c.Title
	m["authors"] = toJson(c.Authors)
	m["body"] = toJson(c.Body)
	return m
}

func (c *Document) fromJson(m map[string]interface{}) {
	c.Title = m["title"].(string)
	c.Authors = nil
	for _, obj := range assertObjList(m["authors"]) {
		c.Authors = append(c.Authors, fromJson(obj).(*Author))
	}
	c.Body = nil
	for _, obj := range assertObjList(m["body"]) {
		c.Body = append(c.Body, fromJson(obj))
	}
}

// Author describes a user who has written something in the document
type Author struct {
	Firstname string
	Lastname  string
	EMail     string
}

func (a *Author) Type() string {
	return AuthorType
}

func (a *Author) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = a.Type()
	m["firstname"] = a.Firstname
	m["lastname"] = a.Lastname
	m["email"] = a.EMail
	return m
}

func (a *Author) fromJson(m map[string]interface{}) {
	a.Firstname = m["firstname"].(string)
	a.Lastname = m["lastname"].(string)
	a.EMail = m["email"].(string)
}

// A Chapter allows the hierarchical titled grouping. Better to keep the level consistent with the hierarchy.
type Chapter struct {
	Title string
	Level int // start by 0 and keep consistent
	Body  []Discriminator
}

func (c *Chapter) Add(e ...Discriminator) *Chapter {
	c.Body = append(c.Body, e...)
	return c
}

func (c *Chapter) NewChapter(title string) *Chapter {
	chap := &Chapter{
		Title: title,
		Level: c.Level + 1,
	}
	c.Body = append(c.Body, chap)
	return chap
}

func (c *Chapter) Text(str string) *Chapter {
	c.Add(&Span{Value: str})
	return c
}

func (c *Chapter) Type() string {
	return ChapterType
}

func (c *Chapter) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = c.Type()
	m["title"] = c.Title
	m["level"] = c.Level
	m["body"] = toJson(c.Body)
	return m
}

func (c *Chapter) fromJson(m map[string]interface{}) {
	c.Title = optString(m, "title")
	c.Level = optInt(m, "level")
	c.Body = nil
	for _, obj := range assertObjList(m["body"]) {
		c.Body = append(c.Body, fromJson(obj))
	}
}

// Newpage creates a new page element
func Newpage() Discriminator {
	return defaultType{name: NewpageType}
}

// Newline creates a new line element
func Newline() Discriminator {
	return defaultType{name: NewlineType}
}

// TOC creates a table of contents based on chapters and their according levels
func TOC() Discriminator {
	return defaultType{name: TOCType}
}

// Italic creates a new body group for cursive typesetting
func Italic(body ...Discriminator) *defaultBody {
	return &defaultBody{name: ItalicType, Body: body}
}

// Bold creates a new body group for fat typesetting
func Bold(body ...Discriminator) *defaultBody {
	return &defaultBody{name: BoldType, Body: body}
}

// Underline creates a new body group for fat typesetting
func Underline(body ...Discriminator) *defaultBody {
	return &defaultBody{name: UnderlineType, Body: body}
}

// A TitlePage is a specially formatted page with a certain meaning.
// The interpretation of the body depends largely on the actual template
// and may put everything or nothing or just the first text.
func TitlePage(body ...Discriminator) *defaultBody {
	return &defaultBody{name: TitlepageType, Body: body}
}

type Span struct {
	Value string
}

func (t *Span) String() string {
	return t.Value
}

func (t *Span) Type() string {
	return TextType
}

func (t *Span) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = t.Type()
	m["value"] = t.Value
	return m
}

func (t *Span) fromJson(m map[string]interface{}) {
	t.Value = optString(m, "value")
}

func Text(str string) *Span {
	return &Span{str}
}

// A Code element contains a bunch of lines and a type hint
type Code struct {
	Hint  string //
	Lines []string
}

func (c *Code) Type() string {
	return CodeType
}

func (c *Code) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = c.Type()
	m["hint"] = c.Hint
	m["lines"] = c.Lines
	return m
}

func (c *Code) fromJson(m map[string]interface{}) {
	c.Hint = optString(m, "hint")
	c.Lines = optStringSlice(m, "lines")
}

// An Image element contains a reference (filename) to a usually local image
type Image struct {
	Src    string
	Width  string
	Height string
}

func (c *Image) Type() string {
	return ImageType
}

func (c *Image) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = c.Type()
	m["src"] = c.Src
	m["width"] = c.Width
	m["height"] = c.Height
	return m
}

func (c *Image) fromJson(m map[string]interface{}) {
	c.Src = optString(m, "src")
	c.Width = optString(m, "width")
	c.Height = optString(m, "height")
}
