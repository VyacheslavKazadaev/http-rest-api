#!/usr/bin/env bash
docker pull dpage/pgadmin4
docker run --rm --network host \
    -v pgadmin4_root:/var/lib/pgadmin \
    -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' \
    -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' \
    -e 'PGADMIN_LISTEN_PORT=5050' \
    -d dpage/pgadmin4