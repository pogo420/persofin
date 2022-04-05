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

// env variables
const TEST_MODE_ENV = "PSF_TEST_MODE" // for persofin test mode
