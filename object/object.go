package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (ig *Integer) Inspect() string { return fmt.Sprintf("%d", ig.Value) }

func (ig *Integer) Type() ObjectType { return INTEGER_OBJ }

type Boolean struct {
	Value bool
}

func (bl *Boolean) Inspect() string { return fmt.Sprintf("%t", bl.Value) }

func (bl *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Null struct{}

func (nl *Null) Inspect() string { return "nil" }

func (nl *Null) Type() ObjectType { return NULL_OBJ }
