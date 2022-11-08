package view_test

import (
	"testing"

	"github.com/kswapd/k13s/internal/config"
	"github.com/kswapd/k13s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestAppNew(t *testing.T) {
	a := view.NewApp(config.NewConfig(ks{}))
	a.Init("blee", 10)

	assert.Equal(t, 11, len(a.GetActions()))
}
