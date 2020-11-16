#!/bin/sh
set -e

./wait-for.sh "$@"

# Sleep long enough to keep the process from exiting
sleep 1h