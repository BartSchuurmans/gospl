package ast

type Visitor interface {
	Visit(n Node)
	End(n Node)
}

type VisitorFunc func(Node)

func (f VisitorFunc) Visit(n Node) {
	if n != nil {
		f(n)
	}
}

func (f VisitorFunc) End(n Node) {
	// Ignore
}

func WalkFunc(n Node, v VisitorFunc) {
	Walk(n, v)
}

func Walk(n Node, v Visitor) {
	// Visit node itself
	v.Visit(n)

	// Visit node children (if any)
	switch nv := n.(type) {
	case *File:
		for _, ce := range nv.Declarations {
			Walk(ce, v)
		}
		for _, ce := range nv.Comments {
			Walk(ce, v)
		}
	case *VariableDeclaration:
		Walk(nv.Type, v)
		Walk(nv.Name, v)
		Walk(nv.Initializer, v)
	case *FunctionDeclaration:
		Walk(nv.ReturnType, v)
		Walk(nv.Name, v)
		Walk(nv.Parameters, v)
		for _, ce := range nv.Variables {
			Walk(ce, v)
		}
		for _, ce := range nv.Statements {
			Walk(ce, v)
		}
	case *FunctionParameters:
		for _, ce := range nv.Parameters {
			Walk(ce, v)
		}
	case *FunctionParameter:
		Walk(nv.Type, v)
		Walk(nv.Name, v)
	case *UnaryExpression:
		Walk(nv.Operand, v)
	case *BinaryExpression:
		Walk(nv.Left, v)
		Walk(nv.Right, v)
	case *FunctionCallExpression:
		Walk(nv.Name, v)
		for _, ce := range nv.Arguments {
			Walk(ce, v)
		}
	case *ParenthesizedExpression:
		Walk(nv.Expression, v)
	case *TupleExpression:
		Walk(nv.Left, v)
		Walk(nv.Right, v)
	case *BlockStatement:
		for _, ce := range nv.List {
			Walk(ce, v)
		}
	case *ReturnStatement:
		Walk(nv.Value, v)
	case *IfStatement:
		Walk(nv.Condition, v)
		Walk(nv.Body, v)
		Walk(nv.Else, v)
	case *WhileStatement:
		Walk(nv.Condition, v)
		Walk(nv.Body, v)
	case *AssignmentStatement:
		Walk(nv.Name, v)
		Walk(nv.Value, v)
	case *FunctionCallStatement:
		Walk(nv.FunctionCall, v)
	case *NamedType:
		Walk(nv.Name, v)
	case *TupleType:
		Walk(nv.Left, v)
		Walk(nv.Right, v)
	case *ListType:
		Walk(nv.ElementType, v)
	}

	// Indicate end of children
	v.End(n)
}
