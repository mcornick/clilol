---
title: "clilol update dns"
---
## clilol update dns

Update a DNS record

### Synopsis

Updates a DNS record.

Specify the ID of the DNS record with the --id flag,
the type of DNS record with the --type flag,
the name of the record with the --name flag,
and the data with the --data flag.

```
clilol update dns [flags]
```

### Options

```
  -d, --data string    Updated data
  -h, --help           help for dns
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

* [clilol update](clilol_update.md)	 - Update things

