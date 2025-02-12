# Descripcion 
This server shortens your URL to a 10 character string. This code was written on Golang and use RestAPI and grpc protocol. To store data code uses Postgresql database with pgx driver, also server can store data in memory.

# How get 

For get code execute command below:
```
go clone github.com/apix76/ShortenURL@latest "Your_path"
```
## Build with Go build
For build code with go you can command below:
```
go build location_ShortenURL
```

## Build with Docker  

For build code with docker you can use:
```
docker build -t shortenurl location_ShortenURL
```

This comman is create image from [Dockerfile](https://github.com/apix76/ShortenURL/blob/main/Dockerfile), which you can run with:
```
docker run -p HttpPort:HttpPort -p GrpcPort:GrpcPort shortenurl
```

# Postgresql
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
