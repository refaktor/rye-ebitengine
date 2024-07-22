package current

import (
	"github.com/refaktor/rye-ebitengine/current/ebitengine"
	"github.com/refaktor/rye/env"
	"github.com/refaktor/rye/evaldo"
)

var Builtins_current = map[string]*env.Builtin{}

func RegisterBuiltins(ps *env.ProgramState) {
	evaldo.RegisterBuiltins2(Builtins_current, ps, "current")
	evaldo.RegisterBuiltinsInContext(ebitengine.Builtins_ebitengine, ps, "ebitengine")
}
