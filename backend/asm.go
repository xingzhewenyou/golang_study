package backend

import (
	"fmt"
	"strings"
)

// generateAssembly takes a slice of intermediate code nodes and returns
// a string containing the corresponding assembly code.
func generateAssembly(nodes []intermediateCodeNode) (string, error) {
	var b strings.Builder

	// Write the assembly file header
	fmt.Fprintf(&b, "section .text\n\n")

	// Loop over each intermediate code node
	for _, node := range nodes {
		switch node.opcode {
		case opcodeAdd:
			// Add the two operands and store the result
			fmt.Fprintf(&b, "add eax, %d\n", node.operand1)
		case opcodeSubtract:
			// Subtract the second operand from the first and store the result
			fmt.Fprintf(&b, "sub eax, %d\n", node.operand2)
		case opcodeMultiply:
			// Multiply the two operands and store the result
			fmt.Fprintf(&b, "imul eax, %d\n", node.operand1)
		case opcodeDivide:
			// Divide the first operand by the second and store the result
			fmt.Fprintf(&b, "cdq\n")
			fmt.Fprintf(&b, "idiv %d\n", node.operand2)
		default:
			// If the opcode is not recognized, return an error
			return "", fmt.Errorf("unknown opcode %d", node.opcode)
		}
	}

	// Write the assembly file footer
	fmt.Fprintf(&b, "\n")

	return b.String(), nil
}
