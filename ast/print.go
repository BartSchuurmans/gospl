package ast

func (b *BadDeclaration) Print() string {
	return "[BAD DECLARATION]"
}

func (b *BadExpression) Print() string {
	return "[BAD EXPRESSION]"
}

func (b *BadStatement) Print() string {
	return "[BAD STATEMENT]"
}

func (v *VariableDeclaration) Print() string {
	return v.Type.Print() + " " + v.Name.Print() + " = " + v.Initializer.Print() + "\n"
}

func (f *FunctionDeclaration) Print() string {
	return f.Type.Return.Print() + " " + f.Name.Print() + "(" + f.Type.Parameters.Print() + ") " + f.Body.Print() + "\n"
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

func (f *File) Print() string {
	out := ""
	for i, decl := range f.Declarations {
		if i > 0 {
			out += "\n"
		}
		out += decl.Print()
	}
	return out
}

func (i *Identifier) Print() string {
	return i.Name
}

func (t *Type) Print() string {
	return t.Name.Print()
}

func (b *BlockStatement) Print() string {
	out := "{\n"
	for _, stmt := range b.List {
		out += "\t" + stmt.Print() + "\n"
	}
	out += "}"
	return out
}

func (r *ReturnStatement) Print() string {
	return "return " + r.Value.Print() + ";"
}
