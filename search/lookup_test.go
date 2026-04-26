package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLookup_Found(t *testing.T) {
	branches := []*Branch{
		{IFSC: "HDFC0000001", BankCode: "HDFC", BankName: "HDFC Bank",
			Branch: "ANDHERI WEST", City: "MUMBAI", State: "MAHARASHTRA",
			NEFT: true, RTGS: true, IMPS: true, UPI: true},
	}
	s, err := NewMemorySearcher(branches)
	require.NoError(t, err)
	t.Cleanup(func() { _ = s.Close() })

	br, err := s.Lookup("HDFC0000001")
	require.NoError(t, err)
	assert.Equal(t, "HDFC0000001", br.IFSC)
	assert.Equal(t, "HDFC", br.BankCode)
	assert.Equal(t, "ANDHERI WEST", br.Branch)
	assert.True(t, br.NEFT)
}

func TestLookup_CaseInsensitive(t *testing.T) {
	branches := []*Branch{
		{IFSC: "HDFC0000001", BankCode: "HDFC", BankName: "HDFC Bank",
			Branch: "ANDHERI WEST", City: "MUMBAI"},
	}
	s, err := NewMemorySearcher(branches)
	require.NoError(t, err)
	t.Cleanup(func() { _ = s.Close() })

	br, err := s.Lookup("hdfc0000001")
	require.NoError(t, err)
	assert.Equal(t, "HDFC0000001", br.IFSC)
}

func TestLookup_NotFound(t *testing.T) {
	branches := []*Branch{
		{IFSC: "HDFC0000001", BankCode: "HDFC", BankName: "HDFC Bank",
			Branch: "ANDHERI WEST", City: "MUMBAI"},
	}
	s, err := NewMemorySearcher(branches)
	require.NoError(t, err)
	t.Cleanup(func() { _ = s.Close() })

	_, err = s.Lookup("ZZZZ0000000")
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestLookup_EmptyCode(t *testing.T) {
	s, err := NewMemorySearcher(nil)
	require.NoError(t, err)
	t.Cleanup(func() { _ = s.Close() })

	_, err = s.Lookup("")
	assert.ErrorIs(t, err, ErrNotFound)
}
