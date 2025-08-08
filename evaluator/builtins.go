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
				return createError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return createError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return createError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}
			return NULL
		},
	},
	"last": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return createError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return createError("argument to `last` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[len(array.Elements)-1]
			}
			return NULL
		},
	},
	"rest": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return createError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return createError("argument to `rest` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)

			length := len(array.Elements)
			if len(array.Elements) > 0 {
				copied := make([]object.Object, length-1)
				copy(copied, array.Elements[1:length])
				return &object.Array{Elements: copied}
			}
			return NULL
		},
	},
	"push": {
		Func: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return createError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return createError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}
			array := args[0].(*object.Array)
			length := len(array.Elements)

			copied := make([]object.Object, length+1)
			copy(copied, array.Elements)

			copied[length] = args[1]
			return &object.Array{Elements: copied}
		},
	},
}
