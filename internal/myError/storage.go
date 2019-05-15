package myError

import "github.com/pkg/errors"

var (
	ErrNoMatches     = errors.New("no matches")
	MPUDoesntConnect = "mpu doesnt connect"
	MPUHalError      = "hal error"
)
