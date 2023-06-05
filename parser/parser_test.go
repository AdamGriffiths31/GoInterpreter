package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foo = 5000;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program statemets, expected 3 got %d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}
	for i, tt := range tests {
		statement := program.Statements[i]
		if !testLetStatement(t, statement, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("expected 'let' got %t", s)
		return false
	}

	letStatement, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement, expected LetStatement got %t", s)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("value, expected %s got %s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("tokenLiteral, expected %s got %s", name,
			letStatement.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func TestRetuenStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 300400;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("statements, expected 3 got %d", len(program.Statements))
	}

	for _, statement := range program.Statements {
		returnStmt, ok := statement.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not ReturnStatement, got %T", statement)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("TokenLiteral not return, got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("statements, expected 1 got %d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement not ExpressionStatement, got %T",
			program.Statements[0])
	}

	ident, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not identifier, got %T", statement.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident value not foobar, got %s", ident.Value)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("tokenLiteral not foobar, got %s", ident.TokenLiteral())
	}
}

func TestIntergerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("statements, expected 1 got %d", len(program.Statements))
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement not ExpressionStatement, got %T", statement)
	}

	literal, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not IntergerLiteral, got %T", statement.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal value not 5, got %d", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("tokenLiteral not '5' got %s", literal.TokenLiteral())
	}
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("statements, expected 1 got %d", len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement not expression, got %T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("expression not prefix, got %T", statement.Expression)
		}

		if expression.Operator != tt.operator {
			t.Fatalf("operator, expected %s got %s", tt.operator,
				expression.Operator)
		}

		if !testIntegerLiteral(t, expression.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("not IntergerLiteral, got %T", il)
	}

	if integer.Value != value {
		t.Errorf("value, expected %d got %d", integer.Value, value)
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("tokenLiteral, expected %d got %s",
			value, integer.TokenLiteral())
		return false
	}
	return true
}
