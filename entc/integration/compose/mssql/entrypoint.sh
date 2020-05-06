#!/bin/bash
set -e

/usr/src/app/import-data.sh &

/opt/mssql/bin/sqlservr