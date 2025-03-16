package runtime

type SyntaxError struct {
	Reason string
}

func (s SyntaxError) ErrorType() string {
	return "SyntaxError"
}

func (s SyntaxError) Reason_() string {
	return s.Reason
}

type GoError struct {
	Reason string
}

func (g GoError) ErrorType() string {
	return "GoError"
}

func (g GoError) Reason_() string {
	return g.Reason
}