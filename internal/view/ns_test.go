package view_test

import (
	"testing"

	"github.com/kswapd/k12s/internal/client"
	"github.com/kswapd/k12s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestNSCleanser(t *testing.T) {
	ns := view.NewNamespace(client.NewGVR("v1/namespaces"))

	assert.Nil(t, ns.Init(makeCtx()))
	assert.Equal(t, "Namespaces", ns.Name())
	assert.Equal(t, 7, len(ns.Hints()))
}
