# Pinyin Generator

**[English Version](README_EN.md) | [中文版本](README.md)**

A simple Chinese pinyin generator that converts Chinese text to pinyin and automatically copies the result to clipboard.

## Features

- Convert Chinese text to pinyin (without tone marks)
- Generate full pinyin and initial letter combination
- Automatically copy result to system clipboard

## Installation

### Build from source

```bash
go build -o pinyingen.exe main.go
```

## Usage

### Basic Usage

```bash
pinyingen.exe "中文文本"
```

### Example

```bash
pinyingen.exe "今日头条"
```

**Output:**
```
已复制到剪贴板: jinritoutiaojrtt
```

The generated result includes:
- `jinritoutiao` - Full pinyin
- `jrtt` - Initial letter combination

## Command Line Arguments

```bash
pinyingen.exe [options] "中文文本"
```

Use `pinyingen.exe -h` to view help information.

## Dependencies

- Go 1.25.0+
- [github.com/atotto/clipboard](https://github.com/atotto/clipboard) - Clipboard operations
- [github.com/mozillazg/go-pinyin](https://github.com/mozillazg/go-pinyin) - Pinyin conversion

## License

This project is licensed under the MIT License.