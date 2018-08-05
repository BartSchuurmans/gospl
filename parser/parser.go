package parser

import (
	"github.com/Minnozz/gompiler/ast"
	"github.com/Minnozz/gompiler/scanner"
	"github.com/Minnozz/gompiler/token"
)

type Parser struct {
	fileInfo *token.FileInfo
	Errors   scanner.ErrorList
	scanner  scanner.Scanner

	// Current scanner token
	pos token.Pos
	tok token.Token
	lit string
}

func (p *Parser) Init(fileInfo *token.FileInfo, src []byte) {
	p.fileInfo = fileInfo
	p.scanner.Init(fileInfo, src, func(pos token.Position, msg string) {
		p.Errors.Add(pos, msg)
	})

	p.next()
}
func (p *Parser) Parse() *ast.File {
	var declarations []ast.Declaration
	for p.tok != token.EOF {
		declarations = append(declarations, p.parseDeclaration())
	}

	return &ast.File{
		Declarations: declarations,
	}
}

func (p *Parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *Parser) error(pos token.Pos, msg string) {
	position := p.fileInfo.Position(pos)
	p.Errors.Add(position, msg)
}

func (p *Parser) errorExpected(pos token.Pos, what string) {
	if p.pos == pos {
		p.error(pos, "expected "+what+", got "+p.tok.String())
	} else {
		p.error(pos, "expected "+what)
	}
}

func (p *Parser) expect(tok token.Token) token.Pos {
	pos := p.pos
	if p.tok != tok {
		p.errorExpected(pos, tok.String())
	}
	p.next()
	return pos
}

func (p *Parser) parseDeclaration() ast.Declaration {
	t := p.parseType()
	name := p.parseIdentifier()

	switch p.tok {
	case token.IS:
		return p.continueVariableDeclaration(t, name)
	case token.ROUND_BRACKET_OPEN:
		return p.continueFunctionDeclaration(t, name)
	default:
		p.errorExpected(p.pos, "declaration")
		p.next()
		return &ast.BadDeclaration{}
	}
}

func (p *Parser) parseType() ast.Type {
	switch p.tok {
	case token.IDENTIFIER:
		name := p.parseIdentifier()
		return &ast.NamedType{
			Name: name,
		}
	case token.ROUND_BRACKET_OPEN:
		p.next()
		left := p.parseType()
		p.expect(token.COMMA)
		right := p.parseType()
		p.expect(token.ROUND_BRACKET_CLOSE)
		return &ast.TupleType{
			Left:  left,
			Right: right,
		}
	case token.SQUARE_BRACKET_OPEN:
		p.next()
		el := p.parseType()
		p.expect(token.SQUARE_BRACKET_CLOSE)
		return &ast.ListType{
			ElementType: el,
		}
	default:
		p.errorExpected(p.pos, "type")
		p.next()
		return &ast.BadType{}
	}
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	name := "-"
	if p.tok == token.IDENTIFIER {
		name = p.lit
		p.next()
	} else {
		p.expect(token.IDENTIFIER)
	}

	return &ast.Identifier{
		Name: name,
	}
}

func (p *Parser) continueVariableDeclaration(t ast.Type, name *ast.Identifier) *ast.VariableDeclaration {
	p.expect(token.IS)
	initializer := p.parseExpression()
	p.expect(token.SEMICOLON)

	return &ast.VariableDeclaration{
		Type:        t,
		Name:        name,
		Initializer: initializer,
	}
}

func (p *Parser) parseExpression() ast.Expression {
	var lhs ast.Expression

	switch p.tok {
	case token.INTEGER, token.EMPTY_LIST:
		lhs = p.parseLiteralExpression()
	case token.IDENTIFIER:
		ident := p.parseIdentifier()
		switch p.tok {
		case token.ROUND_BRACKET_OPEN:
			// Function call
			lhs = p.continueFunctionCall(ident)
		default:
			// Identifier
			lhs = ident
		}
	default:
		p.errorExpected(p.pos, "expression")
		p.next()
		lhs = &ast.BadExpression{}
	}

	switch p.tok {
	case token.PLUS, token.MINUS, token.MULTIPLY, token.DIVIDE, token.MODULO,
		token.EQUALS, token.LESS_THAN, token.GREATER_THAN, token.LESS_THAN_EQUALS, token.GREATER_THAN_EQUALS, token.NOT_EQUALS,
		token.AND, token.OR, token.COLON:
		op := p.tok
		p.next()
		rhs := p.parseExpression()
		return &ast.BinaryExpression{
			Left:     lhs,
			Operator: op,
			Right:    rhs,
		}
	default:
		// Expression ends here
		return lhs
	}
}

func (p *Parser) parseLiteralExpression() *ast.LiteralExpression {
	switch p.tok {
	case token.INTEGER, token.EMPTY_LIST:
		kind, value := p.tok, p.lit
		p.next()
		return &ast.LiteralExpression{
			Kind:  kind,
			Value: value,
		}
	default:
		p.errorExpected(p.pos, "literal expression")
		p.next()
		return &ast.LiteralExpression{
			Kind:  token.INVALID,
			Value: "[BAD LITERAL EXPRESSION]",
		}
	}
}

func (p *Parser) continueFunctionDeclaration(returnType ast.Type, name *ast.Identifier) *ast.FunctionDeclaration {
	p.expect(token.ROUND_BRACKET_OPEN)
	params := p.parseFunctionParameters()
	p.expect(token.ROUND_BRACKET_CLOSE)

	varDecls, stmts := p.parseFunctionBody()

	return &ast.FunctionDeclaration{
		Name: name,
		Type: &ast.FunctionType{
			Return:     returnType,
			Parameters: params,
		},
		Variables:  varDecls,
		Statements: stmts,
	}
}

func (p *Parser) parseFunctionParameters() *ast.FunctionParameters {
	var params []*ast.FunctionParameter

	if p.tok != token.ROUND_BRACKET_CLOSE {
	parameters:
		for {
			params = append(params, p.parseFunctionParameter())

			switch p.tok {
			case token.COMMA:
				p.next()
			case token.ROUND_BRACKET_CLOSE:
				break parameters
			default:
				p.errorExpected(p.pos, token.COMMA.String()+" or "+token.ROUND_BRACKET_CLOSE.String())
				p.next()
				break parameters
			}
		}
	}

	return &ast.FunctionParameters{
		Parameters: params,
	}
}

func (p *Parser) parseFunctionParameter() *ast.FunctionParameter {
	t := p.parseType()
	name := p.parseIdentifier()

	return &ast.FunctionParameter{
		Type: t,
		Name: name,
	}
}

func (p *Parser) parseFunctionBody() ([]*ast.VariableDeclaration, []ast.Statement) {
	p.expect(token.CURLY_BRACKET_OPEN)

	var varDecls []*ast.VariableDeclaration
	var stmts []ast.Statement

	allowVardecl := true
	for p.tok != token.CURLY_BRACKET_CLOSE {
		varDecl, stmt := p.parseVariableDeclarationOrStatement(allowVardecl)
		if varDecl != nil {
			varDecls = append(varDecls, varDecl)
		} else {
			stmts = append(stmts, stmt)
			allowVardecl = false
		}
	}

	p.expect(token.CURLY_BRACKET_CLOSE)

	return varDecls, stmts
}

func (p *Parser) parseVariableDeclarationOrStatement(allowVariableDeclaration bool) (*ast.VariableDeclaration, ast.Statement) {
	switch p.tok {
	case token.RETURN:
		return nil, p.parseReturnStatement()
	case token.IF:
		return nil, p.parseIfStatement()
	case token.CURLY_BRACKET_OPEN:
		return nil, p.parseBlockStatement()
	case token.IDENTIFIER:
		ident := p.parseIdentifier()

		// Possible statements
		switch p.tok {
		case token.IS:
			return nil, p.continueAssignmentStatement(ident)
		case token.ROUND_BRACKET_OPEN:
			return nil, p.continueFunctionCall(ident)
		}

		if !allowVariableDeclaration {
			p.errorExpected(p.pos, "assignment or function call")
			p.next()
			return nil, &ast.BadStatement{}
		}

		// Variable declaration with type ident
		t := &ast.NamedType{
			Name: ident,
		}
		name := p.parseIdentifier()
		return p.continueVariableDeclaration(t, name), nil
	case token.SQUARE_BRACKET_OPEN:
		if !allowVariableDeclaration {
			p.errorExpected(p.pos, "statement")
			p.next()
			return nil, &ast.BadStatement{}
		}
		t := p.parseType()
		name := p.parseIdentifier()
		return p.continueVariableDeclaration(t, name), nil
	case token.WHILE:
		return nil, p.parseWhileStatement()
	default:
		if allowVariableDeclaration {
			p.errorExpected(p.pos, "variable declaration or statement")
		} else {
			p.errorExpected(p.pos, "statement")
		}
		p.next()
		return nil, &ast.BadStatement{}
	}
}

func (p *Parser) parseStatement() ast.Statement {
	_, stmt := p.parseVariableDeclarationOrStatement(false)
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	p.expect(token.RETURN)

	expr := p.parseExpression()

	p.expect(token.SEMICOLON)

	return &ast.ReturnStatement{
		Value: expr,
	}
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	p.expect(token.IF)
	p.expect(token.ROUND_BRACKET_OPEN)

	cond := p.parseExpression()

	p.expect(token.ROUND_BRACKET_CLOSE)

	body := p.parseStatement()

	var elseStmt ast.Statement
	if p.tok == token.ELSE {
		p.next()
		elseStmt = p.parseStatement()
	}

	return &ast.IfStatement{
		Condition: cond,
		Body:      body,
		Else:      elseStmt,
	}
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	p.expect(token.WHILE)
	p.expect(token.ROUND_BRACKET_OPEN)

	cond := p.parseExpression()

	p.expect(token.ROUND_BRACKET_CLOSE)

	body := p.parseStatement()

	return &ast.WhileStatement{
		Condition: cond,
		Body:      body,
	}
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	p.expect(token.CURLY_BRACKET_OPEN)

	var stmts []ast.Statement
	for p.tok != token.CURLY_BRACKET_CLOSE {
		stmts = append(stmts, p.parseStatement())
	}

	p.expect(token.CURLY_BRACKET_CLOSE)

	return &ast.BlockStatement{
		List: stmts,
	}
}

func (p *Parser) continueAssignmentStatement(name *ast.Identifier) *ast.AssignmentStatement {
	p.expect(token.IS)
	value := p.parseExpression()
	p.expect(token.SEMICOLON)

	return &ast.AssignmentStatement{
		Name:  name,
		Value: value,
	}
}

func (p *Parser) continueFunctionCall(name *ast.Identifier) *ast.FunctionCallStatement {
	p.expect(token.ROUND_BRACKET_OPEN)

	var args []ast.Expression

	if p.tok != token.ROUND_BRACKET_CLOSE {
	arguments:
		for {
			args = append(args, p.parseExpression())

			switch p.tok {
			case token.COMMA:
				p.next()
			case token.ROUND_BRACKET_CLOSE:
				break arguments
			default:
				p.errorExpected(p.pos, token.COMMA.String()+" or "+token.ROUND_BRACKET_CLOSE.String())
				p.next()
				break arguments
			}
		}
	}

	p.expect(token.ROUND_BRACKET_CLOSE)

	return &ast.FunctionCallStatement{
		Name:      name,
		Arguments: args,
	}
}
