# snipit
An easy to use, interactive snippet management tool.

## Installing
To install, simply run `go install github.com/pyrod3v/snipit/cmd/snipit@latest` or clone this repository and run `go install ./...`.

## How to use
Simply type `snipit` in your terminal and a menu with all your snippets will appear.  
If you select one, it will open a menu where you can run, copy, print, edit or delete that snippet.  
Running `snipit` with a parameter will try to open the management menu for that snippet. If the snippet doesn't exist, it will open your editor for you to create it.  
Run the program with `-h` or `--help` to show the help message.  
You can also use the following subcommands, followed by a snippet's name:
- `run`: Run the provided snippet. Additional parameters will be passed to the snippet.
- `copy`: Copy the provided snippet.
- `print`: Print the provided snippet.
- `edit`: Edit the provided snippet.
- `delete`: Delete the provided snippet.

## Configuration
The application's configuration is stored in `HOME/.snipit/config.yaml`
You can specify the following config keys:
- `Editor`: The editor to use when creating a new snippet. Defaults to your $EDITOR environment variable, or `nano` if it isn't set.
- `SnippetsDir`: The directory to store the snippets in. Defaults to `HOME/.snipit/snippets`.
You can modify config values by running `snipit -c <key> <value>` (Using `--config` instead of `-c` is supported too).  
You can also have a directory-specific config in `./.snipit/config.yaml`.  

## Contributing
All sorts of contributions are welcome. To contribute:
1. Fork this repository
2. Create your feature branch
3. Commit and push your changes
4. Submit a pull request

Please make sure your commit messages are meaningful.
