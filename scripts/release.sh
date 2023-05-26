#!/bin/sh -e
goreleaser release --clean
cosign sign --yes --key "$HOME/Documents/cosign.key" "$(crane digest ghcr.io/mcornick/clilol:latest --full-ref)"
cosign sign --yes --key "$HOME/Documents/cosign.key" "$(crane digest mcornick/clilol:latest --full-ref)"
