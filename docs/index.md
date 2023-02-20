---
title: Home
---
__clilol__ is a CLI for the various fun services offered by [omg.lol](https://omg.lol/). At present, it supports the statuslog features (aka [status.lol/](https://status.lol/)), with support for more omg.lol services on the way.

Please see the links in the navigation menu to the left to learn about all the different commands provided by clilol.

## Installation

=== "Homebrew"

    I maintain a [Homebrew](https://brew.sh/) tap.

    ```bash
    brew install mcornick/tap/clilol
    ```

=== "Scoop"

    I maintain a [Scoop](https://scoop.sh/) bucket.

    ```powershell
    scoop bucket add mcornick https://github.com/mcornick/scoop-bucket.git
    scoop install clilol
    ```

=== "Binaries and Linux packages"

    I maintain binary releases on GitHub [here](https://github.com/mcornick/clilol/releases). Releases are built for macOS (universal), Linux (i386, amd64, arm64, and armv6) and Windows (i386, amd64). Linux packages are built in RPM, DEB, APK, and Arch Linux pkg.tar.zst formats.

    Binary checksums included on the release pages are signed with my [GPG key](https://github.com/mcornick.gpg).

=== "YUM Repository"

    RPM packages are also available from my Gemfury repository.

    !!! Note

        I do not, and do not intend to, submit clilol to any distribution's official repositories.

    ```
    # /etc/yum.repos.d/mcornick.repo
    [fury]
    name=mcornick yum repo
    baseurl=https://yum.fury.io/mcornick/
    enabled=1
    gpgcheck=0
    ```

=== "APT Repository"

    DEB packages are also available from my Gemfury repository.

    !!! Note

        I do not, and do not intend to, submit clilol to any distribution's official repositories.

    ```
    # /etc/apt/sources.list.d/mcornick.list
    deb [trusted=yes] https://apt.fury.io/mcornick/ /
    ```

=== "Arch User Repository"

    I maintain an [AUR](https://wiki.archlinux.org/title/Arch_User_Repository) for clilol.

    !!! Note

        I do not, and do not intend to, submit clilol to Arch Linux's official AUR.

    ```
    git clone https://github.com/mcornick/clilol-aur.git
    cd clilol-aur
    makepkg -i
    ```

=== "From source"

    The usual: `go install github.com/mcornick/clilol@latest`

    While I do not build or test for platforms other than the ones listed under the Binaries tab, clilol _should_ still work on any platform supported by Go, and if you find that it does not, feel free to file a GitHub issue and I'll take a look.

## Configuration File

clilol expects a configuration file to specify your username and API key.

!!! Note

    You can find your API key on [your omg.lol account page](https://home.omg.lol/account).

 The configuration file should be named either `config.yaml`, `config.toml` or `config.json` depending on which format you prefer, and should located in one of these directories:

- `$HOME/Library/Application Support/clilol` (macOS)
- `$XDG_CONFIG_HOME/clilol` (Unix)
- `/etc/clilol` (macOS or Unix)
- `%AppData%\clilol` (Windows)

The files should look like this, substituting your own username and API key:

=== "config.yaml"

    ```yaml
    ---
    username: tomservo
    apikey: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
    ```

=== "config.toml"

    ```toml
    username = "tomservo"
    apikey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    ```

=== "config.json"

    ```json
    {
      "username": "tomservo",
      "apikey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    }
    ```

## Environment Variables

Configuration is also possible using environment variables. For example, these environment variables would replicate the default configuration of clilol:

```sh
export CLILOL_USERNAME="tomservo"
export CLILOL_APIKEY="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
```

Environment variables take precedence over any configuration file.

## Contributing to clilol

If you think you have a problem, improvement, or other contribution towards the betterment of clilol, please file an issue or, where appropriate, a pull request.

Keep in mind that I'm not paid to write Go code, so I'm doing this in my spare time, which means it might take me a while to respond.

When filing a pull request, please explain what you're changing and why. Please use standard Go formatting (`go fmt` is your friend.) Please limit your changes to the specific thing you're fixing and isolate your changes in a topic branch that I can merge without pulling in other stuff.

clilol uses [Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md) style. Please follow this convention. Scopes are not required in commit messages.

clilol uses the MIT license. Please indicate your acceptance of the MIT license by using [git commit --signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

clilol is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!

## Etcetera

clilol releases are announced on a [Mastodon account](https://social.lol/@mcornick) which you are welcome to follow.

To verify signatures on commits to clilol, you might need [my GPG public key](https://github.com/mcornick.gpg).