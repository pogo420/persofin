package commons

import (
	cm "persofin/src/command"
)

const PROMT = "persofin>"

var Commands = map[string]cm.BaseCommand{
	"bye": &cm.ExitCommand{},
}
