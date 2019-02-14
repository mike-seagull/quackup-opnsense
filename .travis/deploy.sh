#!/usr/bin/env bash

scp $(pwd)/bin/backup-opnsense ${REMOTE_USER}@${REMOTE_SERVER}:/usr/local/bin/backup-opnsense
ssh ${REMOTE_USER}@${REMOTE_SERVER} /bin/chmod +x /usr/local/bin/backup-opnsense
