# ASCII Art Generator

A simple Go-based CLI tool that converts text strings into ASCII art banners using predefined templates.

## Features

- Converts alphanumeric text and symbols into stylized ASCII banners.
- Supports multiple banner styles (standard, shadow, thinkertoy).
- Uses Go's `embed` package to bundle banner files directly into the binary.
- Handles multi-line input via literal `\n` characters.

## Installation

Ensure you have [Go](https://go.dev/doc/install) installed.

1. Clone or navigate to the project directory:
   ```bash
   cd /Go_artify
   ```

2. Build the project:
   ```bash
   go build -o go_artify
   ```

## Usage

Run the program with a single string argument. Use `\n` for new lines.

```bash
./go_artify "Hello\nWorld"
```

### Example

```bash
$ ./go_artify "Hi"
 _    _  _ 
| |  | || |
| |__| || |
|  __  || |
| |  | || |
|_|  |_||_|
           
```

## Banner Files

The tool relies on template files where each character is represented by 8 lines of ASCII art:
- `standard.txt` (Default)
- `shadow.txt`
- `thinkertoy.txt`

Currently, `standard.txt` is embedded by default.

## Testing

Run the included tests to verify functionality:
```bash
go test -v
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
