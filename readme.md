# colorful

`colorful` is a simple Go package that provides colored terminal output, built as a wrapper around the [`aurora`](https://github.com/logrusorgru/aurora) package. It allows global control over whether color output is enabled.

## Features

- Globally toggle terminal color output
- Thread-safe color operations
- Simple, minimal API
- Support for multiple colors and styles
- Predefined color renderers for security severity levels

## Installation

```bash
go get github.com/fishfinal/colorful
```

## Usage

### Colors enabled by default

```go
// Colors are enabled by default
colorful.Colorful.Red("This text will be displayed in red")
```

### Disabling color output

```go
// Disable color output
colorful.SetEnableColors(false)
colorful.Colorful.Green("This text will not be displayed in green")
```

### Re-enabling colors

```go
// Re-enable colors
colorful.SetEnableColors(true)
colorful.Colorful.Blue("This text will be displayed in blue")
```

### Severity color renderers

The package ships with ready-to-use renderers for common severity levels:

```go
colorful.SeverityCritical("critical") // purple
colorful.SeverityHigh("high")         // red
colorful.SeverityMedium("medium")     // orange
colorful.SeverityLow("low")           // yellow
colorful.SeverityInfo("info")         // cyan
colorful.SeverityUnknown("unknown")   // gray
```

### Custom color helpers

```go
colorful.Purple("purple text") // ANSI index 128
colorful.Orange("orange text") // ANSI index 208
colorful.Gray("gray text")     // ANSI index 7
```

### Conditional coloring

Use `ReviseColorfully` to keep a value colored or reset it based on a boolean flag:

```go
colorful.ReviseColorfully(colorful.Colorful.Red("text"), true)  // "\x1b[31mtext\x1b[0m"
colorful.ReviseColorfully(colorful.Colorful.Red("text"), false) // "text"
```

## API Documentation

### Variables

```go
var Colorful aurora.Aurora
```

A globally available color handler instance that provides various color methods such as `Red()`, `Green()`, `Bold()`, etc.

### Functions

```go
func SetEnableColors(enable bool)
```

Sets whether color output is enabled:

- `enable=true`: enable color output
- `enable=false`: disable color output

```go
func ReviseColorfully(value aurora.Value, colorfully bool) aurora.Value
```

Returns the value unchanged when `colorfully` is `true`, or resets it to plain text when `false`.

```go
func Purple(text string) aurora.Value
func Orange(text string) aurora.Value
func Gray(text string) aurora.Value
```

Convenience renderers for specific colors.

### Types

```go
type SeverityColorfully func(text string) aurora.Value
```

A function type for severity-based color renderers. The following package-level variables implement it:

- `SeverityCritical`
- `SeverityHigh`
- `SeverityMedium`
- `SeverityLow`
- `SeverityInfo`
- `SeverityUnknown`

## Supported Color Methods

Through the `Colorful` variable, the following commonly used methods are available (for the full list, refer to the [aurora documentation](https://github.com/logrusorgru/aurora)):

- **Colors**: `Red()`, `Green()`, `Blue()`, `Yellow()`, `Cyan()`, `Magenta()`, `White()`, `Black()`
- **Background colors**: `BgRed()`, `BgGreen()`, `BgBlue()`, etc.
- **Styles**: `Bold()`, `Italic()`, `Underline()`, `Inverse()`

## Best Practices

1. Set color preferences during program initialization.
2. Dynamically enable/disable colors based on terminal capabilities.
3. For long-running programs, consider wrapping `SetEnableColors` with atomic operations.

## License

See the repository for license information.
