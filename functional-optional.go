package main

type FOPattarn struct {
	required string
	optional string
}

type Option func(*FOPattarn) error

func NewFOPattarn(required string, options ...Option) (*FOPattarn, error) {
	fo := FOPattarn{required: required}

	for _, op := range options {
		if err := op(&fo); err != nil {
			// ?: error handling
			continue
		}
	}
	return &fo, nil
}

func option(s string) Option {
	return func(fo *FOPattarn) error {
		fo.optional = s
		return nil
	}
}
