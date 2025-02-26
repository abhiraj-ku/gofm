## gofm - simple go implemetation of linux file management commands

## Usage

```
$ gofm [command] [arg1] [arg2]
```

Use the commands like their Linux counterparts but with `gofm` as the prefix.

## Examples

```
$ gofm ls           # List files in the current directory
$ gofm pwd          # Get the current working directory
$ gofm cat abc.txt  # View the contents of a file
$ gofm mv abc.txt def.txt  # Move a file
$ gofm rm abc.txt   # Delete a file
```

## Available Commands

- `ls` – List files in the current directory
- `rm` – Delete a file
- `pwd` – Print the current working directory
- `mv` – Move or rename a file
- `mkdir` – Create a new directory
- `cat` – View the contents of a file
- `touch` – Create an empty file
- `cp` – Copy a file
- `echo` – Print text to the console
- `info` – Get information about the file system

## Installation

1.  Build the executable:

    ```sh
    go build -o gofm
    ```

2.  Move it to `/usr/local/bin` for system-wide usage:

    ```sh
    sudo mv gofm /usr/local/bin/
    ```

3.  Verify installation:

    ```sh
    gofm --help
    ```
