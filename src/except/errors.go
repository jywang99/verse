package except

type ErrType int

const (
    AuthErr ErrType = iota
    ForbiddenErr
    DbErr
)

type HandledError struct {
    Type ErrType
    Message string
}

func NewHandledError(t ErrType, msg string) HandledError {
    return HandledError{Type: t, Message: msg}
}

func (e HandledError) Error() string {
    return e.Message
}

