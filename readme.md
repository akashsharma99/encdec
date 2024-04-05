# String Encoder and Decoder CLI
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/akashsharma99/encdec/go)

This is a simple cli tool to encode and decode strings using base64 encoding. Will add support for more algorithms in the future.

Uses Golang with flags package to parse command line arguments.

## Steps to run the cli

1. Clone the repository
2. Run `go build -o encdec` to build the binary

## Usage

```bash
encdec -[action-option] -[algo-option] [string-input]
```

### action options

- `e` : To encode a string
- `d` : To decode a string

### algo options

- `base64` : Base64 encoding

