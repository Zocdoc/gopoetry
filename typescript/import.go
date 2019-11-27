package typescript

// ImportDeclaration declares an import
type ImportDeclaration struct {
	module        string
	namedImports  []string
	defaultImport string
}

// DefaultImport creates a default import. import * as foo from 'module';
func DefaultImport(module string, as string) *ImportDeclaration {
	return &ImportDeclaration{
		module:        module,
		defaultImport: as,
	}
}

// NamedImport creates a default import. import * as foo from 'module';
func NamedImport(module string, params ...string) *ImportDeclaration {
	return &ImportDeclaration{
		module:       module,
		namedImports: params,
	}
}

// WriteCode writes the imports to the writer
func (imp *ImportDeclaration) WriteCode(writer CodeWriter) {
	writer.Write("import ")

	if len(imp.namedImports) > 0 {
		writer.Write("{ ")
		for i, name := range imp.namedImports {
			writer.Write(name)
			if i < len(imp.namedImports)-1 {
				writer.Write(", ")
			}
		}
		writer.Write(" }")
	}

	if imp.defaultImport != "" {
		writer.Write("* as " + imp.defaultImport)
	}

	writer.Write(" from '" + imp.module + "';")
}
