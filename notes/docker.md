# docker engine

```s
sudo apt-get update
# sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo apt-get install docker.io
docker version
```

docker config group 設定使用者權限

```s
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
```

docker-compose

```s
sudo curl -L "https://github.com/docker/compose/releases/download/1.10.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

docker run rediscommander (outsideport:internalport)

```s
docker run --rm --name redis-commander -d \
  -p 18081:8081 \
  rediscommander/redis-commander:latest
```

rediscommander connect to redis-server in docker: ip should be

```s
hostname -I | awk '{ print $1}'
```

start docker daemon

```s
sudo systemctl start docker
```

remove images (remove ps first)

```s
docker rmi $(docker ps -aq)
docker rmi $(docker images -aq)
```

docker run bash

```s
docker exec -it {docker-ps-id} (sh|bash)
```
