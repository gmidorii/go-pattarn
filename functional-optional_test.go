package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFOPattarn(t *testing.T) {
	assert := assert.New(t)

	fo, err := NewFOPattarn("req", option("op"))
	if err != nil {
		t.Errorf("create struct error :%v", err)
	}

	assert.Equal("req", fo.required, "should be same")
	assert.Equal("op", fo.optional, "should be same")
}
