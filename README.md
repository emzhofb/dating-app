# dating-app

how to run this

open terminal and clone this
```
if use ssh
git clone git@github.com:emzhofb/dating-app.git
if use http
git clone https://github.com/emzhofb/dating-app.git
```
and then move to dir and set up
```
cd dating-app
go mod init
go mod tidy
go mod vendor
```
and then go to the database mysql and create database name dating-app
```
 ConfigDB{
		DB_Username: "root",
		DB_Password: "password123",
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Database: "dating-app",
	}

this is the config inside configs/config.go
your database need to be set here, if username and password is different, please modify it in the code
```
and then run the main.go and table migration and seed will be created
```
go run main.go
```
finally, you just need to impor the `dating-app.postman_collection.json` to your json app and run it.
thanks!!!
