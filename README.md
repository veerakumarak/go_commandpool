# Command Bus Library in Go

This is a lightweight and flexible Command Bus library for Go, inspired by the Command Bus design pattern. It allows you to decouple the sender of a request (command) from the handler of that request.

## Features

- **Decoupling**: Separates the sender of a request (command) from its handler, enabling better separation of concerns and easier testing.
- **Flexibility**: Provides a simple and extensible API for defining commands and their handlers.
- **Concurrency**: Handles commands asynchronously, making it suitable for high-performance applications.

## Installation

```bash
go get github.com/veerakumarak/go-commandbus
```

## Usage

<pre><div class="dark bg-gray-950 rounded-md"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span></span><span class="" data-state="closed"><button class="flex gap-1 items-center"><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" class="icon-sm"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 4C10.8954 4 10 4.89543 10 6H14C14 4.89543 13.1046 4 12 4ZM8.53513 4C9.22675 2.8044 10.5194 2 12 2C13.4806 2 14.7733 2.8044 15.4649 4H17C18.6569 4 20 5.34315 20 7V19C20 20.6569 18.6569 22 17 22H7C5.34315 22 4 20.6569 4 19V7C4 5.34315 5.34315 4 7 4H8.53513ZM8 6H7C6.44772 6 6 6.44772 6 7V19C6 19.5523 6.44772 20 7 20H17C17.5523 20 18 19.5523 18 19V7C18 6.44772 17.5523 6 17 6H16C16 7.10457 15.1046 8 14 8H10C8.89543 8 8 7.10457 8 6Z" fill="currentColor"></path></svg></button></span></div><div class="p-4 overflow-y-auto"><code class="!whitespace-pre hljs language-go">
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

// Define command payloads
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
    // Supports sequential execution
	bus := commandbus.New("default-command-bus")
  
    // Supports concurrent execution of one or more tasks based on maxWorkers
	//bus := commandbus.NewWithOptions("command-bus-2-100", 2, 100)

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
	_ = bus.Execute(PrintCommand, payload)

	payload, _ = json.Marshal(&GreetCommandPayload{Name: "Alice"})
	_ = bus.Execute(GreetCommand, payload)

    // Execute commands asynchronously
	payload, _ := json.Marshal(&PrintCommandPayload{Message: "Hello World"})
	_ = bus.Dispatch(PrintCommand, payload)

	payload, _ = json.Marshal(&GreetCommandPayload{Name: "Alice"})
	_ = bus.Dispatch(GreetCommand, payload)


	bus.Shutdown()
}
</code></div></div></pre>

### The output will be

```
{Hello World}
{Alice}
{Alice}
{Hello World}
```


## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request for any improvements or new features you'd like to see.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
