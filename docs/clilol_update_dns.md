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
  -d, --data string    updated data
  -h, --help           help for dns
  -i, --id string      ID of DNS record to update
  -n, --name string    updated record name
  -p, --priority int   ipdated priority
  -T, --ttl int        updated TTL (default 3600)
  -t, --type string    updated DNS type
```

### SEE ALSO

* [clilol update](clilol_update.md)	 - Update things

