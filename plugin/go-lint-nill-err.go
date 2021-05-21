// This must be package main
package main

import (
	analyzer "github.com/gvych/go-lint-nil-err-not-reversed/pkg/analyzer"
	"golang.org/x/tools/go/analysis"
)

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
