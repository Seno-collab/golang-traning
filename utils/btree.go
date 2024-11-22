package utils

type Item interface {
	Less(than Item) bool
}

const (
	DefaultFreeSize = 32
)
