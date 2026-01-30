# GoLangCLI File Manager

A powerful, feature-rich command-line file manager written in Go.

## Features

- **File Operations**: Create, delete, copy, move, and rename files and folders.
- **Tree Visualization**: View directory structure with the `tree` command.
- **Content Search**: Search for files by name or search *inside* file contents.
- **Interactive Mode**: Navigate your filesystem with a terminal UI (TUI).
- **File Hashing**: Calculate MD5, SHA1, and SHA256 checksums.
- **Compression**: Compress and extract files (zip).

## Installation

```bash
git clone https://github.com/ooizhenyi/GoLangCLI.git
cd GoLangCLI
go mod tidy
go build -o filemanager .
```

## Usage

Run the tool using `./filemanager` or `go run main.go`.

### Basic Commands

| Command | Usage | Description |
|---------|-------|-------------|
| `cf` | `cf [name]` | Create a folder |
| `dlt` | `dlt [name]` | Delete a file or folder |
| `copyfile` | `copyfile [src] [dst]` | Copy a file |
| `mv` | `mv [src] [dst]` | Move a file or folder |
| `rename` | `rename [old] [new]` | Rename a file or folder |
| `list` | `list` | List files in current directory |
| `properties` | `properties [file]` | View file properties |

### Advanced Features

#### Tree View
Visualize the directory structure recursively.
```bash
filemanager tree .
filemanager tree . --depth 2
```

#### Search
Search for files by name or content.
```bash
# Search by filename
filemanager search "config"

# Search by content
filemanager search "TODO" --content
```

#### Interactive Mode
Launch the Terminal User Interface (TUI) to browse files.
```bash
filemanager interactive
```
Use arrow keys to navigate, Enter to open folders, and Backspace to go up.

#### File Hashing
Verify file integrity.
```bash
filemanager hash myfile.zip
filemanager hash myfile.zip --algo md5
```

## Development

Run tests:
```bash
go test ./...
```