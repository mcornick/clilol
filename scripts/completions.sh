#!/bin/sh -e
rm -rf completions
mkdir completions
for sh in bash zsh fish; do
  go run . completion "$sh" >"completions/clilol.$sh"
done
