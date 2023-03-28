package intermediate

import (
	"bytes"
	"fmt"
)

package intermediate

import (
"bytes"
"fmt"
)

// CodeGenerator represents a code generator for the intermediate code.
type CodeGenerator struct {
	buffer bytes.Buffer // The buffer to hold the generated code.
}

// NewCodeGenerator creates a new code generator.
func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

// Generate generates code for the given node.
func (c *CodeGenerator) Generate(node Node) {
	switch n := node.(type) {
	case *Assignment:
		c.buffer.WriteString(fmt.Sprintf("%s = %s\n", n.Target, n.Operand.String()))
	case *BinaryOp:
		c.buffer.WriteString(fmt.Sprintf("%s %s %s\n", n.Left.String(), n.Op, n.Right.String()))
	case *Integer:
		c.buffer.WriteString(fmt.Sprintf("%d\n", n.Value))
	case *UnaryOp:
		c.buffer.WriteString(fmt.Sprintf("%s%s\n", n.Op, n.Operand.String()))
	}
}

// GetCode returns the generated code.
func (c *CodeGenerator) GetCode() string {
	return c.buffer.String()
}
