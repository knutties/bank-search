package main

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/blevesearch/bleve/v2"
	"github.com/razorpay/ifsc/v2/ifsc-api/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestEndToEnd_CSVThroughHTTP exercises the full pipeline: build a Bleve
// index on disk from a CSV fixture, open it via search.OpenIndex, mount the
// HTTP router, and assert that /search returns the expected branch.
func TestEndToEnd_CSVThroughHTTP(t *testing.T) {
	csvPath := filepath.Join("cmd", "build-index", "testdata", "sample.csv")
	indexDir := filepath.Join(t.TempDir(), "index")

	require.NoError(t, buildSmallIndexFromCSV(t, csvPath, indexDir))

	s, err := search.OpenIndex(indexDir)
	require.NoError(t, err)
	t.Cleanup(func() { _ = s.Close() })

	srv := httptest.NewServer(newRouter(s, search.Version{Tag: "test"}, ""))
	t.Cleanup(srv.Close)

	resp, err := http.Get(srv.URL + "/search?bank=HDFC&q=andheri")
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body search.SearchResults
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	assert.GreaterOrEqual(t, body.Total, 1)
	assert.Equal(t, "HDFC0000001", body.Results[0].IFSC)
}

// buildSmallIndexFromCSV mirrors what cmd/build-index does, but kept inline
// here so this test does not depend on importing main from another package.
func buildSmallIndexFromCSV(t *testing.T, csvPath, indexDir string) error {
	t.Helper()
	f, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	header, err := r.Read()
	if err != nil {
		return err
	}
	cols, err := search.NewColumnIndex(header)
	if err != nil {
		return err
	}

	idx, err := bleve.New(indexDir, search.NewIndexMapping())
	if err != nil {
		return err
	}
	defer idx.Close()

	for {
		row, err := r.Read()
		if err != nil {
			break
		}
		b, err := search.BranchFromCSVRow(cols, row)
		if err != nil {
			continue
		}
		if err := idx.Index(b.IFSC, b); err != nil {
			return err
		}
	}
	return nil
}
