package backend

import "io/ioutil"

// GenerateCode generates code from the AST and writes it to the specified output file.
func GenerateCode(ast *parser.AST, outputFile string) error {
	// Create an intermediate code generator
	icg := intermediateCodeGenerator{}

	// Generate intermediate code from the AST
	nodes, err := icg.generateIntermediateCode(ast)
	if err != nil {
		return err
	}

	// Create an assembly code generator
	asm := asm{}

	// Generate assembly code from the intermediate code nodes
	assembly, err := asm.generateAssembly(nodes)
	if err != nil {
		return err
	}

	// Write the assembly code to the output file
	err = ioutil.WriteFile(outputFile, []byte(assembly), 0644)
	if err != nil {
		return err
	}

	return nil
}
