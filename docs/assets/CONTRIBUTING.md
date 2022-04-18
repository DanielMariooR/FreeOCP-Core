# Contributing

## Contribute to freqtrade

Feel like our bot is missing a feature? We welcome your pull requests! 

Issues labeled good first issue can be good first contributions, and will help get you familiar with the codebase.

Few pointers for contributions:

- Create your PR against the `develop` branch, not `main`.
- New features need to contain unit tests and should be documented with the introduction PR.
- PR's can be declared as `[WIP]` - which signify Work in Progress Pull Requests (which are not finished).

If you are unsure, discuss the feature in an issue before a Pull Request.

## Getting started

Best start by reading the [documentation](https://www.freqtrade.io/) to get a feel for what is possible with the bot, or head straight to the [Developer-documentation](https://www.freqtrade.io/en/latest/developer/) (WIP) which should help you getting started.

## Before sending the PR

### 1. Run unit tests

All unit tests must pass. If a unit test is broken, change your code to 
make it pass. It means you have introduced a regression.

#### Test the whole project

```bash
go test -v -cover -coverprofile=cover.out `go list ./...`
```
## Committer Guide

### Pull Requests

How to prioritize pull requests, from most to least important:

1. Fixes for broken tests. Broken means broken on any supported platform or Python version.
1. Extra tests to cover corner cases.
1. Minor edits to docs.
1. Bug fixes.
1. Major edits to docs.
1. Features.

Ensure that each pull request meets all requirements in the Contributing document.

### Issues

If an issue is a bug that needs an urgent fix, mark it for the next patch release.
Then either fix it or mark as please-help.

For other issues: encourage friendly discussion, moderate debate, offer your thoughts.

### Your own code changes

All code changes, regardless of who does them, need to be reviewed and merged by someone else.
This rule applies to all the core committers.

Exceptions:

- Minor corrections and fixes to pull requests submitted by others.
- While making a formal release, the release manager can make necessary, appropriate changes.
- Small documentation changes that reinforce existing subject matter. Most commonly being, but not limited to spelling and grammar corrections.
