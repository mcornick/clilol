---
title: Developer Notes
---

## Return To GitHub

clilol development started on GitHub in early 2023. In mid-2023, following the 1.0 release, I moved off GitHub, motivated by a desire to distance clilol from some of GitHub's changes in organizational direction, such as scraping public code for Copilot.

After several months away, I moved clilol back to GitHub in February 2024. This was motivated by a few things:

- I want to minimize barriers to collaboration. Hosting on my private Git server, where I do not allow outside registrations, cut collaboration off almost entirely. Mirroring on Sourcehut allowed people to participate if they were willing to learn Sourcehut's contribution flow. For better or worse, lots of people are comfortable working through GitHub, and I don't want to exclude them.

- I am a professional UNIX administrator. Running a server on my personal time to host Git is simple for me to do, but is a (admittedly light) burden and a finger on the scale of work-life balance. I'd like to be out of that business eventually.

- Sourcehut still has my (paid) support, but their January 2024 outage left me scrambling to recover a lot of repositories that I didn't have elsewhere. For whatever flaws GitHub has, I don't see them being down for over a week at a time. Additionally, as mentioned, the email-based workflow that Sourcehut uses for contributions is probably to blame for me having received none.

I hope this clarifies why I've made this decision. I'm keeping the other repositories online because they contain some history that is missing from GitHub, but going forward, GitHub will once again be the canonical source.

## Running tests

Until such time as there's a non-production instance to connect to, or there are mock services available for testing, you'll need an address on the production omg.lol server to run the tests. I use a separate address for this, and recommend you do the same.

You'll need to set some environment variables to reflect your testing address. A [Taskfile](https://taskfile.dev) is provided that can take care of this for you (copy `taskfile.dist.yaml` to `taskfile.yaml` and fill in your details), or you can do it yourself:

```
export CLILOL_ADDRESS="tomservo-testing"
export CLILOL_APIKEY="0123456789abcdef0123456789abcdef"
export CLILOL_EMAIL="tomservo@gizmonics.invalid"
export CLILOL_NAME="Tom Servo's test address"
```

With those set, the tests should run successfully, unless there is some issue connecting to omg.lol (which does happen from time to time, because omg.lol is on the same internet as jerks running botnets.)

## Contributing to clilol

If you think you have a problem, improvement, or other contribution towards the betterment of clilol, please file an issue or, where appropriate, open a pull request.

As of February 2024, I have moved clilol back to GitHub, to reduce barriers to collaboration. Prior versions can be found at [Sourcehut](https://git.sr.ht/~mcornick/clilol).

Keep in mind that I'm not paid to write Go code, so I'm doing this in my spare time, which means it might take me a while to respond.

When opening a pull request, please explain what you're changing and why.  Please use gofmt to format your Go code. Please limit your changes to the specific thing you're fixing and isolate your changes in a topic branch that I can merge without pulling in other stuff.

Please make sure the tests continue to pass. If you're adding new code, please add new passing tests, too.

clilol uses [Conventional Changelog](https://github.com/conventional-changelog/conventional-changelog-angular/blob/master/convention.md) style. Please follow this convention. Scopes are not required in commit messages.

clilol uses the [MPL-2.0 license](https://www.mozilla.org/en-US/MPL/2.0/). Please indicate your acceptance of the MPL-2.0 license by using [git commit --signoff](https://git-scm.com/docs/git-commit#Documentation/git-commit.txt--s).

clilol is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

Thanks for contributing!
