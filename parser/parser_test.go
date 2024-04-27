package parser

import (
	"MonkeyInterpreterByGO/ast"
	"MonkeyInterpreterByGO/lexer"
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	input := `
	let x = 6;
	let y = 21;
	let KFC = 50;
	let is = 10;
	let zz = 20;
`
	// l 是词法分析器
	l := lexer.New(input)
	// p 是语法分析器
	p := New(*l)
	fmt.Println("start parse!")
	program := p.ParseProgram()
	fmt.Printf("program.statement is %d\n", len(program.Statements))
	if program == nil {
		t.Fatalf("ParseProgram() return nil!")
	}
	if len(program.Statements) == 0 {
		t.Fatalf("program.statement is 0")
	}
	checkErrors(t, p)
	validate := []struct {
		expectLiteral string
	}{
		{"x"},
		{"y"},
		{"KFC"},
		{"is"},
		{"zz"},
	}
	for i, stmt := range program.Statements {
		if !testParser(t, stmt, validate[i].expectLiteral) {
			t.Fatalf("false")
		}
	}
}

func checkErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	if len(errors) != 0 {
		t.Errorf("Has %d Error In Parse\n", len(errors))
	}
	for _, msg := range errors {
		t.Errorf("%s\n", msg)
	}
}

func testParser(t *testing.T, stmt ast.StatementNode, expect string) bool {
	if stmt.TokenLiteral() != "let" {
		t.Fatalf("false")
		return false
	}
	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Fatalf("false")
		return false
	}
	if letStmt.Name.TokenLiteral() != expect {
		t.Fatalf("false")
		return false
	}
	if letStmt.Name.Value != expect {
		t.Fatalf("false")
		return false
	}
	return true
}

func TestReturnStatement(t *testing.T) {
	input := "return 5;" +
		"return 10;" +
		"return handle(20);"
	l := lexer.New(input)
	p := New(*l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Parse program is nil!")
	}
	if len(program.Statements) == 0 {
		t.Fatalf("program.statements is 0")
	}
	checkErrors(t, p)
	validate := []struct {
		expectLiteral string
	}{
		{"5"},
		{"10"},
		{"handle"},
	}
	for i, stmt := range program.Statements {
		if !testReturn(t, stmt, validate[i].expectLiteral) {
			t.Fatalf("test Parser fail！")
		}
	}
}

func testReturn(t *testing.T, stmt ast.StatementNode, expect string) bool {
	if stmt.TokenLiteral() != "return" {
		t.Fatalf("expected return, got %s", stmt.TokenLiteral())
	}
	_, ok := stmt.(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("stmt is not ReturnStatement")
	}
	return true
}
