package symbol

import "sync"

type Table struct {
	store map[string]interface{}
	consts map[string]bool
	mu    sync.RWMutex
}

func NewTable() *Table {
	return &Table{
		store:  make(map[string]interface{}),
		consts: make(map[string]bool),
	}
}

func (t *Table) Get(name string) (interface{}, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	val, ok := t.store[name]
	return val, ok
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

