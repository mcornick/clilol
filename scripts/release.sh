#!/bin/sh -e
goreleaser release --clean
cosign sign --yes "$(crane digest ghcr.io/mcornick/clilol:latest --full-ref)"
cosign sign --yes "$(crane digest mcornick/clilol:latest --full-ref)"
