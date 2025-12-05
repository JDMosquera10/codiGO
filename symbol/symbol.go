package symbol

import "sync"

type Table struct {
	store  map[string]interface{}
	consts map[string]bool
	parent *Table // Scope padre para búsqueda recursiva
	mu     sync.RWMutex
}

func NewTable() *Table {
	return &Table{
		store:  make(map[string]interface{}),
		consts: make(map[string]bool),
		parent: nil,
	}
}

// NewTableWithParent crea una nueva tabla con un scope padre
func NewTableWithParent(parent *Table) *Table {
	return &Table{
		store:  make(map[string]interface{}),
		consts: make(map[string]bool),
		parent: parent,
	}
}

func (t *Table) Get(name string) (interface{}, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	val, ok := t.store[name]
	if ok {
		return val, ok
	}
	// Si no se encuentra, buscar en el scope padre
	if t.parent != nil {
		return t.parent.Get(name)
	}
	return nil, false
}

func (t *Table) Set(name string, value interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.consts[name] {
		// No permitir reasignar constantes
		return
	}
	t.store[name] = value
}

func (t *Table) SetConst(name string, value interface{}) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.store[name] = value
	t.consts[name] = true
}

func (t *Table) Has(name string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	_, ok := t.store[name]
	return ok
}

