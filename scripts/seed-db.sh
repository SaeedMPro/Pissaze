#!/usr/bin/env bash
# Seed Pissaze DB: create pissaze_system (schema + Persian test data) and grant app user access.
# Run from project root. Requires: docker compose up -d postgres

set -e
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
STORAGE="$PROJECT_ROOT/back/internal/storage"

echo "Creating database and schema (createdDB.sql)..."
docker compose -f "$PROJECT_ROOT/docker-compose.yml" exec -T postgres \
  psql -U postgres -d postgres < "$STORAGE/createdDB.sql"

echo "Granting pissaze user access to pissaze_system..."
docker compose -f "$PROJECT_ROOT/docker-compose.yml" exec -T postgres \
  psql -U postgres -d pissaze_system -c "
  GRANT CONNECT ON DATABASE pissaze_system TO pissaze;
  GRANT USAGE ON SCHEMA public TO pissaze;
  GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO pissaze;
  GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO pissaze;
  GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO pissaze;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO pissaze;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO pissaze;
"

echo "Seeding Persian test data (persian_dataset.sql)..."
docker compose -f "$PROJECT_ROOT/docker-compose.yml" exec -T postgres \
  psql -U postgres -d pissaze_system < "$STORAGE/persian_dataset.sql"

echo "Done. Backend uses DB_NAME=pissaze_system; restart back if it was already up."