# Descripcion 
This server shortens your URL to a 10 character string. This code was written on Golang and use RestAPI and grpc protocole. To store data code uses Postgresql database with pgx driver, also server can store data in memory.

# Install
For install and build execute command below:
```
go install github.com/apix76/ShortenURL@latest
```

# Setup config 
Table 
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

