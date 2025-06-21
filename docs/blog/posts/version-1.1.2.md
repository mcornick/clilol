---
date:
  created: 2025-06-21
---
# clilol version 1.1.2

This version fixes a bug seen on macOS where configuration would not be found if placed in a directory other than the default, even if `$XDG_CONFIG_HOME` was set (by default, on macOS, it is not.) Thanks to [prosumer](https://prosumer.omg.lol) for the bug report!
