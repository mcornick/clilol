---
date:
  created: 2025-07-15
---
# clilol version 1.1.3

This version fixes a bug where DNS operations (create, update, and get, specifically) would fail if provided a DNS record type (such as A, CNAME, etc.) that's not all uppercase. Now, clilol will convert any provided record type to uppercase before reaching out to the omg.lol API.

As usual, there are a few dependency updates and other less interesting changes, too.
