package search

import (
	"testing"

	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIndexMapping_HasSearchableFields(t *testing.T) {
	m := NewIndexMapping()
	require.NotNil(t, m)

	doc := m.TypeMapping["branch"]
	require.NotNil(t, doc, "branch document mapping must exist")

	expected := []string{"bank_code", "bank_name", "branch", "city", "address",
		"centre", "district", "state"}
	for _, f := range expected {
		field := getField(t, doc, f)
		assert.True(t, field.Index, "field %q must be indexed", f)
	}
}

func TestNewIndexMapping_BoostsMatchSpec(t *testing.T) {
	m := NewIndexMapping()
	doc := m.TypeMapping["branch"]

	cases := map[string]float64{
		"branch":  3.0,
		"city":    2.0,
		"address": 1.0,
	}
	for name, want := range cases {
		got := getField(t, doc, name)
		require.NotNil(t, got, "field %q missing", name)
		assert.InDelta(t, want, fieldBoost(got), 0.0001,
			"field %q boost", name)
	}
}

// helpers
func getField(t *testing.T, doc *mapping.DocumentMapping, name string) *mapping.FieldMapping {
	t.Helper()
	fm, ok := doc.Properties[name]
	if !ok || len(fm.Fields) == 0 {
		t.Fatalf("field %q missing from mapping", name)
	}
	return fm.Fields[0]
}

func fieldBoost(f *mapping.FieldMapping) float64 {
	// Bleve does not surface boost on FieldMapping; the test will read it from
	// a wrapper we add in index.go (see Step 3).
	return FieldBoost(f)
}
