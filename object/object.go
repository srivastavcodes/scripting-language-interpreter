package object

import "fmt"

type ObjectType string

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

func (er *Error) Inspect() string { return "ERROR" + er.Message }
