---
title: "clilol paste create"
---
## clilol paste create

create or update a paste

### Synopsis

Create or update a paste in your pastebin.

Specify a title with the --title flag. If the title is already in use,
that paste will be updated. If the title is not in use, a new paste will
be created.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

```
clilol paste create [flags]
```

### Options

```
  -f, --filename string   file to read paste from (default stdin)
  -h, --help              help for create
  -l, --listed            create paste as listed (default false)
  -t, --title string      title of the paste to create
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol paste](clilol_paste.md)	 - do things with your pastebin

