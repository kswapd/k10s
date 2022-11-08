package view_test

import (
	"testing"

	"github.com/kswapd/k13s/internal/client"
	"github.com/kswapd/k13s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestReferenceNew(t *testing.T) {
	s := view.NewReference(client.NewGVR("references"))

	assert.Nil(t, s.Init(makeCtx()))
	assert.Equal(t, "References", s.Name())
	assert.Equal(t, 4, len(s.Hints()))
}
