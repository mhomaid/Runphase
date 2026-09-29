# Local Postgres and Temporal

This stack is local infrastructure. It does not start the Runphase API, worker, or web app.

PostgreSQL is one server with three databases:

| Database | Owner | Contents |
|---|---|---|
| `runphase` | future Runphase product metadata | empty in this task |
| `temporal` | Temporal | workflow execution history |
| `temporal_visibility` | Temporal | Temporal's own visibility index |

Do not create Runphase tables inside `temporal` or `temporal_visibility`.

## Commands

From the repository root:

```bash
make dev          # start Postgres, Temporal, and the Temporal UI
make dev-status   # show container state
make dev-logs     # show recent logs
make dev-down     # stop and remove containers; keep the Postgres volume
```

`make dev` is safe to run again. It does not delete data.

## Ports

| Service | URL |
|---|---|
| PostgreSQL | `localhost:5432` |
| Temporal frontend | `localhost:7233` |
| Temporal UI | `http://localhost:8080` |

The Temporal UI is a local inspection tool. It is not the Runphase product UI.

Local login, used only on this machine:

- user: `runphase`
- password: `runphase`
- product database: `runphase`
- Temporal namespace: `default`

Override a default by setting the variable before `make dev`: `POSTGRES_PORT`, `TEMPORAL_PORT`, `TEMPORAL_UI_PORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `TEMPORAL_NAMESPACE`.

## Delete local data

This removes the named volume and all local databases:

```bash
docker compose --project-name runphase -f deploy/compose/docker-compose.yml down -v
```
