package driver

const (
	Local = "local"
)

type Driver interface {
	Name() string
	SupportsSeek() bool
}
