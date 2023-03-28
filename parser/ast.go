package parser

package parser

import "github.com/user/golang-interpreter/lexer"

// Node represents a node in the abstract syntax tree (AST).
type Node interface {
	TokenLiteral() string
}

// Statement represents a statement node in the AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node in the AST.
type Expression interface {
	Node
	expressionNode()
}

// Program represents the entire program node in the AST.
type Program struct {
	Statements []Statement // A slice of statement nodes in the program.
}

// TokenLiteral returns the literal value of the token associated with the program node.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement represents a let statement node in the AST.
type LetStatement struct {
	Token lexer.Token // The token.LET token.
	Name  *Identifier // The identifier associated with the let statement.
	Value Expression  // The expression associated with the let statement.
}

// TokenLiteral returns the literal value of the token associated with the let statement node.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// statementNode marks the let statement node as a statement node in the AST.
func (ls *LetStatement) statementNode() {}

// Identifier represents an identifier node in the AST.
type Identifier struct {
	Token lexer.Token // The token.IDENT token.
	Value string      // The value of the identifier.
}

// TokenLiteral returns the literal value of the token associated with the identifier node.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// expressionNode marks the identifier node as an expression node in the AST.
func (i *Identifier) expressionNode() {}

// ReturnStatement represents a return statement node in the AST.
type ReturnStatement struct {
	Token       lexer.Token // The token.RETURN token.
	ReturnValue Expression  // The expression associated with the return statement.
}

// TokenLiteral returns the literal value of the token associated with the return statement node.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// statementNode marks the return statement node as a statement node in the AST.
func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement represents an expression statement node in the AST.
type ExpressionStatement struct {
	Token      lexer.Token // The first token of the expression statement.
	Expression Expression  // The expression associated with the expression statement.
}

// TokenLiteral returns the literal value of the token associated with the expression statement node.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// statementNode marks the expression statement node as a statement node in the AST.
func (es *ExpressionStatement) statementNode() {}

// IntegerLiteral represents an integer literal node in the AST.
type IntegerLiteral struct {
	Token lexer.Token // The token.INT token.
	Value int64       // The value of the integer literal.
}

// TokenLiteral returns the literal value of the token associated with the integer literal node.
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// expressionNode marks the integer literal node as an expression node in the AST.
func (il *IntegerLiteral) expressionNode() {}

// LetStatement represents a let statement node in the AST. It contains the token.LET token, the identifier associated with the let statement, and the expression associated with the let statement.
// The `TokenLiteral` method returns the literal value of the token associated with the let statement node.
// The `statementNode` method marks the let statement node as a statement node in the AST.
type LetStatement struct {
	Token lexer.Token // The token.LET token.
	Name  *Identifier // The identifier associated with the let statement.
	Value Expression  // The expression associated with the let statement.
}

// TokenLiteral returns the literal value of the token associated with the let statement node.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// statementNode marks the let statement node as a statement node in the AST.
func (ls *LetStatement) statementNode() {}

// Identifier represents an identifier node in the AST. It contains the token.IDENT token and the value of the identifier.
// The `TokenLiteral` method returns the literal value of the token associated with the identifier node.
// The `expressionNode` method marks the identifier node as an expression node in the AST.
type Identifier struct {
	Token lexer.Token // The token.IDENT token.
	Value string      // The value of the identifier.
}

// TokenLiteral returns the literal value of the token associated with the identifier node.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// expressionNode marks the identifier node as an expression node in the AST.
func (i *Identifier) expressionNode() {}

// ReturnStatement represents a return statement node in the AST. It contains the token.RETURN token and the expression associated with the return statement.
// The `TokenLiteral` method returns the literal value of the token associated with the return statement node.
// The `statementNode` method marks the return statement node as a statement node in the AST.
type ReturnStatement struct {
	Token       lexer.Token // The token.RETURN token.
	ReturnValue Expression  // The expression associated with the return statement.
}

// TokenLiteral returns the literal value of the token associated with the return statement node.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// statementNode marks the return statement node as a statement node in the AST.
func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement represents an expression statement node in the AST. It contains the first token of the expression statement and the expression associated with the expression statement.
// The `TokenLiteral` method returns the literal value of the token associated with the expression statement node.
// The `statementNode` method marks the expression statement node as a statement node in the AST.
type ExpressionStatement struct {
	Token      lexer.Token // The first token of the expression statement.
	Expression Expression  // The expression associated with the expression statement.
}

// TokenLiteral returns the literal value of the token associated with the expression statement node.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// statementNode marks the expression statement node as a statement node in the AST.
func (es *ExpressionStatement) statementNode() {}

// IntegerLiteral represents an integer literal node in the AST. It contains the token.INT token and the value of the integer literal.
// The `TokenLiteral` method returns the literal value of the token associated with the integer literal node.
// The `expressionNode` method marks the integer literal node as an expression node in the AST.
type IntegerLiteral struct {
	Token lexer.Token // The token.INT token.
	Value int64       // The value of the integer literal.
}

// TokenLiteral
// returns the literal value of the token associated with the integer literal node.
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// expressionNode marks the integer literal node as an expression node in the AST.
func (il *IntegerLiteral) expressionNode() {}

// PrefixExpression represents a prefix expression node in the AST. It contains the first token of the prefix expression, the operator associated with the prefix expression, and the right-hand side expression associated with the prefix expression.
// The TokenLiteral method returns the literal value of the token associated with the prefix expression node.
// The expressionNode method marks the prefix expression node as an expression node in the AST.
type PrefixExpression struct {
	Token lexer.Token // The first token of the prefix expression.
	Operator string // The operator associated with the prefix expression.
	Right Expression // The right-hand side expression associated with the prefix expression.
}

// TokenLiteral returns the literal value of the token associated with the prefix expression node.
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// expressionNode marks the prefix expression node as an expression node in the AST.
func (pe *PrefixExpression) expressionNode() {}

// InfixExpression represents an infix expression node in the AST. It contains the left-hand side expression associated with the infix expression, the operator associated with the infix expression, and the right-hand side expression associated with the infix expression.
// The TokenLiteral method returns the literal value of the token associated with the infix expression node.
// The expressionNode method marks the infix expression node as an expression node in the AST.
type InfixExpression struct {
	Token lexer.Token // The operator token associated with the infix expression.
	Left Expression // The left-hand side expression associated with the infix expression.
	Operator string // The operator associated with the infix expression.
	Right Expression // The right-hand side expression associated with the infix expression.
}

// TokenLiteral returns the literal value of the token associated with the infix expression node.
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// expressionNode marks the infix expression node as an expression node in the AST.
func (ie *InfixExpression) expressionNode() {}

// Boolean represents a boolean literal node in the AST. It contains the token.TRUE or token.FALSE token and the value of the boolean literal.
// The TokenLiteral method returns the literal value of the token associated with the boolean literal node.
// The expressionNode method marks the boolean literal node as an expression node in the AST.
type Boolean struct {
	Token lexer.Token // The token.TRUE or token.FALSE token.
	Value bool // The value of the boolean literal.
}

// TokenLiteral returns the literal value of the token associated with the boolean literal node.
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// expressionNode marks the boolean literal node as an expression node in the AST.
func (b *Boolean) expressionNode() {}

// IfExpression represents an if expression node in the AST. It contains the token.IF token, the condition expression associated with the if expression, the consequence statement associated with the if expression, and the alternative statement associated with the if expression (which can be nil).
// The TokenLiteral method returns the literal value of the token associated with the if expression node.
// The expressionNode method marks the if expression node as an expression node in the AST.
type IfExpression struct {
	Token lexer.Token // The token.IF token.
	Condition Expression // The condition expression associated with the if expression.
	Consequence *BlockStatement // The consequence statement associated with the

	// NodeList represents a list of nodes.
	type NodeList struct {
	Nodes []Node
	Pos   Position
	}

	// Add appends a node to the list.
	func (n *NodeList) Add(node Node) {
	n.Nodes = append(n.Nodes, node)
	}

	// ExprList represents a list of expressions.
	type ExprList struct {
	Exprs []Expr
	Pos   Position
	}

	// Add appends an expression to the list.
	func (e *ExprList) Add(expr Expr) {
	e.Exprs = append(e.Exprs, expr)
	}

	// Program represents a program.
	type Program struct {
	Functions []*Function
	}

	// Function represents a function.
	type Function struct {
	Name       string
	Parameters []*Parameter
	Body       *Block
	Pos        Position
	}

	// Parameter represents a function parameter.
	type Parameter struct {
	Name string
	Type *Type
	Pos  Position
	}

	// Block represents a block of statements.
	type Block struct {
	Statements []Stmt
	Pos        Position
	}

	// Stmt represents a statement.
	type Stmt interface {
	Node
	isStmt()
	}

	// LetStmt represents a let statement.
	type LetStmt struct {
	Name  string
	Type  *Type
	Value Expr
	Pos   Position
	}

	// isStmt implements the Stmt interface.
	func (*LetStmt) isStmt() {}

	// ReturnStmt represents a return statement.
	type ReturnStmt struct {
	Value Expr
	Pos   Position
	}

	// isStmt implements the Stmt interface.
	func (*ReturnStmt) isStmt() {}

	// ExprStmt represents an expression statement.
	type ExprStmt struct {
	Expr Expr
	Pos  Position
	}

	// isStmt implements the Stmt interface.
	func (*ExprStmt) isStmt() {}

	// Type represents a type.
	type Type struct {
	Name string
	Pos  Position
	}

	// IntType represents an integer type.
	type IntType struct {
	Type
	}

	// BoolType represents a boolean type.
	type BoolType struct {
	Type
	}

	// StringType represents a string type.
	type StringType struct {
	Type
	}

	// Identifier represents an identifier.
	type Identifier struct {
	Name string
	Pos  Position
	}

	// IntLiteral represents an integer literal.
	type IntLiteral struct {
	Value int
	Pos   Position
	}

	// BoolLiteral represents a boolean literal.
	type BoolLiteral struct {
	Value bool
	Pos   Position
	}

	// StringLiteral represents a string literal.
	type StringLiteral struct {
	Value string
	Pos   Position
	}

