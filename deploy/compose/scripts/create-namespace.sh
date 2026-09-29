#!/bin/sh
# Ensure one local namespace exists. Safe to run again.
set -eu

NAMESPACE="${DEFAULT_NAMESPACE:-default}"
TEMPORAL_ADDRESS="${TEMPORAL_ADDRESS:-temporal:7233}"
MAX_ATTEMPTS="${TEMPORAL_HEALTH_CHECK_MAX_ATTEMPTS:-30}"
SLEEP_SECONDS="${TEMPORAL_HEALTH_CHECK_SLEEP_SECONDS:-2}"

SERVER_HOST=$(echo "${TEMPORAL_ADDRESS}" | cut -d: -f1)
SERVER_PORT=$(echo "${TEMPORAL_ADDRESS}" | cut -d: -f2)

echo "Waiting for Temporal at ${TEMPORAL_ADDRESS}..."
attempt=1
while ! nc -z -w 10 "${SERVER_HOST}" "${SERVER_PORT}"; do
  if [ "${attempt}" -ge "${MAX_ATTEMPTS}" ]; then
    echo "Temporal port did not open"
    exit 1
  fi
  attempt=$((attempt + 1))
  sleep "${SLEEP_SECONDS}"
done

attempt=1
while ! temporal operator cluster health --address "${TEMPORAL_ADDRESS}"; do
  if [ "${attempt}" -ge "${MAX_ATTEMPTS}" ]; then
    echo "Temporal did not become healthy"
    exit 1
  fi
  attempt=$((attempt + 1))
  sleep "${SLEEP_SECONDS}"
done

if temporal operator namespace describe --namespace "${NAMESPACE}" --address "${TEMPORAL_ADDRESS}" >/dev/null 2>&1; then
  echo "namespace ${NAMESPACE} already exists"
  exit 0
fi

temporal operator namespace create \
  --namespace "${NAMESPACE}" \
  --address "${TEMPORAL_ADDRESS}" \
  --retention 24h
echo "namespace ${NAMESPACE} created"
