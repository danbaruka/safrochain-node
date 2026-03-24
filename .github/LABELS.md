# Suggested GitHub labels

GitHub does not read this file automatically. Use it as a checklist to create labels in **Settings → Labels**, or run the `gh` commands below after replacing the repo if you fork.

## Priority (optional but recommended)

| Label | Color | Meaning |
|-------|-------|---------|
| `priority:P0` | `#B60205` | Critical — immediate attention |
| `priority:P1` | `#D93F0B` | High — next up after P0 |
| `priority:P2` | `#FBCA04` | Medium — normal backlog |
| `priority:P3` | `#0E8A16` | Low — when time allows |
| `priority:needs-triage` | `#EDEDED` | Not yet classified |

## Type / workflow

| Label | Color | Meaning |
|-------|-------|---------|
| `type:bug` | `#D73A4A` | Confirmed defect |
| `type:feature` | `#A2EEEF` | Feature or enhancement |
| `type:task` | `#C5DEF5` | Chore, refactor, CI, docs |
| `triage` | `#D4C5F9` | Needs maintainer triage |

## Area (optional)

| Label | Color | Meaning |
|-------|-------|---------|
| `area:node` | `#1D76DB` | Daemon / core node |
| `area:cli` | `#0052CC` | Command-line tools |
| `area:build` | `#5319E7` | Build, release, packaging |
| `area:docker` | `#2496ED` | Container images |
| `area:ci` | `#FEF2C0` | GitHub Actions / automation |
| `area:tests` | `#BFDADC` | Test suite / interchaintest |

## Create via GitHub CLI (optional)

```bash
REPO=danbaruka/safrochain-node

gh label create "priority:P0" --repo "$REPO" --color B60205 --description "Critical"
gh label create "priority:P1" --repo "$REPO" --color D93F0B --description "High"
gh label create "priority:P2" --repo "$REPO" --color FBCA04 --description "Medium"
gh label create "priority:P3" --repo "$REPO" --color 0E8A16 --description "Low"
gh label create "priority:needs-triage" --repo "$REPO" --color EDEDED --description "Awaiting triage"

gh label create "type:bug" --repo "$REPO" --color D73A4A --description "Bug"
gh label create "type:feature" --repo "$REPO" --color A2EEEF --description "Feature"
gh label create "type:task" --repo "$REPO" --color C5DEF5 --description "Task / chore"
gh label create "triage" --repo "$REPO" --color D4C5F9 --description "Needs triage"

gh label create "area:node" --repo "$REPO" --color 1D76DB --description "Node / daemon"
gh label create "area:cli" --repo "$REPO" --color 0052CC --description "CLI"
gh label create "area:build" --repo "$REPO" --color 5319E7 --description "Build / release"
gh label create "area:docker" --repo "$REPO" --color 2496ED --description "Docker"
gh label create "area:ci" --repo "$REPO" --color FEF2C0 --description "CI"
gh label create "area:tests" --repo "$REPO" --color BFDADC --description "Tests"
```

**Note:** Issue templates apply GitHub’s default **`bug`** and **`enhancement`** labels. After you create `type:*` labels, you can edit the YAML under `.github/ISSUE_TEMPLATE/` to use those names instead if you prefer a single scheme.
