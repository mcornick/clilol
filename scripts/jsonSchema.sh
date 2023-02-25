#!/bin/sh -e
cd docs
go run .. json-schema > config.schema.json
