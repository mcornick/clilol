---
title: "clilol create paste"
---
## clilol create paste

Create or update a paste

### Synopsis

Create or update a paste in your pastebin.

If the specified title is already in use, that paste will be updated.
If the title is not in use, a new paste will be created.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The paste will be created as unlisted by default. To create a listed
paste, use the --listed flag.

```
clilol create paste [title] [flags]
```

### Options

```
  -f, --filename string   file to read paste from (default stdin)
  -h, --help              help for paste
  -l, --listed            create paste as listed (default false)
```

### SEE ALSO

* [clilol create](clilol_create.md)	 - Create things

