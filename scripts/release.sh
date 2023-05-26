#!/bin/sh -e
goreleaser release --clean
cosign sign --yes --key "$HOME/Documents/cosign.key" "$(crane digest git.mcornick.dev/mcornick/clilol:latest --full-ref)"
