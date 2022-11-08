package dialog

import (
	"testing"

	"github.com/derailed/tview"
	"github.com/kswapd/k12s/internal/config"
	"github.com/kswapd/k12s/internal/ui"
	"github.com/stretchr/testify/assert"
)

func TestConfirmDialog(t *testing.T) {
	a := tview.NewApplication()
	p := ui.NewPages()
	a.SetRoot(p, false)

	ackFunc := func() {
		assert.True(t, true)
	}
	caFunc := func() {
		assert.True(t, true)
	}
	ShowConfirm(config.Dialog{}, p, "Blee", "Yo", ackFunc, caFunc)

	d := p.GetPrimitive(confirmKey).(*tview.ModalForm)
	assert.NotNil(t, d)

	dismissConfirm(p)
	assert.Nil(t, p.GetPrimitive(confirmKey))
}
