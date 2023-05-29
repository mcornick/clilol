---
title: Developer Notes
---

## Running tests

Until such time as there's a non-production instance to connect to, or
there are mock services available for testing, you'll need an address on
the production omg.lol server to run the tests. I use a separate address
for this, and recommend you do the same.

You'll need to set some environment variables to reflect your testing
address:

```
export CLILOL_ADDRESS="tomservo-testing"
export CLILOL_APIKEY="0123456789abcdef0123456789abcdef"
export CLILOL_EMAIL="tomservo@gizmonics.invalid"
export CLILOL_NAME="Tom Servo's test address"
```

With those set, the tests should run successfully, unless there is some
issue connecting to omg.lol (which does happen from time to time,
because omg.lol is on the same internet as jerks running botnets.)

## Contributing to clilol

If you think you have a problem, improvement, or other contribution
towards the betterment of clilol, please file a ticket or, where
appropriate, send a patch.

Tickets can be filed at https://todo.sr.ht/~mcornick/clilol, and patches
can be sent to ~mcornick/clilol@lists.sr.ht. If you need help with the
email-based Git workflow, check out https://git-send-email.io/ for a
tutorial.

Keep in mind that I'm not paid to write Go code, so I'm doing this in my
spare time, which means it might take me a while to respond.

When sending a patch, please explain what you're changing and why.
Please use gofmt to format your Go code. Please limit your changes to
the specific thing you're fixing and isolate your changes in a topic
branch that I can merge without pulling in other stuff.

Please make sure the tests continue to pass. If you're adding new code,
please add new passing tests, too.

clilol uses [Conventional
Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md)
style. Please follow this convention. Scopes are not required in commit
messages.

clilol uses the [MPL-2.0
license](https://www.mozilla.org/en-US/MPL/2.0/). Please indicate your
acceptance of the MPL-2.0 license by using [git commit
--signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

clilol is intended to be a safe, welcoming space for collaboration, and
contributors are expected to adhere to the [Contributor
Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!
