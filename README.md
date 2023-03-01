# app_test
api golang para pruebas rápidas

Ejecutar Nginx en docker

docker run --name my-nginx-container \
--mount type=bind,source="$(pwd)"/nginx/nginx.conf,target=/etc/nginx/nginx.conf \
--network host \
-d -p 8000:80 nginx



se creara un archivo vinculado al archivo config del contenedor para editar la configuración

http://localhost:8001

