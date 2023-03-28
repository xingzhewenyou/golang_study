package intermediate

package intermediate

import "fmt"

// Node represents a node in the intermediate code AST.
type Node interface {
	// String returns a string representation of the node.
	String() string
}

// Assignment represents an assignment statement in the AST.
type Assignment struct {
	Target  string // The name of the variable being assigned.
	Operand Node   // The expression being assigned to the variable.
}

// String returns a string representation of the assignment node.
func (a *Assignment) String() string {
	return fmt.Sprintf("%s = %s", a.Target, a.Operand.String())
}

// BinaryOp represents a binary operation in the AST.
type BinaryOp struct {
	Left  Node   // The left operand of the binary operation.
	Op    string // The operator of the binary operation.
	Right Node   // The right operand of the binary operation.
}

// String returns a string representation of the binary operation node.
func (b *BinaryOp) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left.String(), b.Op, b.Right.String())
}

// Integer represents an integer constant in the AST.
type Integer struct {
	Value int // The value of the integer constant.
}

// String returns a string representation of the integer node.
func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

// UnaryOp represents a unary operation in the AST.
type UnaryOp struct {
	Op    string // The operator of the unary operation.
	Operand Node   // The operand of the unary operation.
}

// String returns a string representation of the unary operation node.
func (u *UnaryOp) String() string {
	return fmt.Sprintf("(%s%s)", u.Op, u.Operand.String())
}

