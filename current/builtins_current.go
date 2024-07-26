package current

import (
	"github.com/refaktor/rye-ebitengine/current/ebiten"
	"github.com/refaktor/rye/env"
	"github.com/refaktor/rye/evaldo"
)

var Builtins_current = map[string]*env.Builtin{}

func RegisterBuiltins(ps *env.ProgramState) {
	evaldo.RegisterBuiltins2(Builtins_current, ps, "current")
	evaldo.RegisterBuiltinsInContext(ebiten.Builtins, ps, "ebiten")
}
