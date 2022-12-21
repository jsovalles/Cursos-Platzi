package utils

import "go.uber.org/fx"

var Module = fx.Provide(NewEnv, NewDatabase, NewUtil)
