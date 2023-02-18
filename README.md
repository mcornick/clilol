# clilol

This is a project to create a CLI for the various fun services offered by [omg.lol](https://omg.lol/).

Right now, it supports one command that posts statuses to the statuslog (aka status.lol.) More will be coming, along with better documentation. For now, look on my works, ye mighty, and despair:
## clilol status post

Post a status

### Synopsis

Posts a status to status.lol. Quote the status if it contains spaces.

```
clilol status post [status text] [flags]
```

### Options

```
  -e, --emoji string          Emoji to add to status
  -u, --external-url string   External URL to add to status
  -h, --help                  help for post
```

### Options inherited from parent commands

```
  -k, --apikey string     API key
  -s, --silent            be silent
  -U, --username string   Username
```

## Notes

clilol is available as open source under the terms of the MIT License.