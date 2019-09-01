package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestParseMessage(t *testing.T) {
	data:=[]byte(`{"type":1,"body":"/stock=stock_code"}`)
	_,resend, err := ParseMessage(data)

	assert.Nil(t, err)
	assert.Equal(t, true, resend)
}

