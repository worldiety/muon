package json

import (
	"fmt"
	"reflect"
)

const jsonDiscriminator = "jsonDiscriminator"

type Context struct {
	discriminators map[string]reflect.Type
}

func NewContext() *Context {
	return &Context{discriminators: make(map[string]reflect.Type)}
}

func (c *Context) ByType(t reflect.Type) string {
	for k, v := range c.discriminators {
		if v == t {
			return k
		}
	}
	return ""
}

func (c *Context) AddType(name string, t reflect.Type) {
	if !(t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct) {
		panic("invalid type " + t.String() + " expected pointer to struct type")
	}
	if v, ok := c.discriminators[name]; ok && v != t {
		panic("ambiguous type definition for " + name + ": " + t.String() + " vs " + v.String())
	}
	c.discriminators[name] = t
}

func (c *Context) marshal(v interface{}, parent interface{}) {
	var parentMap map[string]interface{}
	var parentSlice []interface{}
	if m, ok := parent.(map[string]interface{}); ok {
		parentMap = m
	} else if s, ok := parent.([]interface{}); ok {
		parentSlice = s
	} else {
		panic("invalid parent " + reflect.TypeOf(parent).String())
	}

	t := reflect.TypeOf(v)
	disc := c.ByType(t)
	if len(disc) > 0 {
		// is a ptr to a struct
		val := reflect.ValueOf(v)
		for i := 0; i < val.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag
			attrName := t.Name()
			if v, ok := tag.Lookup("json"); ok && len(v) > 0 {
				attrName = v
			}
			if attrName == "-" {
				continue
			}
			switch field.Type.Kind() {
			case reflect.String:
				parentMap[attrName] = val.String()
			case reflect.Slice:
				sliceLen := val.Len()
				if sliceLen > 0 {
					slice := make([]interface{}, 0, sliceLen)
					for i := 0; i < sliceLen; i++ {
						c.marshal(val.Index(i).Interface(), slice)
					}
					parentMap[attrName] = slice
				}
			}
		}

	}
	switch t.Kind() {
	case reflect.Ptr:
		t = t.Elem()
		fmt.Println(t.Elem())
	}

}
