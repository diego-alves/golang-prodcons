package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRequest(t *testing.T) {

	query := `query {
		viewer {
			login
		}
	}`

	res := Request(query)
	assert.Equal(t, `{"data":{"viewer":{"login":"diego-alves"}}}`, string(res))

}
