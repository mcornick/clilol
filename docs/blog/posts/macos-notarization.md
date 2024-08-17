---
date:
  created: 2024-08-17
---
# macOS notarization

I have not, to this point, signed or notarized the macOS builds of clilol. It seemed unnecessary, as well as hard to do within my GitHub Actions workflow.

Two things are changing that are making me reconsider this: [GoReleaser](https://goreleaser.com), my build and release tool of choice, now supports signing and notarization; and macOS Sequoia may make it harder to run stuff that isn't signed and notarized (according to various reports.)

So I'm considering whether future versions of clilol should be signed and notarized. I likely won't make a decision on this until Sequoia is released and I get a chance to see how onerous any new requirements are. For now, if you install clilol via Homebrew on current release versions (i.e. not Sequoia betas) of macOS, everything should be fine. Stay tuned.
