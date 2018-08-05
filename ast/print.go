package ast

import (
	"strings"
)

func (b *BadDeclaration) Print() string {
	return "[BAD DECLARATION]"
}

func (b *BadExpression) Print() string {
	return "[BAD EXPRESSION]"
}

func (b *BadStatement) Print() string {
	return "[BAD STATEMENT]"
}

func (b *BadType) Print() string {
	return "[BAD TYPE]"
}

func (f *File) Print() string {
	out := ""
	for i, decl := range f.Declarations {
		if i > 0 {
			out += "\n\n"
		}
		out += decl.Print()
	}
	return out
}

func (v *VariableDeclaration) Print() string {
	return v.Type.Print() + " " + v.Name.Print() + " = " + v.Initializer.Print() + ";"
}

func (f *FunctionDeclaration) Print() string {
	out := f.Type.Return.Print() + " " + f.Name.Print() + "(" + f.Type.Parameters.Print() + ") {\n"
	if len(f.Variables) > 0 {
		for _, varDecl := range f.Variables {
			out += indent(varDecl.Print()) + "\n"
		}
		out += "\n"
	}
	for _, stmt := range f.Statements {
		out += indent(stmt.Print()) + "\n"
	}
	out += "}"
	return out
}

func (f *FunctionParameters) Print() string {
	out := ""
	for i, param := range f.Parameters {
		if i > 0 {
			out += ", "
		}
		out += param.Print()
	}
	return out
}

func (f *FunctionParameter) Print() string {
	return f.Type.Print() + " " + f.Name.Print()
}

func (l *LiteralExpression) Print() string {
	return l.Value
}

func (b *BinaryExpression) Print() string {
	return b.Left.Print() + " " + b.Operator.Print() + " " + b.Right.Print()
}

func (f *FunctionCallExpression) Print() string {
	out := f.Name.Print() + "("
	for i, expr := range f.Arguments {
		if i > 0 {
			out += ", "
		}
		out += expr.Print()
	}
	out += ")"
	return out
}

func (p *ParenthesizedExpression) Print() string {
	return "(" + p.Expression.Print() + ")"
}

func (t *TupleExpression) Print() string {
	return "(" + t.Left.Print() + ", " + t.Right.Print() + ")"
}

func (i *Identifier) Print() string {
	return i.Name
}

func (n *NamedType) Print() string {
	return n.Name.Print()
}

func (t *TupleType) Print() string {
	return "(" + t.Left.Print() + ", " + t.Right.Print() + ")"
}

func (l *ListType) Print() string {
	return "[" + l.ElementType.Print() + "]"
}

func (b *BlockStatement) Print() string {
	out := "{\n"
	for _, stmt := range b.List {
		out += indent(stmt.Print()) + "\n"
	}
	out += "}"
	return out
}

func (r *ReturnStatement) Print() string {
	out := "return"
	if r.Value != nil {
		out += " " + r.Value.Print()
	}
	out += ";"
	return out
}

func (i *IfStatement) Print() string {
	out := "if(" + i.Condition.Print() + ") " + i.Body.Print()
	if i.Else != nil {
		out += " else " + i.Else.Print()
	}
	return out
}

func (w *WhileStatement) Print() string {
	return "while(" + w.Condition.Print() + ") " + w.Body.Print()
}

func (a *AssignmentStatement) Print() string {
	return a.Name.Print() + " = " + a.Value.Print() + ";"
}

func (f *FunctionCallStatement) Print() string {
	return f.FunctionCall.Print() + ";"
}

func indent(s string) string {
	return "\t" + strings.Replace(s, "\n", "\n\t", -1)
}
