# Fileshare
## For sending things to your friends

A simple, lightweight server for sharing files, it stores files
temporarily in a server and generates unique URLs to share to
your friends and family!

## Usage:

To use, modify the config file, set listen to "\*:80" or the
port you want to use, add a password and a secret (a good ol'
random string) and run:

	$ fileshare [-conf config.json]

You can open the web interface by going to the configured
(in config.json) location (defaults to http://localhost:8080 .)

## Routes:

* GET    /             - Show the password prompt
* POST   /login        - Login the user to the app
* GET    /logout       - Logout the user from the app
* GET    /archives     - List archives
* POST   /archives     - Add an archive to the list
* GET    /archives/:id - Download the file
* PUT    /archives/:id - Refresh the expiration time of the file (+24hrs)
* DELETE /archives/:id - Delete the file from the server

### Note:

All the info lives in memory, so it will be reset to blank
after a reboot or server restart.
