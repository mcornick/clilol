---
title: "clilol update set"
---
## clilol update set

set Now page content

### Synopsis

Sets your Now page content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The Now page will be created as unlisted by default. To create a listed
Now page, use the --listed flag.

```
clilol update set [flags]
```

### Options

```
  -f, --filename string   file to read Now page from (default stdin)
  -h, --help              help for set
  -l, --listed            create Now page as listed (default false)
```

### SEE ALSO

* [clilol update](clilol_update.md)	 - Update things

