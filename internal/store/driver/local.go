package driver

type LocalDriver struct {
	basePath string
}

func NewLocalDriver(basePath string) *LocalDriver {
	return &LocalDriver{basePath: basePath}
}

func (l *LocalDriver) Name() string {
	return Local
}

func (l *LocalDriver) SupportsSeek() bool {
	return true
}
