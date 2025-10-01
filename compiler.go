package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr"
	gen "github.com/nikitakocherinskii/ruby-compiler/codegen"
	p "github.com/nikitakocherinskii/ruby-compiler/parser"
)

func main() {

	outPath := "./out/"
	outFilename := "output.ll"
	astFilename := "ast.txt"

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./compiler <source-file>")
		return
	}

	fileName := os.Args[1]
	sourceCode, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Лексический и синтаксический анализ с использованием ANTLR
	input := antlr.NewInputStream(string(sourceCode))
	lexer := p.NewRubyKLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := p.NewRubyKParser(stream)

	// Парсинг программы
	tree := parser.Program()

	// Создание папки out если она не существует
	if err := os.MkdirAll(outPath, 0755); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	// Вывод AST в текстовом формате и сохранение в файл
	fmt.Println("AST:")
	var astBuilder strings.Builder
	astBuilder.WriteString("AST:\n")
	printFormattedAST(tree, parser.GetRuleNames(), parser, "", true, &astBuilder)

	// Сохранение AST в файл
	if astFile, err := os.Create(outPath + astFilename); err == nil {
		astFile.Write([]byte(astBuilder.String()))
		astFile.Close()
		fmt.Printf("AST saved to %s%s\n", outPath, astFilename)
	} else {
		fmt.Println("Error creating AST file:", err)
	}

	// Генерация кода LLVM IR с использованием llir/llvm
	visitor := gen.NewVisitor(true)

	visitor.VisitProgram(tree.(*p.ProgramContext))

	// Вывод результата
	if visitor.Debug {
		fmt.Println(visitor.Module)
	}

	// Сохранение в файл

	if file, err := os.Create(outPath + outFilename); err == nil {
		file.Write([]byte(visitor.Module.String()))
	}
}

func printFormattedAST(tree antlr.Tree, ruleNames []string, parser antlr.Parser, indent string, isLast bool, builder *strings.Builder) {
	var nodeText string

	if ctx, ok := tree.(antlr.RuleContext); ok {
		ruleIndex := ctx.GetRuleIndex()
		if ruleIndex >= 0 && ruleIndex < len(ruleNames) {
			nodeText = ruleNames[ruleIndex]
		} else {
			nodeText = "UnknownRule"
		}
	} else if terminalNode, ok := tree.(antlr.TerminalNode); ok {
		nodeText = fmt.Sprintf("'%s'", terminalNode.GetText())
	} else {
		nodeText = "Unknown"
	}

	var prefix string
	if isLast {
		prefix = "└── "
	} else {
		prefix = "├── "
	}

	line := fmt.Sprintf("%s%s%s\n", indent, prefix, nodeText)
	// fmt.Print(line)
	builder.WriteString(line)

	childCount := tree.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := tree.GetChild(i)
		var childIndent string
		if isLast {
			childIndent = indent + "    "
		} else {
			childIndent = indent + "│   "
		}
		isLastChild := i == childCount-1
		printFormattedAST(child, ruleNames, parser, childIndent, isLastChild, builder)
	}
}
