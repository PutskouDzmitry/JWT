package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMessage(t *testing.T) {
	require := require.New(t)
	expected := "1"
	actual := "1"
	require.Equal(expected, actual)
}
