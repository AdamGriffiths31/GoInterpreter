package compiler

import "testing"

func TestDefine(t *testing.T) {
	expected := map[string]Symbol{
		"a": {Name: "a", Scope: GlobalScope, Index: 0},
		"b": {Name: "b", Scope: GlobalScope, Index: 1},
		"c": {Name: "c", Scope: LocalScope, Index: 0},
		"d": {Name: "d", Scope: LocalScope, Index: 1},
		"e": {Name: "e", Scope: LocalScope, Index: 0},
		"f": {Name: "f", Scope: LocalScope, Index: 1},
	}

	global := NewSymbolTable()

	a := global.Define("a")
	if a != expected["a"] {
		t.Errorf("a = %v, expected %v", a, expected["a"])
	}

	b := global.Define("b")
	if b != expected["b"] {
		t.Errorf("b = %v, expected %v", b, expected["b"])
	}

	firstLocal := NewEnclosedSymbolTable(global)

	c := firstLocal.Define("c")
	if c != expected["c"] {
		t.Errorf("c = %v, expected %v", c, expected["c"])
	}

	d := firstLocal.Define("d")
	if d != expected["d"] {
		t.Errorf("d = %v, expected %v", d, expected["d"])
	}

	secondLocal := NewEnclosedSymbolTable(firstLocal)

	e := secondLocal.Define("e")
	if e != expected["e"] {
		t.Errorf("e = %v, expected %v", e, expected["e"])
	}

	f := secondLocal.Define("f")
	if f != expected["f"] {
		t.Errorf("f = %v, expected %v", f, expected["f"])
	}
}

func TestResolveGlobal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	expected := map[string]Symbol{
		"a": {Name: "a", Scope: GlobalScope, Index: 0},
		"b": {Name: "b", Scope: GlobalScope, Index: 1},
	}

	for _, sym := range global.store {
		result, ok := global.Resolve(sym.Name)
		if !ok {
			t.Errorf("name %s not resolvable", sym.Name)
			continue
		}
		if result != expected[sym.Name] {
			t.Errorf("name %s = %v, expected %v", sym.Name, result, expected[sym.Name])
		}
	}

}

func TestResolveLocal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	local := NewEnclosedSymbolTable(global)
	local.Define("c")
	local.Define("d")

	expected := map[string]Symbol{
		"a": {Name: "a", Scope: GlobalScope, Index: 0},
		"b": {Name: "b", Scope: GlobalScope, Index: 1},
		"c": {Name: "c", Scope: LocalScope, Index: 0},
		"d": {Name: "d", Scope: LocalScope, Index: 1},
	}

	for _, sym := range local.store {
		result, ok := local.Resolve(sym.Name)
		if !ok {
			t.Errorf("name %s not resolvable", sym.Name)
			continue
		}
		if result != expected[sym.Name] {
			t.Errorf("name %s = %v, expected %v", sym.Name, result, expected[sym.Name])
		}
	}
}

func TestResolveNestedLocal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	firstLocal := NewEnclosedSymbolTable(global)
	firstLocal.Define("c")
	firstLocal.Define("d")

	secondLocal := NewEnclosedSymbolTable(firstLocal)
	secondLocal.Define("e")
	secondLocal.Define("f")

	tests := []struct {
		table    *SymbolTable
		expected []Symbol
	}{
		{
			firstLocal,
			[]Symbol{
				{Name: "a", Scope: GlobalScope, Index: 0},
				{Name: "b", Scope: GlobalScope, Index: 1},
				{Name: "c", Scope: LocalScope, Index: 0},
				{Name: "d", Scope: LocalScope, Index: 1},
			},
		},
		{
			secondLocal,
			[]Symbol{
				{Name: "a", Scope: GlobalScope, Index: 0},
				{Name: "b", Scope: GlobalScope, Index: 1},
				{Name: "e", Scope: LocalScope, Index: 0},
				{Name: "f", Scope: LocalScope, Index: 1},
			},
		},
	}
	for _, tt := range tests {
		for _, sym := range tt.expected {
			result, ok := tt.table.Resolve(sym.Name)
			if !ok {
				t.Errorf("name %s not resolvable", sym.Name)
				continue
			}
			if result != sym {
				t.Errorf("name %s = %v, expected %v", sym.Name, result, sym)
			}
		}
	}
}
