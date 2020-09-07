package lox

type Environment struct {
	enclosing *Environment
	values map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		enclosing: enclosing,
		values: make(map[string]interface{}),
	}
}

// Define creates a new variable
func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

// Assign changes the value of an existing variable
func (e *Environment) Assign(name string, value interface{}) {
	if e.enclosing != nil {
		e.enclosing.Assign(name, value)
		return
	}
	
	if _, ok := e.values[name]; ok {
		e.values[name] = value
	} else {
		panic(&UndefinedVariableError{name})
	}
}

func (e *Environment) Lookup(name string) interface{} {
	if e.enclosing != nil {
		return e.enclosing.Lookup(name)
	}
	
	val, ok := e.values[name]
	if !ok {
		panic(&UndefinedVariableError{name})
	}
	
	return val
}
