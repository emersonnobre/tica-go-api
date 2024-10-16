package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	// Define o argumento para receber o caminho do arquivo Go
	filePath := flag.String("file", "", "Caminho para o arquivo Go contendo a struct")
	flag.Parse()

	// Verifica se o caminho do arquivo foi fornecido
	if *filePath == "" {
		fmt.Println("Por favor, forneça o caminho para o arquivo Go usando o argumento -file")
		os.Exit(1)
	}

	// Analisa o arquivo Go
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, *filePath, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Encontra a struct no arquivo
	var structName string
	var structFields []*ast.Field
	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				structName = typeSpec.Name.Name
				structFields = structType.Fields.List
				return false
			}
		}
		return true
	})

	// Verifica se a struct foi encontrada
	if structName == "" {
		fmt.Println("Nenhuma struct encontrada no arquivo")
		os.Exit(1)
	}

	// Gera a instrução CREATE TABLE
	json := generateJson(structFields)
	fmt.Println(json)
}

// Gera a instrução CREATE TABLE para MySQL
func generateJson(fields []*ast.Field) string {
	var sb strings.Builder
	sb.WriteString("{\n")

	for i, field := range fields {
		// Obtém o nome do campo
		fieldName := strings.ToLower(field.Names[0].Name)

		// Obtém o tipo do campo e o converte para o tipo de dado MySQL
		responseExample := getResponseExample(field.Type)

		sb.WriteRune('\t')
		sb.WriteString("\"")
		sb.WriteString(fieldName)
		sb.WriteString("\": ")
		sb.WriteString(responseExample)

		// Adiciona vírgula se não for o último campo
		if i < len(fields)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("}")
	return sb.String()
}

// Mapeia o tipo Go para o tipo de dado MySQL
func getResponseExample(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return "1"
		case "float32", "float64":
			return "3.2"
		case "string":
			return "\"Teste\""
		case "bool":
			return "true"
		case "time.Time":
			return "2024-02-29 08:43:23"
		default:
			return "\"Teste\""
		}
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok && x.Name == "time" && t.Sel.Name == "Time" {
			return "2024-02-29 08:43:23"
		}
	case *ast.ArrayType:
		return "JSON"
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok && ident.Name == "string" {
			return "\"Teste\""
		}
	case *ast.MapType:
		return "JSON"
	}
	return "\"Teste\""
}
