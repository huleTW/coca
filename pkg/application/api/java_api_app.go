package api

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/infrastructure/ast"
	"github.com/phodal/coca/pkg/infrastructure/ast/api"
	"path/filepath"
)

var allApis []domain.RestApi

type JavaApiApp struct {
}

func (j *JavaApiApp) AnalysisPath(codeDir string, parsedDeps []domain.JClassNode, identifiersMap map[string]domain.JIdentifier, diMap map[string]string) []domain.RestApi {
	files := cocafile.GetJavaFiles(codeDir)
	allApis = nil
	for index := range files {
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Refactoring parse java call: " + displayName)

		parser := ast.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := api.NewJavaApiListener(identifiersMap, diMap)
		listener.AppendClasses(parsedDeps)

		antlr.NewParseTreeWalker().Walk(listener, context)

		currentRestApis := listener.GetClassApis()
		allApis = append(allApis, currentRestApis...)
	}

	return allApis
}
