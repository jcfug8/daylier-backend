package persist

type Client interface {
	Ping()

	Close() error
}
