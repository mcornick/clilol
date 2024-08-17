---
title: "clilol create status"
---

Create a status

## Synopsis

Posts a status to status.lol.

Quote the text if it contains spaces.

You can specify an emoji with the --emoji flag. This must be an
actual emoji, not a :emoji: style code. If not set, the sparkles
emoji will be used.

If you have enabled cross-posting to Mastodon in your statuslog
settings, you can skip cross-posting to Mastodon by setting the
--skip-mastodon-post flag.

```
clilol create status [text] [flags]
```

## Options

```
  -e, --emoji string         emoji to add to status (default sparkles)
  -h, --help                 help for status
      --skip-mastodon-post   do not cross-post to Mastodon
```

## SEE ALSO

* [clilol create](clilol_create.md)	 - Create things
