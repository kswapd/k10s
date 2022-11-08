package view_test

import (
	"testing"

	"github.com/kswapd/k13s/internal/client"
	"github.com/kswapd/k13s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestStatefulSetNew(t *testing.T) {
	s := view.NewStatefulSet(client.NewGVR("apps/v1/statefulsets"))

	assert.Nil(t, s.Init(makeCtx()))
	assert.Equal(t, "StatefulSets", s.Name())
	assert.Equal(t, 12, len(s.Hints()))
}
