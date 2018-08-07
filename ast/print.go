package ast

import (
	"strings"
)

func PrintSource(node Node) string {
	switch n := node.(type) {
	// File
	case *File:
		// TODO: Comments are not printed
		out := ""
		for i, decl := range n.Declarations {
			if i > 0 {
				out += "\n\n"
			}
			out += PrintSource(decl)
		}
		return out

	// Declarations
	case *VariableDeclaration:
		return PrintSource(n.Type) + " " + PrintSource(n.Name) + " = " + PrintSource(n.Initializer) + ";"
	case *FunctionDeclaration:
		out := PrintSource(n.ReturnType) + " " + PrintSource(n.Name) + "(" + PrintSource(n.Parameters) + ") {\n"
		if len(n.Variables) > 0 {
			for _, varDecl := range n.Variables {
				out += indent(PrintSource(varDecl)) + "\n"
			}
			out += "\n"
		}
		for _, stmt := range n.Statements {
			out += indent(PrintSource(stmt)) + "\n"
		}
		out += "}"
		return out
	case *FunctionParameters:
		out := ""
		for i, param := range n.Parameters {
			if i > 0 {
				out += ", "
			}
			out += PrintSource(param)
		}
		return out
	case *FunctionParameter:
		return PrintSource(n.Type) + " " + PrintSource(n.Name)
	case *BadDeclaration:
		return "/* BAD DECLARATION */"

	// Expressions
	case *LiteralExpression:
		return n.Value
	case *UnaryExpression:
		return n.Operator.Print() + PrintSource(n.Expression)
	case *BinaryExpression:
		return PrintSource(n.Left) + " " + n.Operator.Print() + " " + PrintSource(n.Right)
	case *FunctionCallExpression:
		out := PrintSource(n.Name) + "("
		for i, expr := range n.Arguments {
			if i > 0 {
				out += ", "
			}
			out += PrintSource(expr)
		}
		out += ")"
		return out
	case *ParenthesizedExpression:
		return "(" + PrintSource(n.Expression) + ")"
	case *TupleExpression:
		return "(" + PrintSource(n.Left) + ", " + PrintSource(n.Right) + ")"
	case *Identifier:
		return n.Name
	case *BadExpression:
		return "/* BAD EXPRESSION */"

	// Statements
	case *BlockStatement:
		out := "{\n"
		for _, stmt := range n.List {
			out += indent(PrintSource(stmt)) + "\n"
		}
		out += "}"
		return out
	case *ReturnStatement:
		out := "return"
		if n.Value != nil {
			out += " " + PrintSource(n.Value)
		}
		out += ";"
		return out
	case *IfStatement:
		out := "if(" + PrintSource(n.Condition) + ") " + PrintSource(n.Body)
		if n.Else != nil {
			out += " else " + PrintSource(n.Else)
		}
		return out
	case *WhileStatement:
		return "while(" + PrintSource(n.Condition) + ") " + PrintSource(n.Body)
	case *AssignmentStatement:
		return PrintSource(n.Name) + " = " + PrintSource(n.Value) + ";"
	case *FunctionCallStatement:
		return PrintSource(n.FunctionCall) + ";"
	case *BadStatement:
		return "/* BAD STATEMENT */"

	// Types
	case *NamedType:
		return PrintSource(n.Name)
	case *TupleType:
		return "(" + PrintSource(n.Left) + ", " + PrintSource(n.Right) + ")"
	case *ListType:
		return "[" + PrintSource(n.ElementType) + "]"
	case *BadType:
		return "/* BAD TYPE */"

	default:
		return "/* UNKNOWN AST NODE */"
	}
}

func indent(s string) string {
	return "\t" + strings.Replace(s, "\n", "\n\t", -1)
}
