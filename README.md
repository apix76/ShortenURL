# Descripcion 
This server shortens your URL to a 10 character string. This code was written on Golang and use RestAPI and grpc protocole. To store data code uses Postgresql database with pgx driver, also server can store data in memory.

# Install
For install and build execute command below:
```
go install github.com/apix76/ShortenURL@latest
```
# Postgrasql
Setup of table psql is been in the [CreateDb.sql](https://github.com/apix76/ShortenURL/blob/main/CreateDb.sql).

# Setup config 

if field HttpPort/GrpcPort is empty, than RestAPI/grpc server won't start.
HttpPort & GrpcPort both can't be empty.
If field PgsqlNameServe is empty in config.cfg, server will be store data in memory.

Example config.cfg:
```json
{
  "HttpPort":":8080",
  "GrpcPort":":8081",
  "PgsqlNameServe":"postgres://PsqlUserName:PsqlUserPassword@DOMEN:PORT/NameYourDb"
}
```

Put config file in work directory.

# Example RestAPI request 
Post:
```cmd
curl -X POST http://domen:port -d {"URL":"YourURL"}
resp: {"ShortURL":"shortenrul"}
```
Get:
```cmd
curl -X GET http://domen:port -d {"ShortURL":"shortenrul"}
resp: {"URL":"YourURL"}
```
