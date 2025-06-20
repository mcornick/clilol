---
date:
  created: 2025-06-20
---
# clilol version 1.1.1

I meant to tag this as 1.1.0-pre1, but fatfingered my `git tag` command, and
then ran into an issue with the resulting 1.1.0 release, so YOLO... it's
version 1.1.1! This version makes the jump to 1.1 with a few changes that are
not breaking, but still significant enough to justify the jump.

## Significant Homebrew Change

Starting with this version, the Homebrew formula for clilol has been replaced
with a Homebrew cask. This is more Homebrew-ly correct, and it's what
GoReleaser now expects. If you had the formula installed previously, you'll
need to `brew rm mcornick/tap/clilol` before you install the new version with
`brew install --cask mcornick/tap/clilol`.

## macOS Binary Signed And Notarized

As part of the move to Homebrew Cask, the `clilol` macOS binary needs to be
signed by me and notarized by Apple. It is now! While the Homebrew formula
should've taken care of the scary `THIS APP IS DAMAGED!!!1!` warnings, they are
now definitively in the past if you saw them previously.

## Shimmer And Shine

I'm now using Charm's [fang](https://github.com/charmbracelet/fang) library to
give the UI a little pizazz. This is subject to change as Charm works more on
fang.

## One Man Page To Rule Them All

Starting with this version, clilol has one `man` page briefly summarizing
commands, instead of a whole slew of pages for each subcommand. This makes more
sense to me, but if you want something more exhaustive, [the online
documentation](https://clilol.readthedocs.io/latest/commands/clilol/) still has
a page for each subcommand.

## Miscellaneous

This version fixes a few bugs, works around some changes in the omg.lol API,
and otherwise has changes that are meaningful only to clilol's developer
(that's me, Mark.)

## In Closing

Thanks to everyone who's still using this thing I wrote to learn some Go a few
years back. Thanks also to Adam at [omg.lol](https://omg.lol) for support,
running an ethical service, and being a generally stand-up dude.
