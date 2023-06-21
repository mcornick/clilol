---
title: Home
---
__clilol__ is a CLI for the various fun services offered by
[omg.lol](https://omg.lol/). All the services supported by the [omg.lol
API](https://api.omg.lol) are supported.

![Screenshot](clilol.gif "Screenshot")

Please see the links in the navigation menu to the left to learn about
all the different commands provided by clilol.

## Installation

You can install clilol in any of these ways. (These are the only
supported builds of clilol. I don't submit them to "official"
repositories, at least not yet. If you find a problem with anyone else's
build, please try my builds instead.)

### Homebrew

I maintain a [Homebrew](https://brew.sh/) tap.

```bash
brew tap mcornick/tap https://git.mcornick.dev/mcornick/homebrew-tap
brew install mcornick/tap/clilol
```

### Scoop

I maintain a [Scoop](https://scoop.sh/) bucket.

```powershell
scoop bucket add mcornick https://git.mcornick.dev/mcornick/scoop-bucket
scoop install clilol
```

### Container Images

I maintain container images on my Git server
[here](https://git.mcornick.dev/mcornick/clilol/packages).

```bash
docker run --rm git.mcornick.dev/mcornick/clilol
```

Container manifests are signed with
[Cosign](https://docs.sigstore.dev/cosign/overview/). The signatures are
created with my [Cosign key](https://mcornick.dev/mcornick.cosign):

```bash
cosign verify --key https://mcornick.dev/mcornick.cosign git.mcornick.dev/mcornick/clilol
```

### Binaries and Linux packages

I maintain binary releases on my Git server
[here](https://git.mcornick.dev/mcornick/clilol/releases). Releases are built
for macOS (universal), Linux (i386, amd64, arm64, and armv6) and Windows
(i386, amd64). Linux packages are built in RPM, DEB, APK, and Arch Linux
pkg.tar.zst formats.

Binary checksums included on the release pages are signed with my [GPG
key](https://git.mcornick.dev/mcornick.gpg).

### YUM Repository

RPM packages are also available from my Gemfury repository.

```
# /etc/yum.repos.d/mcornick.repo
[mcornick]
name=mcornick yum repo
baseurl=https://yum.fury.io/mcornick/
enabled=1
gpgcheck=0
```

### APT Repository

DEB packages are also available from my Gemfury repository.

```
# /etc/apt/sources.list.d/mcornick.list
deb [trusted=yes] https://apt.fury.io/mcornick/ /
```

### Arch User Repository

I maintain an
[AUR](https://wiki.archlinux.org/title/Arch_User_Repository) for clilol.

```
git clone https://git.mcornick.dev/mcornick/clilol-aur
cd clilol-aur
makepkg -i
```

### From source

The usual: `go install git.mcornick.dev/mcornick/clilol@latest`

While I do not build or test for platforms other than the ones listed
under the Binaries tab, clilol _should_ still work on any platform
supported by Go, and if you find that it does not, feel free to file an
issue, and I'll take a look.

## Configuration File

clilol expects a configuration file to specify your address, login
email, and API key. You can find your API key on [your omg.lol account
page](https://home.omg.lol/account).

The configuration file should be named either `config.yaml`,
`config.toml` or `config.json` depending on which format you prefer, and
should be located in one of these directories:

- `$HOME/Library/Application Support/clilol` (macOS)
- `$XDG_CONFIG_HOME/clilol` (Unix)
- `/etc/clilol` (macOS or Unix)
- `%AppData%\clilol` (Windows)

The file should look like one of these, substituting your own details:

=== "config.yaml"

    ```yaml
    ---
    address: tomservo
    email: tomservo@gizmonics.invalid
    apikey: 0123456789abcdef0123456789abcdef
    ```

=== "config.toml"

    ```toml
    address = "tomservo"
    email = "tomservo@gizmonics.invalid"
    apikey = "0123456789abcdef0123456789abcdef"
    ```

=== "config.json"

    ```json
    {
      "address": "tomservo",
      "email": "tomservo@gizmonics.invalid",
      "apikey": "0123456789abcdef0123456789abcdef"
    }
    ```
A [JSON Schema](config.schema.json) for the configuration file is
available, for editors that support it.

!!! Note

    Your email address is only needed to identify your account for the
    `clilol account` commands. It is not used by clilol for anything
    else, such as spamming you.

## Environment Variables

Configuration is also possible using environment variables:

```sh
export CLILOL_ADDRESS="tomservo"
export CLILOL_EMAIL="tomservo@gizmonics.invalid"
export CLILOL_APIKEY="0123456789abcdef0123456789abcdef"
```

Environment variables are the easiest way to pass configuration when
running the container images:

```bash
docker run --rm -ti --env CLILOL_ADDRESS=tomservo --env CLILOL_APIKEY=0123456789abcdef0123456789abcdef --env CLILOL_EMAIL=tomservo@gizmonics.invalid git.mcornick.dev/mcornick/clilol ...
# or put the configuration in a dotenv file:
docker run --rm -ti --env-file .env git.mcornick.dev/mcornick/clilol ...
```

Environment variables take precedence over any configuration file.

## Etc.

clilol releases are announced on [my Mastodon
account](https://social.sdf.org/@mcornick) which you are welcome to follow.

To verify signatures on commits to clilol, you might need [my SSH public
key](https://git.mcornick.dev/mcornick.keys) or [my GPG public
key](https://git.mcornick.dev/mcornick.gpg).

Thanks to the following people for helping to improve clilol:

- [Andy Piper](https://github.com/andypiper)
- [Elliott Street](https://github.com/ejstreet)

clilol is a labor of love. I do not expect, do not request, and will not
accept any payment for it. If you like clilol, please show your support
by subscribing (or continuing to subscribe) to omg.lol, a service that
does not sell you or your data as a product, and thus relies on paid
subscribers to keep the lights on.

clilol is made by [Mark Cornick](https://mcornick.omg.lol) who is solely
responsible for it. clilol is not a product of Neatnik/omg.lol; please
don't bug them for support. Thanks!
