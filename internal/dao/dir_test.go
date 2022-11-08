package dao_test

import (
	"context"
	"testing"

	"github.com/kswapd/k13s/internal"
	"github.com/kswapd/k13s/internal/dao"
	"github.com/stretchr/testify/assert"
)

func TestNewDir(t *testing.T) {
	d := dao.NewDir(nil)
	ctx := context.WithValue(context.Background(), internal.KeyPath, "testdata/dir")
	oo, err := d.List(ctx, "")

	assert.Nil(t, err)
	assert.Equal(t, 2, len(oo))
}
