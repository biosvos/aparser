package aparser

type Argument struct {
	Flags []string

	// Required 설정되면
	//   - Default = nil
	Required bool

	Description string

	// Default 설정되면
	//   - Required = false
	Default string

	ConfigKey string

	// Mandatory 설정되면
	//  - Required = true
	//  - Flags = nil
	//  - Default = nil
	//  - ConfigKey = nil
	Mandatory string

	EnvironmentKey string
}

type Options struct {
	ConfigKey      string
	EnvironmentKey string
}

type Option func(*Options)

func WithConfigKey(key string) Option {
	return func(opts *Options) {
		opts.ConfigKey = key
	}
}

func WithEnvKey(key string) Option {
	return func(opts *Options) {
		opts.EnvironmentKey = key
	}
}

func NewOptionalArgument(flags []string, description string, Default string, opts ...Option) *Argument {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	return &Argument{
		Flags:          flags,
		Description:    description,
		Default:        Default,
		ConfigKey:      options.ConfigKey,
		EnvironmentKey: options.EnvironmentKey,

		Mandatory: "",
		Required:  false,
	}
}

func NewRequiredArgument(flags []string, description string, opts ...Option) *Argument {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}

	return &Argument{
		Required: true,

		Flags:          flags,
		Description:    description,
		ConfigKey:      options.ConfigKey,
		EnvironmentKey: options.EnvironmentKey,

		Default:   "",
		Mandatory: "",
	}
}

func NewMandatoryArgument(mandatory string, description string) *Argument {
	return &Argument{
		Required: true,

		Description: description,
		Mandatory:   mandatory,

		Flags:          nil,
		Default:        "",
		ConfigKey:      "",
		EnvironmentKey: "",
	}
}
