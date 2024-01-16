package domain_errors

import "errors"

var CannotAcceptJob = errors.New("cannot accept job")
var CannotStartJob = errors.New("cannot start job")
var CannotCompleteJob = errors.New("cannot complete job")
