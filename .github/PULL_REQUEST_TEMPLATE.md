## Description

## Problem
A short description of the problem this PR is addressing.

## Solution
A short description of the chosen method to resolve the problem
with an overview of the logic and implementation details when needed.

## Notes
Other notes that you want to share but do not fit into _Problem_ or _Solution_.

### Checklist
- [ ] The title starts either with `feat(area)`, `fix(area)`, or `chore(area)`
  - `feat` should be used if this pull request implements a new feature
  - `fix` should be used if this pull request fixes a bug
  - `chore` should be used for maintenance changes
  - `area` should be describing the relevant section of the project
- [ ] ran `go mod tidy`
- [ ] ran `go mod vendor`
- [ ] ran `docker run --rm -ti -v $PWD:/workspace -w /workspace eunts/minigo:latest --template -o /workspace/README.md /workspace/README.md.template`

