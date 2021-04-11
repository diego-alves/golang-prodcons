package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentities(t *testing.T) {

	ids := GetExternalIdentities("")
	assert.Equal(t, 0, ids.TotalCount)

}
