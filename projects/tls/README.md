# Reminder

## How to run
```
cd projects/tls
go run main.go
```


## Generate private key
```
openssl genrsa -out private.key 2048
```
## Generate public key
```
openssl req -new -x509 -sha256 -key private.key -out public.crt -days 3650
```

Country Name (2 letter code) [AU]:BR
State or Province Name (full name) [Some-State]:Minas Gerais
Locality Name (eg, city) []:Uberlandia                                            
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Learning Go
Organizational Unit Name (eg, section) []:dev
Common Name (e.g. server FQDN or YOUR name) []:learning-go.dev
Email Address []: