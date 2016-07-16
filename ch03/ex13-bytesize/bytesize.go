// Package bytesize provides constants to describe unit of byte size.
package bytesize

const kb = 1024

const (
	KiB = kb << (10 * iota) // 1024
	MiB                     // 1048576
	GiB                     // 1073741824
	TiB                     // 1099511627776
	PiB                     // 1125899906842624
	EiB                     // 1152921504606846976
	ZiB                     // 180591620717411303424
	YiB                     // 12089258196146297470676
)
