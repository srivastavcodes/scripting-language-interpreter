package object

import "fmt"

type ObjectType string

const (
	COLOR_RED   = "\033[31m"
	COLOR_RESET = "\033[0m"
	// COLOR_GREEN = "\033[32m"
)

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (ig *Integer) Type() ObjectType { return INTEGER_OBJ }

func (ig *Integer) Inspect() string { return fmt.Sprintf("%d", ig.Value) }

type Boolean struct {
	Value bool
}

func (bl *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (bl *Boolean) Inspect() string { return fmt.Sprintf("%t", bl.Value) }

type Null struct{}

func (nl *Null) Inspect() string { return "nil" }

func (nl *Null) Type() ObjectType { return NULL_OBJ }

type Return struct {
	Value Object
}

func (rv *Return) Type() ObjectType { return RETURN_VALUE_OBJ }

func (rv *Return) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (er *Error) Type() ObjectType { return ERROR_OBJ }

func (er *Error) Inspect() string {
	return fmt.Sprintf("%sERROR::%s %s", COLOR_RED, COLOR_RESET, er.Message)
}

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object)}
}

func (env *Environment) Get(name string) (Object, bool) {
	obj, ok := env.store[name]
	return obj, ok
}

func (env *Environment) Set(name string, val Object) Object {
	env.store[name] = val
	return val
}
