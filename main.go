package main

import "github.com/killi1812/go-terminal-template/util/cmd"

func main() {
	println("Bokic")

	var str string
	var br int
	cmd.Setup([]cmd.Flag{
		{
			Ptr:     &str,
			Name:    "Test",
			Message: "Test message",
		},
		{
			Ptr:          &br,
			DefaultValue: 1,
			Name:         "Test2",
			Message:      "Test message",
		},
	})
}
