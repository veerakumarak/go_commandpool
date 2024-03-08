package commandbus

import (
	"errors"
)

type Command string

func (c Command) valid() error {
	if c == "" {
		return errors.New("command name can not be empty")
	}

	return nil
}
