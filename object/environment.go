package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Object), outer: nil}
}

func (env *Environment) Get(name string) (Object, bool) {
	ob, ok := env.store[name]
	if !ok && env.outer != nil {
		ob, ok = env.outer.Get(name)
	}
	return ob, ok
}

func (env *Environment) Set(name string, val Object) Object {
	env.store[name] = val
	return val
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
