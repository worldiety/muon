package doc

import "reflect"

const typeAttrName = "type"
const WorkspaceType = "workspace"
const DocumentType = "document"
const ChapterType = "chapter"
const AuthorType = "author"
const NewlineType = "newline"
const NewpageType = "newpage"
const ItalicType = "italic"
const BoldType = "bold"
const UnderlineType = "underline"
const CodeType = "code"
const ImageType = "image"
const TOCType = "toc"
const TitlepageType = "titlepage"
const TextType = "text"

func assertObjList(v interface{}) []map[string]interface{} {
	var res []map[string]interface{}
	if slice, ok := v.([]interface{}); ok {
		for _, o := range slice {
			if m, ok := o.(map[string]interface{}); ok {
				res = append(res, m)
			}
		}
	}
	return res
}

func toJson(genericSlice interface{}) []interface{} {
	if genericSlice == nil {
		return nil
	}
	slice := reflect.ValueOf(genericSlice)
	res := make([]interface{}, 0, slice.Len())
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i).Interface()
		res = append(res, item.(Discriminator).toJson())
	}
	return res
}

func fromJson(m map[string]interface{}) Discriminator {
	typeName := m["type"].(string)
	var obj Discriminator
	switch typeName {
	case WorkspaceType:
		obj = &Workspace{}
	case DocumentType:
		obj = &Document{}
	case AuthorType:
		obj = &Author{}
	case ChapterType:
		obj = &Chapter{}
	case TextType:
		obj = &Span{}
	case TOCType:
		obj = TOC()
	case NewlineType:
		obj = Newline()
	case ItalicType:
		obj = Italic()
	case BoldType:
		obj = Bold()
	case UnderlineType:
		obj = Underline()
	case CodeType:
		obj = &Code{}
	case ImageType:
		obj = &Image{}
	case TitlepageType:
		obj = TitlePage()
	case NewpageType:
		obj = Newpage()
	default:
		panic("unknown format type: " + typeName)
	}
	obj.fromJson(m)
	return obj
}
func optString(m map[string]interface{}, key string) string {
	if str, ok := m[key].(string); ok {
		return str
	}
	return ""
}

func optStringSlice(m map[string]interface{}, key string) []string {
	if str, ok := m[key].([]string); ok {
		return str
	}
	return nil
}

func optInt(m map[string]interface{}, key string) int {
	if i, ok := m[key].(int); ok {
		return i
	}
	if i, ok := m[key].(int64); ok {
		return int(i)
	}
	return 0
}

type defaultType struct {
	name string
}

func (d defaultType) Type() string {
	return d.name
}

func (d defaultType) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m[typeAttrName] = d.Type()
	return m
}

func (d defaultType) fromJson(map[string]interface{}) {

}

type defaultBody struct {
	name string
	Body []Discriminator
}

func (d *defaultBody) Type() string {
	return d.name
}

func (d *defaultBody) toJson() map[string]interface{} {
	m := make(map[string]interface{})
	m["body"] = toJson(d.Body)
	return m
}

func (d *defaultBody) fromJson(m map[string]interface{}) {
	d.Body = nil
	for _, obj := range assertObjList(m["body"]) {
		d.Body = append(d.Body, fromJson(obj))
	}
}
