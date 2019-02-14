#!/usr/bin/env bash
scp -o LogLevel=quiet ./bin/backup-opnsense ${REMOTE_USER}@${REMOTE_SERVER}:/usr/local/bin/backup-opnsense
ssh -o LogLevel=quiet ${REMOTE_USER}@${REMOTE_SERVER} "/bin/chmod +x /usr/local/bin/backup-opnsense"
