---
title: "clilol status update"
---
## clilol status update

update a status

### Synopsis

Updates a status on status.lol.
Specify the ID of the status to update with the --id flag. The
status can be found as the last element of the status URL.

Quote the status text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used. Note that the omg.lol API does not preserve
the existing emoji if you don't specify one, so if you don't want
to change it, you'll still need to specify it again.

```
clilol status update [status text] [flags]
```

### Options

```
  -e, --emoji string   Emoji to add to status (default sparkles)
  -h, --help           help for update
  -i, --id string      ID of the status to update
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol status](clilol_status.md)	 - do things with statuses

