---
title: Home
---
__clilol__ is a CLI for the various fun services offered by [omg.lol](https://omg.lol/). All of the services supported by the [omg.lol API](https://api.omg.lol) are supported.

![Screenshot](clilol.gif "Screenshot")

Please see the links in the navigation menu to the left to learn about all the different commands provided by clilol.

!!! Note

    clilol is still under heavy development, and things may still change before a 1.0.0 release. _Caveat lector_.

## Installation

You can install clilol in any of these ways. (These are the only supported builds of omglol. I don't submit them to "official" repositories, at least not yet. If you find a problem with anyone else's build, please try my builds instead.)

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

=== "Container Images"

    I maintain container images on [GitHub](https://github.com/mcornick/clilol/pkgs/container/clilol) and [Docker Hub](https://hub.docker.com/repository/docker/mcornick/clilol).

    ```bash
    docker run --rm ghcr.io/mcornick/clilol
    # or, for Docker Hub
    docker run --rm mcornick/clilol
    ```

    Container manifests are signed with [Cosign](https://docs.sigstore.dev/cosign/overview/). The signatures are created with Cosign's "keyless" mode, which requires setting `COSIGN_EXPERIMENTAL=1` when using Cosign versions prior to 2.0.0:

    ```bash
    env COSIGN_EXPERIMENTAL=1 cosign verify ghcr.io/mcornick/clilol
    # or, for Docker Hub
    env COSIGN_EXPERIMENTAL=1 cosign verify mcornick/clilol
    ```

=== "Binaries and Linux packages"

    I maintain binary releases on GitHub [here](https://github.com/mcornick/clilol/releases). Releases are built for macOS (universal), Linux (i386, amd64, arm64, and armv6) and Windows (i386, amd64). Linux packages are built in RPM, DEB, APK, and Arch Linux pkg.tar.zst formats.

    Binary checksums included on the release pages are signed with my [GPG key](https://github.com/mcornick.gpg).

=== "YUM Repository"

    RPM packages are also available from my Gemfury repository.

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

    ```
    # /etc/apt/sources.list.d/mcornick.list
    deb [trusted=yes] https://apt.fury.io/mcornick/ /
    ```

=== "Arch User Repository"

    I maintain an [AUR](https://wiki.archlinux.org/title/Arch_User_Repository) for clilol.

    ```
    git clone https://github.com/mcornick/clilol-aur.git
    cd clilol-aur
    makepkg -i
    ```

=== "From source"

    The usual: `go install github.com/mcornick/clilol@latest`

    While I do not build or test for platforms other than the ones listed under the Binaries tab, clilol _should_ still work on any platform supported by Go, and if you find that it does not, feel free to file a GitHub issue and I'll take a look.

## Configuration File

clilol expects a configuration file to specify your address, login email, and API key. You can find your API key on [your omg.lol account page](https://home.omg.lol/account).

The configuration file should be named either `config.yaml`, `config.toml` or `config.json` depending on which format you prefer, and should be located in one of these directories:

- `$HOME/Library/Application Support/clilol` (macOS)
- `$XDG_CONFIG_HOME/clilol` (Unix)
- `/etc/clilol` (macOS or Unix)
- `%AppData%\clilol` (Windows)

The file should look like one of these, substituting your own details:

=== "config.yaml"

    ```yaml
    ---
    address: tomservo
    email: tomservo@gizmonics.example.com
    apikey: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
    ```

=== "config.toml"

    ```toml
    address = "tomservo"
    email = "tomservo@gizmonics.example.com"
    apikey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    ```

=== "config.json"

    ```json
    {
      "address": "tomservo",
      "email": "tomservo@gizmonics.example.com",
      "apikey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    }
    ```

## Environment Variables

Configuration is also possible using environment variables:

```sh
export CLILOL_USERNAME="tomservo"
export CLILOL_EMAIL="tomservo@gizmonics.example.com"
export CLILOL_APIKEY="XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
```

Environment variables take precedence over any configuration file.

Your email address is only needed for the `clilol account` commands. It is not used by clilol for anything else, such as spamming you. If you don't use the `clilol account` commands, feel free to not specify it.

## Contributing to clilol

If you think you have a problem, improvement, or other contribution towards the betterment of clilol, please file an issue or, where appropriate, a pull request.

Keep in mind that I'm not paid to write Go code, so I'm doing this in my spare time, which means it might take me a while to respond.

When filing a pull request, please explain what you're changing and why. Please use [gofumpt](https://github.com/mvdan/gofumpt) to format your Go code. Please limit your changes to the specific thing you're fixing and isolate your changes in a topic branch that I can merge without pulling in other stuff.

clilol uses [Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md) style. Please follow this convention. Scopes are not required in commit messages.

clilol uses the [MPL-2.0 license](https://www.mozilla.org/en-US/MPL/2.0/). Please indicate your acceptance of the MPL-2.0 license by using [git commit --signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

clilol is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!

## Etcetera

clilol releases are announced on [my social.lol account](https://social.lol/@mcornick) which you are welcome to follow.

To verify signatures on commits to clilol, you might need [my GPG public key](https://github.com/mcornick.gpg).
