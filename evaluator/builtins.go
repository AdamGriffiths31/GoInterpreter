package evaluator

import (
	"interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GetBuiltinByName("len"),
	"push":  object.GetBuiltinByName("push"),
	"print": object.GetBuiltinByName("print"),
	"first": object.GetBuiltinByName("first"),
	"last":  object.GetBuiltinByName("last"),
	"rest":  object.GetBuiltinByName("rest"),
}
