package ui

import (
	"fmt"
	"strings"

	"github.com/derailed/tview"
	"github.com/kswapd/k13s/internal/config"
)

/*
// LogoSmall K9s small log.

	var LogoSmall = []string{
		` ____  __.________       `,
		`|    |/ _/   __   \______`,
		`|      < \____    /  ___/`,
		`|    |  \   /    /\___ \ `,
		`|____|__ \ /____//____  >`,
		`        \/            \/ `,
	}

// LogoBig K9s big logo for splash page.

	var LogoBig = []string{
		` ____  __.________      _________ .____    .___ `,
		`|    |/ _/   __   \_____\_   ___ \|    |   |   |`,
		`|      < \____    /  ___/    \  \/|    |   |   |`,
		`|    |  \   /    /\___ \\     \___|    |___|   |`,
		`|____|__ \ /____//____  >\______  /_______ \___|`,
		`        \/            \/        \/        \/    `,
	}

	var LogoSmall = []string{
		` ____  __.________       `,
		`|    |/ _/   __   \______`,
		`|      < \____    /  ___/`,
		`|    |  \   /    /\___ \ `,
		`|____|__ \ /____//____  >`,
		`        \/            \/ `,
	}
*/
var LogoBig = []string{
	` _______  _______  __    _  _______  _______  ___   _______ `,
	`|       ||       ||  |  | ||       ||       ||   | |       |`,
	`|    ___||    ___||   |_| ||    ___||  _____||   | |  _____|`,
	`|   | __ |   |___ |       ||   |___ | |_____ |   | | |_____ `,
	`|   ||  ||    ___||  _    ||    ___||_____  ||   | |_____  |`,
	`|   |_| ||   |___ | | |   ||   |___  _____| ||   |  _____| |`,
	`|_______||_______||_|  |__||_______||_______||___| |_______|`,
}

var LogoSmall = []string{
	` _______                   `,
	`(_______)                  `,
	` _   ___ _____ ____  _____ `,
	`| | (_  | ___ |  _ \| ___ |`,
	`| |___) | ____| | | | ____|`,
	` \_____/|_____)_| |_|_____)`,
}

// Splash represents a splash screen.
type Splash struct {
	*tview.Flex
}

// NewSplash instantiates a new splash screen with product and company info.
func NewSplash(styles *config.Styles, version string) *Splash {
	s := Splash{Flex: tview.NewFlex()}
	s.SetBackgroundColor(styles.BgColor())

	logo := tview.NewTextView()
	logo.SetDynamicColors(true)
	logo.SetTextAlign(tview.AlignCenter)
	s.layoutLogo(logo, styles)

	vers := tview.NewTextView()
	vers.SetDynamicColors(true)
	vers.SetTextAlign(tview.AlignCenter)
	s.layoutRev(vers, version, styles)

	s.SetDirection(tview.FlexRow)
	s.AddItem(logo, 10, 1, false)
	s.AddItem(vers, 1, 1, false)

	return &s
}

func (s *Splash) layoutLogo(t *tview.TextView, styles *config.Styles) {
	logo := strings.Join(LogoBig, fmt.Sprintf("\n[%s::b]", styles.Body().LogoColor))
	fmt.Fprintf(t, "%s[%s::b]%s\n",
		strings.Repeat("\n", 2),
		styles.Body().LogoColor,
		logo)
}

func (s *Splash) layoutRev(t *tview.TextView, rev string, styles *config.Styles) {
	fmt.Fprintf(t, "[%s::b]Revision [red::b]%s", styles.Body().FgColor, rev)
}
