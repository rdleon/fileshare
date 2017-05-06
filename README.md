# Fileshare for sending things to your friends

A simple, lightweight server for sharing files, it stores files
temporarily in a server and generates unique URLs to share to
your friends and family!

## Usage:

	$ fileshare -c config.json

You can open the web interface by going to the configured
(in config.json) location (defaults to http://localhost:8080 .)

### Note:

All the info lives in memory, so it will be reset to blank
after a reboot or server restart.
