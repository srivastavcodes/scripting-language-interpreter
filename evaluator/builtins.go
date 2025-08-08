package evaluator

import "Interpreter_in_Go/object"

var builtIns = map[string]*object.BuiltIn{
	"len": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return createError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return createError("argument to `len()` not supported, got %s", args[0].Type())
			}
		},
	},
}
