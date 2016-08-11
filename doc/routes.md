GET    /          - Show the password prompt
POST   /login     - Login the user to the app
GET    /logout    - Logout the user from the app
GET    /files     - List archives
POST   /files     - Add an archive to the list
GET    /files/:id - Download the file
PUT    /files/:id - Modify the expiration and name of the file
DELETE /files/:id - Delete the file
