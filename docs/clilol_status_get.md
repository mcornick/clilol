---
title: "clilol status get"
---
## clilol status get

get status

### Synopsis

Gets status(es) for a single user from status.lol.

The username can be specified with the --username flag. If not set,
it defaults to your own username.

The number of statuses returned can be specified with the --limit
flag. If not set, it will return all statuses for the user. (This
does not currently work if you also use the --json flag. This is
a bug, and I'll fix it eventually.)

See the statuslog commands to get statuses for all users.

```
clilol status get [flags]
```

### Options

```
  -h, --help              help for get
  -l, --limit int         how many status(es) to get (default all; ignored if --json is set)
  -u, --username string   username whose status(es) to get (default "mcornick")
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol status](clilol_status.md)	 - do things with statuses

