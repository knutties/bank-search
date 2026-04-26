package search

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

// ErrNotFound is returned by Lookup when no branch exists for the given code.
var ErrNotFound = errors.New("ifsc code not found")

// Lookup returns the Branch indexed under the given IFSC code, or
// ErrNotFound if no such document exists. The lookup is case-insensitive.
func (b *bleveSearcher) Lookup(code string) (*Branch, error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if code == "" {
		return nil, ErrNotFound
	}

	q := bleve.NewDocIDQuery([]string{code})
	sr := bleve.NewSearchRequestOptions(q, 1, 0, false)
	sr.Fields = []string{"*"}

	res, err := b.idx.Search(sr)
	if err != nil {
		return nil, fmt.Errorf("bleve lookup: %w", err)
	}
	if res.Total == 0 {
		return nil, ErrNotFound
	}
	hit := res.Hits[0]
	br := branchFromFields(hit.Fields)
	br.IFSC = hit.ID
	return br, nil
}
