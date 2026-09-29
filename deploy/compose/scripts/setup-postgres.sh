#!/bin/sh
# Create Temporal's own databases. Safe to run again after the data volume already exists.
set -eu

: "${POSTGRES_SEEDS:?POSTGRES_SEEDS is required}"
: "${POSTGRES_USER:?POSTGRES_USER is required}"
: "${POSTGRES_PWD:?POSTGRES_PWD is required}"

export SQL_PASSWORD="${POSTGRES_PWD}"
DB_PORT="${DB_PORT:-5432}"
PLUGIN=postgres12
SCHEMA_ROOT=/etc/temporal/schema/postgresql/v12

echo "Waiting for PostgreSQL at ${POSTGRES_SEEDS}:${DB_PORT}..."
attempt=1
while ! nc -z -w 5 "${POSTGRES_SEEDS}" "${DB_PORT}"; do
  if [ "${attempt}" -ge 60 ]; then
    echo "PostgreSQL did not accept connections"
    exit 1
  fi
  attempt=$((attempt + 1))
  sleep 1
done

setup_database() {
  db="$1"
  schema_dir="$2"

  if temporal-sql-tool \
    --plugin "${PLUGIN}" \
    --ep "${POSTGRES_SEEDS}" \
    -u "${POSTGRES_USER}" \
    -p "${DB_PORT}" \
    --db "${db}" \
    create; then
    echo "created database ${db}"
  else
    echo "database ${db} was not created; continuing because it may already exist"
  fi

  if temporal-sql-tool \
    --plugin "${PLUGIN}" \
    --ep "${POSTGRES_SEEDS}" \
    -u "${POSTGRES_USER}" \
    -p "${DB_PORT}" \
    --db "${db}" \
    setup-schema -v 0.0; then
    echo "initialized schema for ${db}"
  else
    echo "schema for ${db} was not initialized; continuing because it may already exist"
  fi

  temporal-sql-tool \
    --plugin "${PLUGIN}" \
    --ep "${POSTGRES_SEEDS}" \
    -u "${POSTGRES_USER}" \
    -p "${DB_PORT}" \
    --db "${db}" \
    update-schema -d "${schema_dir}"
  echo "schema for ${db} is current"
}

setup_database temporal "${SCHEMA_ROOT}/temporal/versioned"
setup_database temporal_visibility "${SCHEMA_ROOT}/visibility/versioned"

echo "Temporal PostgreSQL schema setup complete"
