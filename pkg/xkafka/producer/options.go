package producer

type Options struct {
	Topic        string
	EventHandler EventHandlerFunc
	IsDebug      bool
}

type Option func(opts *Options)

func EventHandler(handler EventHandlerFunc) Option {
	return func(opts *Options) {
		opts.EventHandler = handler
	}
}

func IsDebug(isDebug bool) Option {
	return func(opts *Options) {
		opts.IsDebug = isDebug
	}
}
