---
title: "clilol create dns"
---
## clilol create dns

Create a DNS record

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
  -d, --data string    data to store in the DNS record
  -h, --help           help for dns
  -n, --name string    name of the DNS record to create
  -p, --priority int   priority of the DNS record
  -T, --ttl int        time to live of the DNS record (default 3600)
  -t, --type string    type of DNS record to create
```

### SEE ALSO

* [clilol create](clilol_create.md)	 - Create things

