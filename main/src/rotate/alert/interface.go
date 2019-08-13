package alert

type Sendable interface {
	Send(info string)
}
