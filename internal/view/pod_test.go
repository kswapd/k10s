package view_test

import (
	"context"
	"testing"

	"github.com/kswapd/k13s/internal"
	"github.com/kswapd/k13s/internal/client"
	"github.com/kswapd/k13s/internal/config"
	"github.com/kswapd/k13s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestPodNew(t *testing.T) {
	po := view.NewPod(client.NewGVR("v1/pods"))

	assert.Nil(t, po.Init(makeCtx()))
	assert.Equal(t, "Pods", po.Name())
	assert.Equal(t, 25, len(po.Hints()))
}

// Helpers...

func makeCtx() context.Context {
	cfg := config.NewConfig(ks{})
	return context.WithValue(context.Background(), internal.KeyApp, view.NewApp(cfg))
}
