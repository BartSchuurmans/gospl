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
		p.next()
		return p.continueVariableDeclaration(t, name)
	case token.ROUND_BRACKET_OPEN:
		p.next()
		return p.continueFunctionDeclaration(t, name)
	default:
		p.errorExpected(p.pos, "declaration")
		p.next()
		return &ast.BadDeclaration{}
	}
}

func (p *Parser) parseType() *ast.Type {
	name := p.parseIdentifier()

	return &ast.Type{
		Name: name,
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

func (p *Parser) continueVariableDeclaration(t *ast.Type, name *ast.Identifier) *ast.VariableDeclaration {
	initializer := p.parseExpression()
	p.expect(token.SEMICOLON)

	return &ast.VariableDeclaration{
		Type:        t,
		Name:        name,
		Initializer: initializer,
	}
}

func (p *Parser) parseExpression() ast.Expression {
	switch p.tok {
	case token.INTEGER:
		val := p.lit
		p.next()
		return &ast.LiteralExpression{
			Value: val,
		}
	default:
		p.errorExpected(p.pos, "expression")
		p.next()
		return &ast.BadExpression{}
	}
}

func (p *Parser) continueFunctionDeclaration(returnType *ast.Type, name *ast.Identifier) *ast.FunctionDeclaration {
	params := p.parseFunctionParameters()
	p.expect(token.ROUND_BRACKET_CLOSE)

	body := p.parseBlockStatement()

	return &ast.FunctionDeclaration{
		Name: name,
		Type: &ast.FunctionType{
			Return:     returnType,
			Parameters: params,
		},
		Body: body,
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.tok {
	case token.RETURN:
		p.next()
		expr := p.parseExpression()
		p.expect(token.SEMICOLON)
		return &ast.ReturnStatement{
			Value: expr,
		}
	default:
		p.errorExpected(p.pos, "statement")
		p.next()
		return &ast.BadStatement{}
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
