#!/bin/sh
# Run createdDB.sql (creates pissaze_system + schema) then Persian seed.
# Uses POSTGRES_USER (pissaze) - the image does not create a "postgres" role.
set -e
echo "Running createdDB.sql..."
psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /storage/createdDB.sql
echo "Granting pissaze access to pissaze_system..."
psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d pissaze_system -c "
  GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO pissaze;
  GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO pissaze;
  GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO pissaze;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO pissaze;
  ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO pissaze;
"
echo "Seeding persian_dataset.sql..."
psql -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d pissaze_system -f /storage/persian_dataset.sql
echo "Init done."
