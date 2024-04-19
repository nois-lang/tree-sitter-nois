package tree_sitter_nois_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-nois"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_nois.Language())
	if language == nil {
		t.Errorf("Error loading Nois grammar")
	}
}
