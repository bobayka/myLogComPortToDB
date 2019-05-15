package comPort

import "regexp"

var validMessage = regexp.MustCompile(`-?\d*\.\d*`)
var newLine = regexp.MustCompile(`:((-?\d*\.\d*,){6})\s`)
