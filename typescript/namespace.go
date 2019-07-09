package typescript

// NamespaceDeclaration declares a namespace
type NamespaceDeclaration struct {
	namespace     string
	imports       []Writable
	declarations  []Writable
	referencePath *string
}

// AddImports adds imports to the namespace
func (ns *NamespaceDeclaration) AddImports(imports ...Writable) *NamespaceDeclaration {
	ns.imports = append(ns.imports, imports...)
	return ns
}

// NamedImport adds a named import
func (ns *NamespaceDeclaration) NamedImport(module string, params ...string) *NamespaceDeclaration {
	ns.AddImports(NamedImport(module, params...))
	return ns
}

// DefaultImport adds a default import
func (ns *NamespaceDeclaration) DefaultImport(module string, as string) *NamespaceDeclaration {
	ns.AddImports(DefaultImport(module, as))
	return ns
}

// AddDeclarations adds declarations to the namespace
func (ns *NamespaceDeclaration) AddDeclarations(declarations ...Writable) *NamespaceDeclaration {
	ns.declarations = append(ns.declarations, declarations...)
	return ns
}

// WithReference adds a file reference for namespace compilation
func (ns *NamespaceDeclaration) WithReference(filePath string) *NamespaceDeclaration {
	ns.referencePath = &filePath
	return ns
}

// Namespace creates a NamespaceDeclaration
func Namespace(namespace string) *NamespaceDeclaration {
	return &NamespaceDeclaration{
		namespace:    namespace,
		imports:      []Writable{},
		declarations: []Writable{},
	}
}

// WriteCode writes the namespace to the writer
func (ns *NamespaceDeclaration) WriteCode(writer CodeWriter) {
	if len(ns.imports) > 0 {
		for _, using := range ns.imports {
			using.WriteCode(writer)
			writer.Eol()
		}
		writer.Eol()
	}

	if ns.referencePath != nil {
		writer.Write("/// <reference path=\"")
		writer.Write(*ns.referencePath)
		writer.Write("\" />")
		writer.Eol()
	}

	writer.Write("namespace " + ns.namespace)
	writer.Begin()

	for index, class := range ns.declarations {
		if index > 0 {
			writer.Eol()
		}
		class.WriteCode(writer)
	}
	writer.End()
}
