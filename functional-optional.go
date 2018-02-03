package main

type FOPattarn struct {
	required string
	optional string
}

type FOOption func(*FOPattarn) error

func NewFOPattarn(required string, options ...FOOption) (*FOPattarn, error) {
	fo := FOPattarn{required: required}

	for _, op := range options {
		if err := op(&fo); err != nil {
			// ?: error handling
			continue
		}
	}
	return &fo, nil
}

func optional(s string) FOOption {
	return func(fo *FOPattarn) error {
		fo.optional = s
		return nil
	}
}
