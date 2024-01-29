package domain_errors

import "errors"

var ErrItemNotEnough = errors.New("item not enough")
var ErrItemNotFound = errors.New("item not found")
