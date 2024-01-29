package domain_errors

import "errors"

var ErrCannotAcceptJob = errors.New("cannot accept job")
var ErrCannotStartJob = errors.New("cannot start job")
var ErrCannotCompleteJob = errors.New("cannot complete job")
var ErrCannotAddEquipment = errors.New("cannot add equipment")
var ErrCannotAllocateJob = errors.New("cannot Allocate job")
