package grifts

import (
	"github.com/alexlovelltroy/mnemonic_buffalo/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
