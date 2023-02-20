---
title: "clilol completion zsh"
---
## clilol completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(clilol completion zsh); compdef _clilol clilol

To load completions for every new session, execute once:

#### Linux:

	clilol completion zsh > "${fpath[1]}/_clilol"

#### macOS:

	clilol completion zsh > $(brew --prefix)/share/zsh/site-functions/_clilol

You will need to start a new shell for this setup to take effect.


```
clilol completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -j, --json     output json
  -s, --silent   be silent
```

### SEE ALSO

* [clilol completion](clilol_completion.md)	 - Generate the autocompletion script for the specified shell

