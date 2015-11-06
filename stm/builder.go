package stm

type Builder interface {
	// Content() string
	Add(interface{}) Builder
	// AddWithErr(url interface{}) (Builder, error)
	// location() *Location

	isFull() bool
	isFinalized() bool

	finalize()
	write()
	run()
}
