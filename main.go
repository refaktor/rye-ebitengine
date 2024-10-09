package main

import (
	/*RYEGEN: BEGIN IMPORTS*/
	"github.com/refaktor/rye-ebitengine/ryegen_bindings/github_com_hajimehoshi_ebiten_v2"
	/*RYEGEN: END IMPORTS*/

	"github.com/refaktor/rye/env"
	"github.com/refaktor/rye/evaldo"
	"github.com/refaktor/rye/runner"
)

func main() {
	runner.DoMain(func(ps *env.ProgramState) {
		/*RYEGEN: BEGIN BUILTINS*/
		evaldo.RegisterBuiltinsInContext(github_com_hajimehoshi_ebiten_v2.Builtins, ps, "ebiten")
		/*RYEGEN: END BUILTINS*/
	})
}
