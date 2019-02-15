#!/usr/bin/env bash
scp -o LogLevel=quiet ./bin/quackup-opnsense ${REMOTE_USER}@${REMOTE_SERVER}:/usr/local/bin/quackup-opnsense
ssh -o LogLevel=quiet ${REMOTE_USER}@${REMOTE_SERVER} "/bin/chmod +x /usr/local/bin/quackup-opnsense"
