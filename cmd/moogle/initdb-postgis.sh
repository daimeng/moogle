#!/bin/sh

set -e

export PGUSER="$POSTGRES_USER"
DB=geocoder

echo "Loading PostGIS extensions into $DB"
psql --dbname="$DB" <<-'EOSQL'
  CREATE EXTENSION IF NOT EXISTS postgis;
  CREATE EXTENSION IF NOT EXISTS postgis_topology;
  CREATE EXTENSION IF NOT EXISTS fuzzystrmatch;
  CREATE EXTENSION IF NOT EXISTS postgis_tiger_geocoder;
  CREATE EXTENSION IF NOT EXISTS address_standardizer;
EOSQL

echo "Loading PostGIS extensions into $DB"
psql --dbname="$DB" <<-'EOSQL'
  INSERT INTO tiger.loader_platform(os, declare_sect, pgbin, wget, unzip_command, psql, path_sep,
        loader, environ_set_command, county_process_command)
  SELECT 'moogle', 'TMPDIR="${staging_fold}/temp/"
UNZIPTOOL=unzip
WGETTOOL="/usr/bin/wget"
export PGBIN=/usr/local/bin
export PGPORT=5432
export PGHOST=localhost
PSQL=${PGBIN}/psql
SHP2PGSQL=shp2pgsql
cd ${staging_fold}
', pgbin, wget, unzip_command, psql, path_sep,
      loader, environ_set_command, county_process_command
    FROM tiger.loader_platform
    WHERE os = 'sh';
EOSQL

# SELECT loader_generate_nation_script('moogle');
