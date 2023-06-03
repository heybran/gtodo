<p>
  <img src="https://github.com/heybran/gtodo/blob/main/screenshot.png">
</p>

# Go CLI Todo App
This is a simple command-line todo application written in Go. It allows you to create, read, update, and delete tasks into/from a local JSON file.

## Installation
To install the application, you will need to have Go installed on your system. Once you have installed Go, you can download and install the application using the following command:

```bash
go install github.com/heybran/gtodo@latest
```
This will download the source code and install the application in your `$GOPATH/bin` directory.

## Usage
Be sure to run `gtodo init` to generate an empty JSON file in your home directory to store todo tasks.

To use the application, simply run the `gtodo` command followed by one of the following subcommands:

- `add`: Add a new task.
- `list`: List all tasks.
- `update`: Update the details of an existing task.
- `delete`: Delete a task.

Each subcommand has its own set of options and arguments. Here are some examples of how to use the application:

```bash
# Add a new task
gtodo add -task "Recording tutorials on web components :host selector" -cat "Tutorial"

# List all tasks
gtodo list
gtodo list -done 1 # list tasks that are completed
gtodo list -done 1 -cat "Project" # list tasks that are completed and belong to category "Project"

# Update an existing task
gtodo update -id 1 -task "Hrmm, maybe write a blog post about :host selector instead"
gtodo update -id 1 -cat "Blog"

# Delete a task
gtodo delete -id 1
```

## Uninstall
To uninstall the gtodo CLI app that you installed using ﻿go install, you can run the following command in your terminal:

```bash
go clean -i github.com/heybran/gtodo
```

This will remove the gtodo binary from your `$GOPATH/bin` directory. If you want to remove the source code as well, you can delete the directory `$GOPATH/src/github.com/heybran/gtodo`.
  
## Contributing
If you find a bug or have a feature request, please open an issue on the GitHub repository. Pull requests are also welcome!

## License
This application is licensed under the MIT License. See the ﻿LICENSE file for details.
