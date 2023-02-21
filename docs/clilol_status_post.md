---
title: "clilol status post"
---
## clilol status post

post a status

### Synopsis

Posts a status to status.lol.

Specify the status text with the --text flag.
Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used.

You can specify an external URL with the --external-url flag. This
will be shown as a "Respond" link on the statuslog. If not set, no
external URL will be used.

```
clilol status post [flags]
```

### Options

```
  -e, --emoji string          Emoji to add to status (default sparkles)
  -a, --external-url string   External URL to add to status
  -h, --help                  help for post
  -t, --text string           Status text
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol status](clilol_status.md)	 - do things with statuses

