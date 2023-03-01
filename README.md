# app_test
api golang para pruebas rápidas

Ejecutar Nginx en docker
volumen 
docker volume create vol-nginx-config 
docker volume ls
docker volume inspect vol-nginx-config

docker run --name my-nginx-container \
-v vol-nginx-config:/etc/nginx \
-d -p 8000:80 nginx

docker run --name my-nginx-container \
--mount type=bind,source="$(pwd)"/nginx/nginx.conf,target=/etc/nginx/nginx.conf \
--network host \
-d -p 8000:80 nginx




docker run --name my-nginx-container -v ./nginx.conf:/etc/nginx/nginx.conf:ro -d -p 8000:80 nginx
se creara un archivo vinculado al archivo config del contenedor para editar la configuración

http://localhost:8001

