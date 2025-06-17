# events Package

The `events` package provides a robust and flexible event handling mechanism for your Go applications. It emphasizes a "build-it-yourself" approach, allowing you to clearly define and manage event listeners upfront.

---

## Installation

To use the `events` package, simply install it with `go get`:

```bash
go get github.com/ogiusek/events
```

---

## Core Concepts

The package revolves around two main components:

* **`Builder`**: This is your starting point for configuring and setting up your event system. It uses a fluent API to chain configuration options.
* **`Events`**: This is the compiled event dispatcher, ready to `Emit` events and notify your registered listeners.

---

## Getting Started

Here's a basic example of how to set up an event dispatcher and emit events:

```go
package main

import (
	"fmt"
	"github.com/ogiusek/events"
)

func main() {
	e := events.NewBuilder().
		// Listen for integer events
		events.Listen(func(num int) {
			fmt.Printf("Received int: %d\n", num)
		}).
		// Listen for string events
		events.Listen(func(text string) {
			fmt.Printf("Received string: %s\n", text)
		}).
		Build()

	events.Emit(e, "Hello, events!")
	events.Emit(e, 123)
}
```

---

## The Builder

The `Builder` provides several methods to customize your event system:

### `NewBuilder()`

This is the constructor for the `Builder`. Always start your event system configuration by calling this function. Failing to use the constructor will result in a panic.

```go
builder := events.NewBuilder()
```

### `Listen[Event event](b Builder, listener func(Event)) Builder`

This function registers a listener for a specific event type. The `listener` function's signature must match the `Event` type. You can register multiple listeners for the same event type.

```go
// Listen for int events
builder = events.Listen(builder, func(num int) {
	fmt.Printf("Number: %d\n", num)
})

// Listen for string events
builder = events.Listen(builder, func(text string) {
	fmt.Printf("Text: %s\n", text)
})
```

### `Wrap(wrapped func(Builder) Builder) Builder`

The `Wrap` method allows for a more organized and modular way to register listeners or apply other builder configurations. It takes a function that receives the current `Builder` and returns a modified `Builder`. This is particularly useful for grouping related listener registrations or applying common patterns.

```go
// Grouping string listeners
e := events.NewBuilder().
	Wrap(func(b events.Builder) events.Builder {
		return events.Listen(b, func(text string) { print(text) })
	}).
	Wrap(func(b events.Builder) events.Builder {
		return events.Listen(b, func(text string) { print("Another string listener: " + text) })
	}).
	Build()
```

### `GoroutinePerListener(use bool) Builder`

This method controls whether each listener should be executed in its own goroutine when an event is emitted.

* If `use` is `true`, each listener will run concurrently in a new goroutine. This is useful for long-running listeners that shouldn't block the event emission process.
* If `use` is `false` (default), listeners are executed sequentially in the same goroutine as the `Emit` call.

```go
builder := events.NewBuilder().GoroutinePerListener(true) // Listeners will run in separate goroutines
```

### `Build() Events`

After you've configured your `Builder` with all the desired listeners and settings, call `Build()` to finalize the event system and get an `Events` instance.

```go
eventSystem := events.NewBuilder().
	// ... add listeners and configurations
	Build()
```

---

## Emitting Events

Once you have an `Events` instance, you can use the `Emit` function to dispatch events.

### `Emit[Event event](e Events, event Event)`

The `Emit` function dispatches an `event` to all registered listeners of that specific `Event` type.

```go
events.Emit(eventSystem, "This is a string event")
events.Emit(eventSystem, 42) // This is an integer event
```

## Contributing

Contributions are welcome\! Please feel free to open issues or submit pull requests on the [GitHub repository](https://www.google.com/search?q=https://github.com/ogiusek/relay).
