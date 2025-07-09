// Package cmd used for registering comand line arguments
package cmd

import (
	"flag"
	"reflect"
)

// Flag struct
type Flag struct {
	Ptr          any
	Name         string
	DefaultValue any
	Message      string
}

func Setup(flags []Flag) {
	for _, f := range flags {
		RegisterFlag(f)
	}

	flag.Parse()
}

func RegisterFlag(f Flag) {
	ftype := reflect.TypeOf(f.Ptr)

	switch ftype.String() {
	case "*string":
		ptr, ok := f.Ptr.(*string)
		if !ok {
			panic("Not a string pointer")
		}

		var dvalue string
		if f.DefaultValue != nil {
			dvalue, ok = f.DefaultValue.(string)
			if !ok {
				panic("Not a string default value")
			}
		}
		flag.StringVar(ptr, f.Name, dvalue, f.Message)

	case "*int":
		ptr, ok := f.Ptr.(*int)
		if !ok {
			panic("Not a int pointer")
		}
		var dvalue int
		if f.DefaultValue != nil {
			dvalue, ok = f.DefaultValue.(int)
			if !ok {
				panic("Not a int default value")
			}
		}
		flag.IntVar(ptr, f.Name, dvalue, f.Message)

	case "*bool":
		ptr, ok := f.Ptr.(*bool)
		if !ok {
			panic("Not a bool pointer")
		}
		var dvalue bool
		if f.DefaultValue != nil {
			dvalue, ok = f.DefaultValue.(bool)
			if !ok {
				panic("Not a bool default value")
			}
		}
		flag.BoolVar(ptr, f.Name, dvalue, f.Message)

	case "*float64":
		ptr, ok := f.Ptr.(*float64)
		if !ok {
			panic("Not a float64 pointer")
		}
		var dvalue float64
		if f.DefaultValue != nil {
			dvalue, ok = f.DefaultValue.(float64)
			if !ok {
				panic("Not a float64 default value")
			}
		}
		flag.Float64Var(ptr, f.Name, dvalue, f.Message)

	}
}
