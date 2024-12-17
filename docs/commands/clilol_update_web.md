---
title: "clilol update web"
---
## clilol update web

set webpage content

### Synopsis

Sets your webpage content.

If you specify a filename with the --filename flag, the content of the file
will be used. If you do not specify a filename, the content will be read
from stdin.

The webpage will be created as unpublished by default. To create a published
webpage, use the --publish flag.

```
clilol update web [flags]
```

### Options

```
  -f, --filename string   file to read webpage from (default stdin)
  -h, --help              help for web
  -p, --publish           publish the updated page (default false)
```

### SEE ALSO

* [clilol update](clilol_update.md)	 - Update things
* [clilol update web pfp](clilol_update_web_pfp.md)	 - set your profile picture

