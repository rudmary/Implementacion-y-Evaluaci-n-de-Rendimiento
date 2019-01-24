# Proyecto distribuidos

Implementación de diseño para repartir la carga entre dos instancias.

# Development

- Depedencias GO para server de archivos estáticos y api

Gin

```sh
github.com/gin-gonic/gin
```

- Dependencias microservicio

Mysql

```sh
go get -u github.com/go-sql-driver/mysql
```

Redis

```sh
go get -u https://github.com/voidabhi/gredis
```

Grpc

```sh
go get -u google.golang.org/grpc
```

Go env

```sh
go get github.com/joho/godotenv/cmd/godotenv
```

- Instalación de redis

```sh
git clone https://github.com/google/protobuf.git .protobuf
cd  .protobuf
./autogen.sh
./configure
make
make install
make ldconfig
cd ..

```

# Instancias

- Ubuntu 16.04

## Instancia mysql

- Instalacion ngix(proxy server) - api gateway

```sh
sudo su
debconf-set-selections <<< 'mysql-server mysql-server/root_password password mysqldb'
debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password mysqldb'
apt-get update
apt-get install -y mysql-server
```

- Instalacion nginx

```sh
sudo apt-get install nginx-server
```

- Reemplazar el archivo de configuracion nginx en \_etc/nginx/sites-available/default

```txt
upstream microservicio {
    server localhost:3001;
	server localhost:3000;
}

server {
    listen 80;
    location /api {
        proxy_pass http://microservicio;
        proxy_next_upstream     error timeout invalid_header http_500;
        proxy_connect_timeout   2;
		proxy_set_header        Host            $host;
    }

    location / {
		proxy_pass "http://127.0.0.1:3005";
	}
}

```

- Correr los archivos

```sh
nohup ./server &
```
