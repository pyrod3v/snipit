# snipit
An easy to use, interactive snippet management tool.

## Installing
To install, simply run `go install github.com/pyrod3v/snipit/cmd/snipit@latest` or clone this repository and run `go install`.

## How to use
Simply type `snipit` in your terminal and a menu with all your snippets will appear.
If you click on one, it will open a menu where you can run, copy, print, edit or delete that snippet.
Running `snipit` with a parameter will try to open the management menu for that snippet. If the snippet doesn't exist, it will open your editor for you to create it.
Run the program with `-h` or `--help` to show the help message.
All snippets are stored in `$USER/.config/snipit/snippets` on unix-like systems and in `%appdata%\Roaming\snipit\snippets` on Windows.

## Contributing
All sorts of contributions are welcome. To contribute:
1. Fork this repository
2. Create your feature branch
3. Commit and push your changes
4. Submit a pull request

Please make sure your commit messages are meaningful.