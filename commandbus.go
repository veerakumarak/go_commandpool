package commandbus

import (
	"encoding/json"
	"errors"
	"github.com/veerakumarak/go-workerpool"
)

type ICommandBus interface {
	Register(Command, HandlerFunc) error
	Execute(Command, json.RawMessage) (json.RawMessage, error)
	Dispatch(Command, json.RawMessage) error
	Shutdown()
}

type bus struct {
	handlers map[Command]HandlerFunc
	pool     workerpool.IWorkerPool
	quit     bool
}

func New(name string) ICommandBus {
	pool := workerpool.New(name, 1, 1)
	pool.Start()
	return &bus{
		handlers: make(map[Command]HandlerFunc),
		pool:     pool,
	}
}

func NewWithOptions(name string, maxWorkers int, queueSize int) ICommandBus {
	pool := workerpool.New(name, maxWorkers, queueSize)
	pool.Start()
	return &bus{
		handlers: make(map[Command]HandlerFunc),
		pool:     pool,
	}
}

func (b *bus) Register(command Command, fn HandlerFunc) error {
	if err := command.valid(); err != nil {
		return err
	}

	if err := fn.valid(); err != nil {
		return err
	}

	if _, err := b.getHandler(command); err == nil {
		return errors.New("command already registered with a handler")
	}

	b.handlers[command] = fn

	return nil
}

func (b *bus) Execute(command Command, payload json.RawMessage) (json.RawMessage, error) {
	if b.quit {
		return nil, errors.New("shutting down, cannot execute the command")
	}
	if err := b.validate(command, payload); err != nil {
		return nil, err
	}
	return b.execute(command, payload)
}

func (b *bus) Dispatch(command Command, payload json.RawMessage) error {
	if b.quit {
		return errors.New("shutting down, cannot dispatch the command")
	}

	if err := b.validate(command, payload); err != nil {
		return err
	}
	return b.pool.Submit(createAsyncTask(b, command, payload))
}

func (b *bus) Shutdown() {
	b.quit = true
	b.pool.Shutdown()
}

func createAsyncTask(b *bus, command Command, payload json.RawMessage) workerpool.Task {
	return func() {
		_, _ = b.execute(command, payload)
	}
}

func (b *bus) validate(command Command, payload json.RawMessage) error {
	if err := command.valid(); err != nil {
		return err
	}

	return Message(payload).valid()
}

func (b *bus) execute(command Command, payload json.RawMessage) (json.RawMessage, error) {
	if err := b.validate(command, payload); err != nil {
		return nil, err
	}

	h, err := b.getHandler(command)
	if err != nil {
		return nil, err
	}

	return h(payload)
}

func (b *bus) getHandler(command Command) (HandlerFunc, error) {
	if h, ok := b.handlers[command]; ok {
		return h, nil
	}

	return nil, errors.New("handler not found for command")
}
