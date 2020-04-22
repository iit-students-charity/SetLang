package gen

import "fmt"

const (
	prog = `package main

import (
	"fmt"
)

type Set struct {
	elems map[string]Value
}

type Value struct {
	val string
}

func NewValue(v string) Value {
	return Value{val: v}
}

func NewSet(vals ...Value) Set {
	elems := make(map[string]Value)
	for _, v := range vals {
		elems[v.val] = v
	}
	return Set{elems: elems}
}

func (v Value) get() string {
	return v.val
}

func (v Value) PRINT() string {
	fmt.Println(v.val)
	return v.val
}

func (s Set) PRINT() string {
	fmt.Print("{")
	for _, v := range s.elems {
		fmt.Printf("%s, ", v.val)
	}
	fmt.Println("}")
	return ""
}

func (s Set) ADDELEM(val Value) string {
	s.elems[val.get()] = val
	return ""
}

func dump(a interface{})  {
	
}

`
)

type Set struct {
	elems map[string]Value
}

type Value struct {
	val string
}

func NewValue(v string) Value {
	return Value{val: v}
}

func NewSet(vals ...Value) Set {
	elems := make(map[string]Value)
	for _, v := range vals {
		elems[v.val] = v
	}
	return Set{elems: elems}
}

func (v Value) get() string {
	return v.val
}

func (v Value) PRINT() string {
	fmt.Println(v.val)
	return v.val
}

func (s Set) PRINT() string {
	fmt.Print("{")
	for _, v := range s.elems {
		fmt.Printf("%s, ", v.val)
	}
	fmt.Println("}")
	return ""
}

func (s Set) ADDELEM(val Value) string {
	s.elems[val.get()] = val
	return ""
}

func dump(a interface{}) {

}