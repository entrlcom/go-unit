package unit

const (
	B = 1 << (10 * iota) //nolint:gomnd // OK.
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
)
