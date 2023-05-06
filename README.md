## How To Install?
### Using [`GO`](https://go.dev/dl) toolchain

```sh
go install github.com/heybran/gtodo@main
```

## Usage

### Init
You need to run `init` subcommand to create an empty JSON file to store your todo items.

```bash
gtodo init
```

### Add
- `-task` flag for the todo content.
- `-cat` flag for the todo category, will default to `Uncategorized` if not provided.

```bash
gtodo add -task="Record tutorial for this cli todo app" -cat="Tutorial"
```

### List

```bash
gtodo list
```

### Delete
- `-id` flag for the id of todo to be deleted.
```bash
gtodo delete -id=1
```

### Update
#### Marked a todo as `done`
```bash
gtodo update -id=1 -done=1
```

#### Update todo content
- `-task` to update content of this todo item.
- `-cat` to update category of this todo item.
```bash
gtodo update -id=1 -task="Hrmm, maybe just write a blog post for this cli todo app" -cat="Blog"
```
