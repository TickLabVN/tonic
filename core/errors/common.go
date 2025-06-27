package errors

import "errors"

var ErrUnimplemented = errors.New("unimplemented")
var ErrParseValidateTag = errors.New("failed to parse validate tag")