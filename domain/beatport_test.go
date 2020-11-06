package domain_test

import (
	"fmt"
	"github.com/jeanmorais/beatport/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenre_ParseKey(t *testing.T) {
	tables := []struct {
		input        domain.Genre
		expectedID   string
		expectedName string
	}{
		{domain.Genre{Key: "afro-house-15"}, "15", "afro-house"},
		{domain.Genre{Key: "progressive-10"}, "10", "progressive"},
		{domain.Genre{Key: "techno-raw-deep-hypnotic-92"}, "92", "techno-raw-deep-hypnotic"},
	}

	for _, tc := range tables {
		id, name := tc.input.ParseKey()
		assert.Equal(t, tc.expectedID, id)
		assert.Equal(t, tc.expectedName, name)
	}
}

func TestGenre_Validate(t *testing.T) {
	tables := []struct {
		input       domain.Genre
		expectedErr error
	}{
		{domain.Genre{Key: "afro-house-150000"}, fmt.Errorf("Invalid argument: genreKey [afro-house-150000]")},
		{domain.Genre{Key: "progressive-10"}, nil},
	}

	for _, tc := range tables {
		err := tc.input.Validate()
		assert.Equal(t, tc.expectedErr, err)
	}
}
