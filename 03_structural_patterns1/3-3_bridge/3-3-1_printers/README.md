# Two printers and two ways of printing for each

## Requirements and acceptance criteria

- A `PrinterAPI` that accepts a message to print
- An implementation of the API that simply prints the message to the console
- An implementation of the API that prints to an `io.Writer` interface
- A `Printer` abstraction with a `Print` method to implement in printing types
- A `normal` printer object, which will implement the `Printer` and the `PrinterAPI` interface
- The `normal` printer will forward the message directly to the implementation
- A `Packt` printer, which will implement the `Printer` abstraction and the `PrinterAPI` interface
- The `Packt` printer will append the message `Message from Packt:` to all prints
