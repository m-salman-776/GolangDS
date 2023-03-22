package main

import (
	"fmt"
)

type Animal interface {
	Walk()
	Breath()
}

type Lion struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func (l *Lion) Walk() {
	fmt.Println(l.Name, "is Waling")
}
func (l *Lion) Breath() {
	fmt.Println(l.Name, "is Breathing")
}

func main() {
	//var lion Animal
	//lion = Lion{Name: "Some name", Age: 123}
	lion := Lion{
		Name: "Some lion",
		Age:  54,
	}
	encoded, err := JSONEncode(lion)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(encoded))
	//types := reflect.TypeOf(lion)
	//values := reflect.ValueOf(lion)
	//fmt.Println(types, "values->", values)
}

// interface variable => (interface Type , interface Value) => (Animal,----)
// Interface Value => (Underlying type(reflect.TypeOf()) , Underlying value reflect.ValuesOf()) => (Lion , {Name : "xyx})
