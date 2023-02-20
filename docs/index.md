---
title: Home
---
This is a project to create a CLI for the various fun services offered by [omg.lol](https://omg.lol/).

Please see the links in the navigation menu to the left to learn about all the different commands provided by clilol.

## You've Got Questions, I've Got Answers

I'd be lying if I said they were frequently asked, but it is what it is.

This section will be rewritten and improved in the future.

### What do I need to run clilol?

1. A computer and operating system supported by [Go](https://go.dev)
2. An [omg.lol](https://omg.lol) account
3. The API key for your omg.lol account, which you can find [here](https://home.omg.lol/account)

### How do I install clilol?

1. Get [Go](https://go.dev)
2. Clone the repo from [GitHub](https://github.com/mcornick/clilol)
3. `go build` to get a binary named `clilol`
4. Put that `clilol` binary wherever you like

Once clilol is a little more mature, I'll provide pre-built binaries, Linux packages, Homebrew, etc.

### Where do I put that API key?

Create a file `config.yaml` that looks like this:

```yaml
username: tomservo
apikey: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

Fill in the X's with the API key you found above, and use your own username, of course.

Then make one of these directories, depending on what operating system you use:

- `$HOME/Library/Application Support/clilol` (macOS)
- `$XDG_CONFIG_HOME/clilol` (Linux/Unix)
- `%AppData%\clilol` (Windows)

Finally, put the `config.yaml` file in that directory you just created.

If you prefer not to keep this information on disk, on macOS and Linux/Unix systems you can set the `CLILOL_USERNAME` and `CLILOL_APIKEY` environment variables. I have no idea if this works on Windows or if there is a Windows compatible alternative.

Now you should be ready to lol.

### How do I specify emoji when posting a status?

Use the `--emoji` option:

```
clilol status post --emoji '🇫🇷' 'Ooh la la, ah oui oui!'
```

You will need to type in the actual emoji, not a :emoji: style code. I think this is a limitation of the omg.lol API, but I'm going to see if it will accept codes.

If you don't specify an emoji, the default is sparkles (✨)

### How do you pronounce clilol?

I pronounce it "see ell eye lol." "See ell eye ell oh ell" works too.

### Is this an official/supported omg.lol product?

No. I am only a happy customer, not otherwise affiliated with Neatnik/omg.lol. So don't bug them if you find a problem. Bug me by opening a GitHub issue.
