#!/bin/sh -e
rm -rf manpages
mkdir manpages
cd manpages
go run .. man
