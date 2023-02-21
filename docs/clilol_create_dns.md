---
title: "clilol create dns"
---
## clilol create dns

create a DNS record

### Synopsis

Creates a DNS record.

Specify the type of DNS record with the --type flag,
the name of the record with the --name flag,
and the data with the --data flag.

```
clilol create dns [flags]
```

### Options

```
  -d, --data string    Data to store in the DNS record
  -h, --help           help for dns
  -n, --name string    Name of the DNS record to create
  -p, --priority int   Priority of the DNS record
  -T, --ttl int        Time to live of the DNS record (default 3600)
  -t, --type string    Type of DNS record to create
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol create](clilol_create.md)	 - create things

