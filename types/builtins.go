package types

import (
	"fmt"
)

var builtins = map[string]*Builtin{
	"print": {
		Fun: print,
	},
	// ;)
	"generatePassword": {
		Fun: generatePassword,
	},
}

func print(args ...Object) Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}

func generatePassword(args ...Object) Object {
	return &String{Value: "password1234"}
}
