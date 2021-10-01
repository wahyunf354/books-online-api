sudo groupadd docker

sudo usermod -aG docker $USER

newgrp docker

docker build -t books_online:1.0.1 .
