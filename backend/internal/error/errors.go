package error

import "errors"

var ErrNegativePrice = errors.New("price cannot be < 0")
var ErrIncorrectData = errors.New("data format should be MM-YYYY")
var ErrNoService = errors.New("user dont have subscription on this service")
var ErrInvalidDataPeriod = errors.New("time start need to be earlier than end time")
var ErrUserWithoutSub = errors.New("user dont have sub at this time")
