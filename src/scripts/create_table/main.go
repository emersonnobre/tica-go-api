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
	createTableStmt := generateCreateTableStatement(structName, structFields)
	fmt.Println(createTableStmt)
}

// Gera a instrução CREATE TABLE para MySQL
func generateCreateTableStatement(structName string, fields []*ast.Field) string {
	var sb strings.Builder
	sb.WriteString("CREATE TABLE IF NOT EXISTS ")
	sb.WriteString(strings.ToLower(structName))
	sb.WriteString(" (\n")

	for i, field := range fields {
		// Obtém o nome do campo
		fieldName := strings.ToLower(field.Names[0].Name)

		// Obtém o tipo do campo e o converte para o tipo de dado MySQL
		fieldType := getMySQLType(field.Type)

		sb.WriteString("  ")
		sb.WriteString(fieldName)
		sb.WriteString(" ")
		sb.WriteString(fieldType)

		if fieldName == "id" {
			sb.WriteString(" AUTO_INCREMENT PRIMARY KEY NOT NULL")
		} else {
			sb.WriteString(" NULL")
		}

		// Adiciona vírgula se não for o último campo
		if i < len(fields)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString(");\n\n")
	sb.WriteString("DROP TABLE IF EXISTS ")
	sb.WriteString(strings.ToLower(structName))
	sb.WriteString(";\n")
	return sb.String()
}

// Mapeia o tipo Go para o tipo de dado MySQL
func getMySQLType(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		fmt.Println(t.Name)
		switch t.Name {
		case "int", "int8", "int16", "int32":
			return "INT"
		case "int64":
			return "BIGINT"
		case "uint", "uint8", "uint16", "uint32", "uint64":
			return "INT UNSIGNED"
		case "float32", "float64":
			return "FLOAT"
		case "string":
			return "VARCHAR(255)"
		case "bool":
			return "BOOLEAN"
		case "time.Time":
			return "DATETIME"
		default:
			return "TEXT" // Tipo padrão para tipos desconhecidos
		}
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok && x.Name == "time" && t.Sel.Name == "Time" {
			return "DATETIME"
		}
	case *ast.ArrayType:
		return "JSON" // Usa JSON para armazenar arrays
	case *ast.StarExpr: // Adiciona suporte para ponteiros de string
		if ident, ok := t.X.(*ast.Ident); ok && ident.Name == "string" {
			return "VARCHAR(255)"
		}
	case *ast.MapType:
		return "JSON" // Usa JSON para armazenar maps
	}
	return "TEXT" // Tipo padrão para tipos desconhecidos
}
