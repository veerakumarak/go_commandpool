package main

import (
	"encoding/json"
	"fmt"
	"github.com/veerakumarak/go-commandbus"
	"time"
)

// Define commands
const (
	PrintCommand commandbus.Command = "print_command"
	GreetCommand                    = "greet_command"
)

// Define command payload
type PrintCommandPayload struct {
	Message string
}

type GreetCommandPayload struct {
	Name string
}

// Define command handlers
func PrintCommandHandler(payload json.RawMessage) (json.RawMessage, error) {
	// parse payload
	var printCommandPayload PrintCommandPayload
	if err := json.Unmarshal(payload, &printCommandPayload); err != nil {
		return nil, err
	}

	// execute the command
	time.Sleep(time.Second * 2)
	fmt.Println(printCommandPayload)

	// return the result
	return nil, nil
}

func GreetCommandHandler(payload json.RawMessage) (json.RawMessage, error) {
	// parse payload
	var greetCommandPayload GreetCommandPayload
	if err := json.Unmarshal(payload, &greetCommandPayload); err != nil {
		return nil, err
	}

	// execute the command
	fmt.Println(greetCommandPayload)

	// return the result
	return nil, nil
}

func main() {
	//bus := commandbus.New("default-command-bus")
	bus := commandbus.NewWithOptions("command-bus-2-100", 2, 100)

	// Register command handlers
	err := bus.Register(PrintCommand, PrintCommandHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bus.Register(GreetCommand, GreetCommandHandler)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Execute commands synchronously
	payload, _ := json.Marshal(&PrintCommandPayload{Message: "Hello World"})
	_, _ = bus.Execute(PrintCommand, payload)

	payload, _ = json.Marshal(&GreetCommandPayload{Name: "Alice"})
	_, _ = bus.Execute(GreetCommand, payload)

	// Execute commands asynchronously
	payload, _ = json.Marshal(&PrintCommandPayload{Message: "Hello World"})
	_ = bus.Dispatch(PrintCommand, payload)

	payload, _ = json.Marshal(&GreetCommandPayload{Name: "Alice"})
	_ = bus.Dispatch(GreetCommand, payload)

	bus.Shutdown()
}
