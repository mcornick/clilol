---
title: Home
---
# clilol, a CLI for omg.lol

__clilol__ is a CLI for the various fun services offered by [omg.lol](https://omg.lol/). All the services supported by the [omg.lol API](https://api.omg.lol) are supported.

![Screenshot](clilol.gif "Screenshot")

Please see the Commands tab in the navigation menu to learn about all the different commands provided by clilol.

## Installation

You can install clilol in any of these ways. (These are the only supported builds of clilol. I don't submit them to "official" repositories, at least not yet. If you find a problem with anyone else's build, please try my builds instead.)

=== "Homebrew"

    ```bash
    brew tap mcornick/tap https://github.com/mcornick/homebrew-tap.git
    brew install mcornick/tap/clilol
    ```

=== "Scoop"

    ```powershell
    scoop bucket add mcornick https://github.com/mcornick/scoop-bucket.git
    scoop install mcornick/clilol
    ```

=== "YUM Repository"

    ```
    # /etc/yum.repos.d/mcornick.repo
    [mcornick]
    name=mcornick yum repo
    baseurl=https://yum.fury.io/mcornick/
    enabled=1
    gpgcheck=0
    ```

=== "APT Repository"

    ```
    # /etc/apt/sources.list.d/mcornick.list
    deb [trusted=yes] https://apt.fury.io/mcornick/ /
    ```

=== "Arch User Repository"

    ```
    git clone https://github.com/mcornick/clilol-aur.git
    cd clilol-aur
    makepkg -i
    ```

=== "Nix"

    ```
    # flake.nix
    inputs.mcornick = {
      url = "github:mcornick/nixpkgs";
      inputs.nixpkgs.follows = "nixpkgs";
    };
    ```

    Then add  `mcornick.packages.x86_64-linux.clilol` to `environment.systemPackages` in your NixOS configuration, or to `home.packages` in your Home Manager configuration. If you're not on a `x86_64-linux` system, use your platform type instead.

    !!! Note

    The flake _theoretically_ supports nix-darwin, but this is untested. If you try it and it works, let me know!

=== "Container Images"

    ```bash
    docker run --rm ghcr.io/mcornick/clilol
    ```

    Container manifests are signed with [Cosign](https://docs.sigstore.dev/cosign/overview/). Ephemeral keys from GitHub are used, so you'll need to specify a certificate identity that matches the tag you're trying to verify.

    ```bash
    cosign verify --certificate-identity=https://github.com/mcornick/clilol/.github/workflows/goreleaser.yaml@refs/tags/vX.Y.Z --certificate-oidc-issuer=https://token.actions.githubusercontent.com ghcr.io/mcornick/clilol:vX.Y.Z
    ```

=== "Binaries and Linux packages"

    I maintain binary releases on GitHub [here](https://github.com/mcornick/clilol/releases). Releases are built for macOS (universal), Linux (i386, amd64, arm64, and armv6) and Windows (i386, amd64). Linux packages are built in RPM, DEB, APK, and Arch Linux pkg.tar.zst formats.

    Binary checksums included on the release pages are signed with my [PGP key](https://github.com/mcornick.gpg).

    !!! Note

        macOS will likely complain that the `clilol` binary is from an
        unidentified developer. To avoid this, install clilol with
        Homebrew.

=== "From source"

    The usual: `go install github.com/mcornick/clilol@latest`

    While I do not build or test for platforms other than the ones listed above, clilol _should_ still build and run on any platform supported by Go, and if you find that it does not, feel free to file an issue, and I'll take a look.

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
A [JSON Schema](https://raw.githubusercontent.com/mcornick/clilol/main/docs/config.schema.json) for the configuration file is available, for editors that support it.

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

Environment variables are the easiest way to pass configuration when running the container images:

```bash
docker run --rm -ti --env CLILOL_ADDRESS=tomservo --env CLILOL_APIKEY=0123456789abcdef0123456789abcdef --env CLILOL_EMAIL=tomservo@gizmonics.invalid ghcr.io/mcornick/clilol ...
# or put the configuration in a dotenv file:
docker run --rm -ti --env-file .env ghcr.io/mcornick/clilol ...
```

Environment variables take precedence over any configuration file.

## Reading apikey from a command

Rather than hardcoding your API key in the configuration file or environment, you can specify a command which, when run, will return the API key on standard output, such as:

=== "config.yaml"

    ```yaml
    ---
    address: tomservo
    email: tomservo@gizmonics.invalid
    apikeycmd: gopass -o omg.lol/tomservo
    ```

=== "config.toml"

    ```toml
    address = "tomservo"
    email = "tomservo@gizmonics.invalid"
    apikeycmd = "gopass -o omg.lol/tomservo"
    ```

=== "config.json"

    ```json
    {
      "address": "tomservo",
      "email": "tomservo@gizmonics.invalid",
      "apikeycmd": "gopass -o omg.lol/tomservo"
    }
    ```

=== "environment"

    ```sh
    export CLILOL_ADDRESS="tomservo"
    export CLILOL_EMAIL="tomservo@gizmonics.invalid"
    export CLILOL_APIKEYCMD="gopass -o omg.lol/tomservo"
    ```

In this example, clilol would use the output of `gopass -o omg.lol/tomservo` as the API key. If the command fails, clilol will print an error stating that the API key is missing.

If apikeycmd is specified, it takes precedence over apikey if that is also specified.

## Etc.

SBOMs and SLSA provenance are generated for each release. You can use these with tools like [grype](https://github.com/anchore/grype) and [slsa-verifier](https://github.com/slsa-framework/slsa-verifier).

```bash
grype clilol_X.Y.Z_darwin_all.tar.gz.sbom
grype ghcr.io/mcornick/clilol
```

```bash
slsa-verifier verify-artifact clilol_X.Y.Z_darwin_all.tar.gz --provenance-path multiple.intoto.jsonl --source-uri github.com/mcornick/clilol --source-tag vX.Y.Z
slsa-verifier verify-image ghcr.io/mcornick/clilol:vX.Y.Z@sha256:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx --source-uri github.com/mcornick/clilol --source-tag vX.Y.Z
```

clilol releases are announced on [my Mastodon account](https://social.lol/@mcornick) and [my Bluesky account](https://bsky.app/profile/mcornick.lol), both of which you are welcome to follow.

To verify signatures on git commits to clilol, you might need [my SSH public key](https://github.com/mcornick.keys) or [my PGP public key](https://github.com/mcornick.gpg).

Thanks to the following people for helping to improve clilol:

- [Andy Piper](https://github.com/andypiper)
- [Elliott Street](https://github.com/ejstreet)
- [John Lynch](https://github.com/jtlynchjr)

clilol is a labor of love. I do not expect, do not request, and will not accept any payment for it. If you like clilol, please show your support by subscribing (or continuing to subscribe) to omg.lol, a service that does not sell you or your data as a product, and thus relies on paid subscribers to keep the lights on.

clilol is made by [Mark Cornick](https://mcornick.lol) who is solely responsible for it. clilol is not a product of Neatnik/omg.lol; please don't bug them for support. Thanks!
