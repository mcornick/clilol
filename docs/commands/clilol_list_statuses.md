---
title: "clilol list statuses"
---
## clilol list statuses

List statuses

### Synopsis

Lists statuses for a single user from status.lol.

The address can be specified with the --address flag. If not set,
it defaults to your own address.

The number of statuses returned can be specified with the --limit
flag. If not set, it will return all statuses for the user. If
set to more statuses than exist, it will return all statuses.

See the statuslog commands to get statuses for all users.

```
clilol list statuses [flags]
```

### Options

```
  -a, --address string   address whose status(es) to get
  -h, --help             help for statuses
  -l, --limit int        how many status(es) to get (default all)
```

### SEE ALSO

* [clilol list](clilol_list.md)	 - List things

