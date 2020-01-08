package cocafile

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	. "github.com/phodal/coca/languages/java"
	ignore "github.com/sabhiram/go-gitignore"
	"os"
	"path/filepath"
)

func GetJavaFiles(codeDir string) []string {
	return GetFilesWithFilter(codeDir, JavaCodeFileFilter)
}

func GetFilesWithFilter(codeDir string, filter func(path string) bool) []string {
	files := make([]string, 0)
	gitIgnore, err := ignore.CompileIgnoreFile(".gitignore")
	if err != nil {
		//fmt.Println(err)
	}

	fi, err := os.Stat(codeDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if fi.Mode().IsRegular() {
		files = append(files, codeDir)
		return files
	}

	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if gitIgnore != nil {
			if gitIgnore.MatchesPath(path) {
				return nil
			}
		}

		if filter(path) {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func GetJavaTestFiles(codeDir string) []string {
	return GetFilesWithFilter(codeDir, JavaTestFileFilter)
}

func ProcessJavaFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewJavaParser(stream)
	return parser
}

func ProcessJavaString(code string) *JavaParser {
	is := antlr.NewInputStream(code)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewJavaParser(stream)
	return parser
}
