package document

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func blub(id string, bla int) {

}

type Endpoint struct {
	Name string `param:"path"`
	Id   int    `param:"path"`
}

func (e Endpoint) Apply() (interface{}, error) {
	return nil, nil
}

func TestNew(t *testing.T) {
	value := reflect.ValueOf(blub).Pointer()
	myfunc := runtime.FuncForPC(value)
	fname := myfunc.Name()

	fmt.Println(fname)

	for i := 0; i < reflect.TypeOf(blub).NumIn(); i++ {
		fmt.Println("in " + reflect.TypeOf(blub).In(i).String())
	}

	declareEndpoint("/session/:id/:name",endpoint,)
}

func declareEndpoint(path string,myFunc interface{},params...interface{}){

}

type UserResponse struct{

}

func endpoint(name string, id int)(*UserResponse,error){
	return nil,nil
}


