# Contributing to Runphase

Thanks for your interest. Runphase is pre-alpha; expect things to move.

## Before you start

- Open an issue for anything larger than a small fix so we can agree on the approach.
- Read `ROADMAP.md`. Features outside the release-lifecycle scope will likely be declined.

## Tooling

- **Go** for `cmd/` and `internal/` (one module).
- **bun** for `apps/web` (Next.js). Do not use npm, pnpm, or yarn.
- **uv** for any Python tooling. Do not use pip, conda, or poetry.

## Development

Once V0 lands:

```bash
make dev     # Postgres, Temporal, API, worker, operator, web
make test    # unit + envtest
make e2e     # kind end-to-end
make lint
```

## Pull requests

- Keep PRs focused; one concern per PR.
- Include tests. Workflow changes need a replay test; controller changes need envtest coverage.
- Update docs and ADRs when behavior or architecture changes.
- Sign off your commits (DCO): `git commit -s`.

## Definition of done

State transitions defined, invalid input rejected, idempotency and timeouts considered, failures visible to the user, tests for the critical path, docs updated.

## Code of conduct

This project follows the [Code of Conduct](CODE_OF_CONDUCT.md).
