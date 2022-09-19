package main

import (
	"fmt"
	"reflect"
)

var typeRegistry = make(map[string]reflect.Type)

func registerType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	typeRegistry[t.PkgPath()+"."+t.Name()] = t
}

type MyString string
type myString string
type Screen struct {
	Width float32
}

func init() {
	registerType((*MyString)(nil))
	registerType((*myString)(nil))
	registerType((*Screen)(nil))
	// ...
}

func makeInstance(name string) interface{} {
	return reflect.New(typeRegistry[name]).Elem().Interface()
}

func main() {
	for k := range typeRegistry {
		fmt.Println(k)
	}
	fmt.Printf("%T\n", makeInstance("main.MyString"))
	fmt.Printf("%T\n", makeInstance("main.myString"))
	fmt.Printf("%#v\n", makeInstance("main.Screen"))
}
