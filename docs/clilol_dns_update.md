---
title: "clilol dns update"
---
## clilol dns update

update a DNS record

### Synopsis

Updates a DNS record.

Specify the ID of the DNS record with the --id flag,
the type of DNS record with the --type flag,
the name of the record with the --name flag,
and the data with the --data flag.

```
clilol dns update [flags]
```

### Options

```
  -d, --data string    Updated data
  -h, --help           help for update
  -i, --id string      ID of DNS record to update
  -n, --name string    Updated record name
  -p, --priority int   Updated priority
  -T, --ttl int        Updated TTL (default 3600)
  -t, --type string    Updated DNS type
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol dns](clilol_dns.md)	 - do things with DNS

