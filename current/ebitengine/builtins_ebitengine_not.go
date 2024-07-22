//go:build !b_ebiten
// +build !b_ebiten

package ebiten

import (
	"github.com/refaktor/rye/env"
)

var Builtins_ebiten = map[string]*env.Builtin{}
