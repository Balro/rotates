package check

type Checkable interface {
	Check(q chan<- string)
	GetInfo() string
}
