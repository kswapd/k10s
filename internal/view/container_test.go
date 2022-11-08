package view_test

import (
	"testing"

	"github.com/kswapd/k12s/internal/client"
	"github.com/kswapd/k12s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestContainerNew(t *testing.T) {
	c := view.NewContainer(client.NewGVR("containers"))

	assert.Nil(t, c.Init(makeCtx()))
	assert.Equal(t, "Containers", c.Name())
	assert.Equal(t, 18, len(c.Hints()))
}
