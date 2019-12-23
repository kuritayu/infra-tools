package lstar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveUser(t *testing.T) {
	actual := ResolveUser(501)
	expected := "kuritayu"
	assert.Equal(t, expected, actual)
}

func TestNotResolveUser(t *testing.T) {
	actual := ResolveUser(9999)
	expected := "9999"
	assert.Equal(t, expected, actual)
}

func TestResolveGroup(t *testing.T) {
	actual := ResolveGroup(20)
	expected := "staff"
	assert.Equal(t, expected, actual)
}

func TestNotResolveGroup(t *testing.T) {
	actual := ResolveGroup(9999)
	expected := "9999"
	assert.Equal(t, expected, actual)
}
