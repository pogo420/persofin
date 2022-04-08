package command

import (
	dbi "persofin/src/db_interface"
)

type BaseCommand interface {
	Description() string
	Execute(dbi.BaseDbInterface, string) int
}
