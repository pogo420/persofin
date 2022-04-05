// File containing static data
package commons

import (
	cm "persofin/src/command"
)

// promt string
const PROMT = "persofin>"

// All commands to struct mapping
var Commands = map[string]cm.BaseCommand{
	"bye": &cm.ExitCommand{},
}
