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
sudo curl -L "https://github.com/docker/compose/releases/download/v2.4.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version
```

docker login harbor
```s
docker login --username 帳號 --password 密碼 img.paradise-soft.com.tw
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

docker run mysql workbench
```s
docker run -d \
  --name=mysql-workbench \
  --cap-add=IPC_LOCK \
  -e PUID=1000 \
  -e PGID=1000 \
  -e TZ=Etc/UTC \
  -p 3000:3000 \
  -p 3001:3001 \
  -v /path/to/config:/config \
  --shm-size="1gb" \
  --restart unless-stopped \
  lscr.io/linuxserver/mysql-workbench:latest
```