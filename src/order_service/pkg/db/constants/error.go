package constants

import (
	"monorepo/src/libs/constants"
)

const (
	// ErrRowsAffectedZero indicates nothing has been affected after sql query
	ErrRowsAffectedZero = constants.Sentinel("no rows affected by sql command")
	ErrRecordNotFound   = constants.Sentinel("record not found")
	// ErrRecordNotFound indicates that record not found in database
)
