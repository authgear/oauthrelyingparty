---
name: dep-audit
description: Audit and fix Go dependency vulnerabilities. Runs govulncheck and upgrades affected modules.
argument-hint: "--fix"
---

Audit and fix dependency vulnerabilities in this project. Follow the steps below in order.

## Step 1: Go Vulnerability Check

Run `make govulncheck` in the repository root.

Parse the output:
- If there are **no vulnerabilities**, note it and stop.
- If there are **vulnerabilities**, for each affected module:
  1. Run `go list -m -u <module>` to find the latest available version.
  2. Compare the current version with the latest:
     - If the major version changes (e.g. `v1.x.x` → `v2.x.x`), generate a **Breaking Change Report** (see below), defer the upgrade, and continue with any patchable work.
     - If only minor/patch version changes, proceed automatically.
  3. After all patchable fixes are identified, run `go get <module>@latest` then `go mod tidy` in the repository root.

Important:
- `go get <module>@latest` is for Go module dependencies only. Do **not** use it to change the Go toolchain version.
- If the requested fix is a Go version bump, update `.tool-versions` and `go.mod` directly, then run `go mod tidy`.

**Breaking Change Report (major version bumps)** must include:
- Module name, current version → proposed version
- Link to the module's changelog or migration guide if available
- Known incompatibilities (import path changes, removed/renamed symbols)
- Ask: "Do you want to apply this major version upgrade? (yes/no)"

## Step 2: Verify and Commit

After updating Go deps:
- Run `go build ./...` to verify the build.
- If the build breaks, report the compiler errors and ask the user how to proceed. Do not commit.
- If vulnerabilities cannot be fixed (no fix available), note them in an **Unfixable Issues Report** and notify the user.
- If fixes were applied and build passes, stage and commit `go.mod` and `go.sum` with commit message: `chore: fix Go dependency vulnerabilities`

## Step 3: Final Summary

Output a summary with three sections:

### Fixed
List every module that was updated (name, old version → new version).

### Breaking Changes Applied
List any breaking changes that were confirmed and applied.

### Unfixable Issues
List any vulnerabilities that could not be resolved, including:
- Module name and version
- CVE/advisory ID
- Why it cannot be fixed
- Recommended action
