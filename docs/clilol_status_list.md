---
title: "clilol status list"
---
## clilol status list

list statuses

### Synopsis

Lists statuses for a single user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

The number of statuses returned can be specified with the --limit
flag. If not set, it will return all statuses for the user.

See the statuslog commands to get statuses for all users.

```
clilol status list [flags]
```

### Options

```
  -a, --address string   address whose status(es) to get (default "mcornick")
  -h, --help             help for list
  -l, --limit int        how many status(es) to get (default all)
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol status](clilol_status.md)	 - do things with statuses

