[![Build Status](https://travis-ci.com/mike-seagull/quackup-opnsense.svg?branch=master)](https://travis-ci.com/mike-seagull/quackup-opnsense)

<p>backs up OPNSense config file to a remote server using scp and a ssh-key<p/>
<p>files are saved on the remote server as <code>config.${EPOCH_TIME}.xml</code></p>
<h4>Required Environment Variables</h4>

* SERVER_USER = user on remote server
* SERVER_IP = ip to the remote server

<h5>To start a backup</h5>
<code>quackup-opnsense /remote/backup/dir/</code>
<h5>Example Crontab</h5>
<code>1 2 * * * (bash -c "SERVER_USER=remoteuser SERVER_IP=your.remote.server.ip quackup-opnsense /remote/backup/dir") >> /local/log/path/quackup-opnsense.log</code>
