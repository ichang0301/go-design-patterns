# 3-2-1_incompatible_interface

## Requirements and acceptance criteria

- Create an adapter object that implements the `ModernPrinter` interface
- The new adapter object must contain an instance of the `LegacyPrinter` interface
- When using `ModernPrinter`, it must call the `LegacyPrinter` interface under the hood, prefixing it with the text `Adapter`
