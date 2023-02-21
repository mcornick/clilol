---
title: "clilol web set"
---
## clilol web set

set webpage content

### Synopsis

Sets your webpage content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

Set the --publish flag to true publish your webpage. By default, it will not
be published.

```
clilol web set [flags]
```

### Options

```
  -f, --filename string   file to read webpage from (default stdin)
  -h, --help              help for set
  -p, --publish           publish the updated page (default false)
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol web](clilol_web.md)	 - do things with your webpage

