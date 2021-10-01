sudo groupadd docker

sudo usermod -aG docker $USER

newgrp docker

docker build -t books_online:1.0.1 .

docker login
docker tag 557bec4698b6 moryku/kampus_merdeka:1.0.0
docker push moryku/kampus_merdeka:1.0.0

docker image pull moryku/kampus_merdeka:1.0.3

docker run -d -it -p 8080:8080 --name=books_online3 wahyunf354/books_online:1.0.3
