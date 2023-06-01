package analyzer

import (
	"go/ast"

	"../../../internal/utils"
)

type Analyzer struct {
	ast *ast.File
	findings []string
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(scope string) {
	a.ast = utils.GenerateAST(scope)
	a.findings = []string{}
	a.analyze()
}