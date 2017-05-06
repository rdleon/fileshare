GET    /             - Show the password prompt
POST   /login        - Login the user to the app
GET    /logout       - Logout the user from the app
GET    /archives     - List archives
POST   /archives     - Add an archive to the list
GET    /archives/:id - Download the file
PUT    /archives/:id - Refresh the expiration and name of the file
DELETE /archives/:id - Delete the file from the server
