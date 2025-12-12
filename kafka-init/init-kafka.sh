#!/usr/bin/env bash
set -e

TOPICS=("uploaded_files" "task_result" "tasks")

echo "Starting Kafka topic initialization..."

for i in {1..30}; do
  echo "Attempt $i to create topics..."

  for topic in "${TOPICS[@]}"; do
    kafka-topics \
      --bootstrap-server "${KAFKA_BROKER}" \
      --create \
      --topic "$topic" \
      --partitions 1 \
      --replication-factor 1 \
      --if-not-exists || true
  done

  echo "Topics created (or already exist)"
  exit 0
  sleep 5
done

echo "Failed to initialize topics after 30 attempts"
exit 1
