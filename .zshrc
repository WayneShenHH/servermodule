# If you come from bash you might have to change your $PATH.
export ZSH="/home/shen/.oh-my-zsh"
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin:$HOME/bin
export PATH=$PATH:$HOME/.local/bin

ZSH_THEME="robbyrussell"
export GOOGLE_CLOUD_PROJECT="940773501583"
export GEMINI_API_KEY="AIzaSyDqIsX1oaRerXZDHmuKneGsvvSaNLYPidM"

# go build 使用了 c語言的 lpsolve 函式庫，需要設定編譯參數
# export	CGO_CFLAGS="-I/usr/include/lpsolve"
# export	CGO_LDFLAGS="-llpsolve55 -lm -ldl -lcolamd"

plugins=(git)

source $ZSH/oh-my-zsh.sh

setopt no_nomatch

function s:audio(){
  systemctl --user restart pulseaudio
}

alias goland="~/go/bin/jetbrains-toolbox"

# alias prime-run="__NV_PRIME_RENDER_OFFLOAD=1 __GLX_VENDOR_LIBRARY_NAME=nvidia __VK_LAYER_NV_optimus=NVIDIA_only"
alias zsh="source ~/.zshrc"
alias pid="ps aux | awk '{print \$2 \"\\t\" \$11}' | grep  $1"
alias s:docker="sudo systemctl start docker"
alias docker:cls="docker image prune -a"

alias wayne="cd ~/projects/servermodule"
alias c:wayne="code ~/projects/servermodule"

alias fgw="cd ~/projects/fgw"
alias c:fgw="code ~/projects/fgw/kkgame.code-workspace"

alias s:vpn="echo zxc123 | sudo -S openvpn ~/work/pfSense-UDP4-1194/pfSense-UDP4-1194-config.ovpn"
# main project
alias egame="cd ~/projects/fgw/deploy/egame-service" 
alias c:egame="code ~/projects/fgw/deploy/egame-service" 

alias argocd="cd ~/projects/fgw/deploy/argocd-kkgame"
alias c:argocd="code ~/projects/fgw/deploy/argocd-kkgame"

alias pmi="cd ~/projects/fgw/deploy/migration"
alias c:pmi="code ~/projects/fgw/deploy/migration"
function u:pmi() {
  cd ~/projects/fgw/deploy/migration
  CGO_ENABLED=0 go build -o migration
  cp ./migration ~/projects/fgw/deploy/sql/kkgame/migration_$1
  file ~/projects/fgw/deploy/sql/kkgame/migration_$1
}

alias pvc="cd ~/projects/fgw/deploy/versioncontroller"
alias c:pvc="code ~/projects/fgw/deploy/versioncontroller"

alias pgc="cd ~/projects/fgw/gameconnector"
alias c:pgc="code ~/projects/fgw/gameconnector"

alias pbe="cd ~/projects/fgw/backendmodules"
alias c:pbe="code ~/projects/fgw/backendmodules"
function u:pbe(){
  go get gitlab.geax.io/demeter/backendmodules@$1
  go mod tidy
  go mod vendor
}
alias pgs="cd ~/projects/fgw/gameservice"
alias c:pgs="code ~/projects/fgw/gameservice"

alias pgas="cd ~/projects/mbg/game-api-service"
alias c:pgas="code ~/projects/mbg/game-api-service"

alias psdk="cd ~/projects/fgw/cocos/game-sdk"
alias c:psdk="code ~/projects/fgw/cocos/game-sdk"

alias kkgame="cd ~/projects/paradise/kkgame"

function u:sdk(){
  nvm use 16.20.2
  psdk
  npm run build
  cp ./dist/sdk_polyfill.js ~/projects/paradise/kkgame/Common/SDK/sdk_polyfill.js
  cd ~/projects/paradise/kkgame/Common
  gpa "refactor(sdk): update sdk"
}

alias pgpa="cd ~/projects/fgw/go-public-api"
alias c:pgpa="code ~/projects/fgw/go-public-api"

alias pghw="cd ~/projects/fgw/web/game-hall-web"
alias c:pghw="code ~/projects/fgw/web/game-hall-web"

alias pgm="cd ~/projects/fgw/web/gamemanual"
alias c:pgm="code ~/projects/fgw/web/gamemanual"

alias pgrw="cd ~/projects/fgw/web/game-records-web"
alias c:pgrw="code ~/projects/fgw/web/game-records-web"

alias pds="cd ~/projects/fgw/web/demo-site"
alias c:pds="code ~/projects/fgw/web/demo-site"

alias pgt="cd ~/projects/fgw/web/gametool"
alias c:pgt="code ~/projects/fgw/web/gametool"

alias d:nginx="systemctl stop nginx.service"

alias s:ibus="ibus-daemon &"

function add:crt(){
    cp $1 /usr/local/share/ca-certificates
    sudo update-ca-certificates
}

alias gmod="git submodule update --init --recursive"
alias gh="git rev-parse --short=8 HEAD"
alias gh6="git rev-parse --short=6 HEAD"
alias gtag="git describe --tag"
alias gtd="git tag -d $1"
alias gtdp="git push --delete origin $1"
alias gcpredis="gcloud compute ssh fsbs-forwarder --zone asia-east1-b -- -N -L 6386:10.0.0.3:6379"
alias gcpredisl1="kubectl -n=lab1 port-forward service/redis 6386:6379"
alias gcpnsq="kubectl -n=fsbs port-forward nsqlookupd-4k5qt 4161:4161"
alias gcplog="kubectl port-forward svc/kibana 5601:443 -n=logging"
alias nsqlook="cd ~/nsqlog && nsqlookupd"
alias nsq="cd ~/nsqlog && nsqd --lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1"
alias nsqui="cd ~/nsqlog && nsqadmin --lookupd-http-address=127.0.0.1:4161"
alias lintfix="golangci-lint run --fix"
alias c:zsh="code ~/.zshrc"
alias c:host="sudo vi /etc/hosts"
alias l:host="cat /etc/hosts"
alias redisui="nohup redisdm >/dev/null &"
alias kc="kubectl"
alias sshkey="cat ~/.ssh/id_rsa.pub"
alias goclear="go clean --modcache"
alias jbox="~/go/bin/jbox"
alias dockercls="docker system prune"
alias snipasteUI="snipaste >/dev/null &"
alias wu="wayneutil"

function wu(){
  wayneutil $1 $2 $3
}

function repair:zsh(){
  cd ~
  mv .zsh_history .zsh_history_bad
  strings -eS .zsh_history_bad > .zsh_history
  fc -R .zsh_history
}

function gud(){
  pgpa 
  gco dev
  gpr

  pre
  gco dev
  gpr

  pgs
  gco dev
  gpr

  pgc
  gco dev
  gpr

  pbe
  gco dev
  gpr

  ppo
  gco dev
  gpr
}

function gum(){
  pgpa 
  gco master
  gpr

  pre
  gco master
  gpr

  pgs
  gco master
  gpr

  pgc
  gco master
  gpr

  pbe
  gco master
  gpr

  ppo
  gco master
  gpr

  pmi
  gco master
  gpr

  pegame
  gco master
  gpr

  pvc
  gco master
  gpr
}

function hostch(){
  if [ -z "$1" ];then
    echo "must input one of [27, 36, 37, 49, 75, 151]"
    return
  fi

  sudo sed -i '$d' /etc/hosts
  sudo sed -i "$ a 10.200.6.$1 reverse-proxy.sit-gm.svc.cluster.local" /etc/hosts
  cat /etc/hosts
}

function u:npm(){
  rm -rf ./node_modules
  rm package-lock.json
  npm i
}

function pprof:heap(){
  # 傳入 service 網址(ex: http://localhost:6060)，後方路徑(/debug/pprof/heap)為預設，可能會依照各服務改寫
  go tool pprof $1/debug/pprof/heap
}
function pprof:ui(){
  # 傳入快照檔案路徑
  go tool pprof -http=:8080 $1
}

function arango_docker_import(){
  # arango_docker_import /home/shen/Downloads/20250904_arangodb AccountStatistics account_statistics.csv
  # 指定資料夾預先放置好檔案
  docker run --rm \
  -v $1:/dump \
  arangodb/arangodb \
  arangoimport \
  --server.endpoint tcp://10.127.6.12:8529 \
  --server.database Converter \
  --server.username root \
  --server.password password \
  --collection $2 \
  --file "/dump/$3" --type csv
}

function arango_docker_restore(){
  # 指定資料夾預先放置好檔案
  docker run --rm \
  -v $1:/dump \
  arangodb/arangodb \
  arangorestore \
  --server.endpoint tcp://10.127.6.12:8529 \
  --server.database Database \
  --server.username root \
  --server.password password \
  --input-directory /dump
}

function arango_docker_dump(){
  # 輸出指定資料夾
  docker run --rm \
  -v $1:/dump \
  arangodb/arangodb \
  arangodump \
  --server.endpoint tcp://10.127.6.12:8529 \
  --server.database Database \
  --server.username root \
  --server.password password \
  --overwrite true \
  --output-directory /dump
}

function arango_restore(){
  # 指定資料夾預先放置好檔案，包含兩個檔案 (*.data.json | *.data.json.gz) & *.structure.json
  arangorestore --server.authentication=false --create-database=true --server.database "Database" \
--input-directory "$1"
}

function dbimport(){
  arangoimport --file "$1" --type jsonl --server.database "Database" --collection "$2"
}

function devdump(){
  outpath=./lastestDump.tar.gz
  if [ -z "$1" ];then
    echo "default output to the same place"
  else
    outpath=$1/lastestDump.tar.gz
    if [ $1 = 'loc' ];then
      outpath=$HOME/projects/paradise/fortest/service_template/local/arangodb/lastestDump.tar.gz
    fi
  fi
  echo "output to:" $outpath

  DEV="10.200.6.37"
  DB_URL=http://10.200.6.37:8888/job/DEV/job/build-image-and-deploy-service/job/arangodb-dump/job/arangoDB-dump/ws/$DEV/lastestDump.tar.gz
  wget --auth-no-challenge \
	--user=admin --password=1161eb66da93f00301bdeab20c8decfd71 \
	"$DB_URL" \
    -O $outpath
}

function mgd(){
  # https://mega.nz/file/TZ8nFKzL#NfwZWTabsBUG3GyKbZStV_yslg73ENQ3OEQhd-KPGDc
  # https://mega.co.nz/#!TZ8nFKzL!NfwZWTabsBUG3GyKbZStV_yslg73ENQ3OEQhd-KPGDc
  echo $1
  url=${1//'#'/'!'} # replace sharp to exclamation mark
  ori="mega.nz/file/"
  rep="mega.co.nz/#!"
  link=${url//$ori/$rep}
  echo $link
  megadl $link --path=./mega
}

function gbuild() {
  cd ~/projects/paradise/roomservice/games/game$1/algorithm-test
  CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ~/projects/paradise/fortest/gamesimulation/game$1_simulation/simulation.exe
  go build -o ~/projects/paradise/fortest/gamesimulation/game$1_simulation/simulation
  prs
}

function compress(){
  # $1: destination file name
  # $2: source diretory
  tar zcvf $1.tar.gz $2
}

function fgw:clone(){
    cd ~/projects/fgw
    git clone https://gitlab.fgw/egame/backendmodules.git
    git clone https://gitlab.fgw/egame/go-public-api.git
    git clone https://gitlab.fgw/egame/gameconnector.git
    git clone https://gitlab.fgw/egame/gameservice.git

    cd ~/projects/fgw/deploy
    git clone https://gitlab.fgw/deploy/versioncontroller.git
    git clone https://gitlab.fgw/deploy/migration.git
    git clone https://gitlab.fgw/deploy/egame-service.git
    git clone https://gitlab.fgw/deploy/automation-testing.git
    git clone http://gitlab.paradise-soft.com.tw/deployscript/devops/utils/argocd-kkgame.git

    cd ~/projects/fgw/web
    git clone https://gitlab.fgw/egame/web/game-hall-web.git
    git clone https://gitlab.fgw/egame/web/gamemanual.git
    git clone https://gitlab.fgw/egame/web/game-records-web.git
    git clone https://gitlab.fgw/egame/web/demo-site.git
    git clone https://gitlab.fgw/egame/web/gametool.git
}

function sucode(){
  sudo code $1 --user-data-dir
}

function gostr(){
  filepath=$1
  file=$(echo $filepath |rev|cut -d "/" -f 1 |rev)
  echo $file

  ver1=$(echo $file |cut -d "." -f 1 )
  ver2=$(echo $file |cut -d "." -f 2 )
  ver3=$(echo $file |cut -d "." -f 3 )
  ver=$(echo $ver1.$ver2.$ver3)
  echo $ver

  ver4=$(echo $file |cut -d "-" -f 1 )
  echo $ver4
}

function gvi() {
  echo "source: $1"
  filepath=$1
  file=$(echo $filepath |rev|cut -d "/" -f 1 |rev)

  ver=$(echo $file |cut -d "-" -f 1 )
  echo "version: $ver"

  godir="/usr/local/go"
  verdir="/usr/local/gvm/$ver"
  if [ -d $godir ]; then
      sudo rm -rf $godir
      echo "uninstall previous version"
  fi

  if [ -d $verdir ]; then
    sudo rm -rf $verdir
    echo "remove old $verdir"
  fi
  sudo mkdir -p $verdir
  if [ $? != 0 ];then
    echo "fail to install $ver..."
    return
  fi

  sudo tar -C $verdir -xzf $1
  if [ $? != 0 ];then
    echo "fail to install $ver..."
    return
  fi
  echo "unpackage to $verdir success"
  
  sudo cp -r $verdir/go /usr/local/go
  if [ $? != 0 ];then
    echo "fail to install $ver..."
    return
  fi

  echo "install $ver success"
  go version
}

function gvm() {
  if [ -z "$1" ];then
    echo "usage: gvm (ls | lastest | go-version)"
    return
  fi

  if [ $1 = 'ls' ];then
    ls /usr/local/gvm | awk '{print $1 "\t"}'
    return
  fi
  
  if [ $1 = 'lastest' ];then
    max="go1.1.1"
    all=$(ls /usr/local/gvm)
    array=($(echo $all | tr ' ' "\n"))
    for ver in "${array[@]}"
    do
      if [[ "$max" < "$ver" ]];then
        max=$ver        
      fi
    done
    echo "using $max as lastest"
    gvm $max
    return
  fi

  godir="/usr/local/go"
  verdir="/usr/local/gvm/$1"
  if [ -d $verdir ]; then
      if [ -d $godir ]; then
        sudo rm -rf $godir
        echo "delete $godir"
      fi
      sudo cp -r $verdir/go $godir
    else
      echo "$1 not exist..."
      return
  fi

  echo "switch to $1 success"
  go version
}

function gcpredis(){
  kubectl -n=$KUBE_NAME_SPACE port-forward pod/$1 6386:6379
}
function gcpfwd(){
  kubectl -n=$KUBE_NAME_SPACE port-forward pod/$1 16888:$2
}
function kubespace() {
  export KUBE_NAME_SPACE=$1
  echo "set kubernetes name-space to: "$1
}
function kubeurl(){
  kubectl get virtualservice -n=istio-system | grep $KUBE_NAME_SPACE
}
function kubesvc(){
  kubectl get svc -n=$KUBE_NAME_SPACE
}
function kubef(){
  sudo kubefwd svc -n=$KUBE_NAME_SPACE -d k8s.tw
}
function helmdel(){
  helm del --purge fsbs-$KUBE_NAME_SPACE
}
function kubeapply(){
  # kubectl apply -f /Users/wayne/projects/doc-devops/kubernetes/fsbs/$1/stbk.yaml
  kubectl -n=$1 apply -f /Users/wayne/projects/k8s/rd-lab/lab$1/$2.yaml
}
function kubepod(){
  kubectl get pods -n=$KUBE_NAME_SPACE | grep $1
  #kubectl -n=$1 get pods -l worker=stbk-$2
}
function kubelog(){
  kubectl -n=$KUBE_NAME_SPACE logs $1 $2
}
function kubedesc(){
  kubectl -n=$KUBE_NAME_SPACE describe pod $1
}
function kubedel(){
  kubectl delete pods -n=$KUBE_NAME_SPACE $1
}
function kubecfg(){
  kubectl -n=$KUBE_NAME_SPACE get configmaps $1 -o yaml
}

function dockerauth(){
  sudo groupadd docker
  sudo usermod -aG docker $USER
  newgrp docker
}

function gitcls(){
  git fetch
  git rebase
  git checkout master
  git pull
  if git branch --merged master | grep -v '^[ *]*master$';then
      git branch --merged master | grep -v '^[ *]*master$' | xargs git branch -d
  else
    echo "not found any merged branch"
  fi
  git fetch --prune
}

function gpm(){
  if [ -z "$1" ];then
    echo "need input a branch to merge"
    return
  fi
  cur=$(git branch --show-current)
  echo "current:" $cur
  echo "update:" $1 "\n"
  gco $1
  gup
  gco $cur
  echo "\n"
  echo $cur "merge from:" $1
  gm $1
}

function doloop(){
  for i in {1..$1};
  do
    # echo "doloop"
    # sh -c $2
    # npm run dev
    nc -zv 10.10.66.93 4222
    nc -zv 10.10.66.93 4223
    nc -zv 10.10.66.93 4224
    sleep 1s
  done
}

function extest(){
  cd ./games/game$1/example
  go build -o debug
  ./debug
  rm ./debug
  cd ../../..
}

function bddtest(){
  go clean -testcache && go test -v -timeout 50s ./$1 -tags integration -run $2
}

function shtest(){
  if git branch --merged master | grep -v '^[ *]*master$';then
    echo "ok" + $(git branch --merged master)
  else
    echo "not found" + $(git branch --merged master)
  fi
}

function svctest(){
  # go clean -testcache && go test -v -timeout 5s ./module/mq/mockqueue/mockqueue_test.go
  # go clean -testcache && go test -v -timeout 5s ./handlers -run Test_Name
  TRACE_LEVEL=debug go clean -testcache && go test -v -timeout 50h ./$1 -run $2
}

function bentest(){
  go test -v $1 -bench=$2
}

function unitest(){
  go test -v -cover $(go list ./... | grep -E -v "vendor|integration|wayne")
}

function lstest(){
  go test -v -cover $(go list ./$1/... | grep -E -v "vendor|integration|wayne")
}

function setps(){
  export PS=$1
}
function killps(){
  export PID=$(ps aux | awk '{print $2 "\t" $11}' | grep $PS | awk '{print $1}')
  kill -9 $PID
}

function goswagger(){
  swagger generate spec -m -o ~/projects/$1/swagger.json
  swagger serve -F=swagger ~/projects/$1/swagger.json
}

function docker-stop-all(){
  sudo docker stop $(docker ps -a -q)
}

function s:env(){
  s:arango
  s:redis
  s:nats
}

function d:env(){
  d:arango
  d:nats
  d:redis
}

function s:all(){
  s:arango
  s:redis
  s:nats
  s:psvc
}

function d:all(){
  d:arango
  d:nats
  d:redis
  d:psvc
}

function s:redis(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/wayne/redis/docker-compose.yml up -d
  # config cluster
  REDIS_CLUSTER_IP=$(hostname -I | awk '{ print $1}')
  redis-cli -c -p 7001 --cluster create $REDIS_CLUSTER_IP:7001 $REDIS_CLUSTER_IP:7002 $REDIS_CLUSTER_IP:7003 $REDIS_CLUSTER_IP:7004 $REDIS_CLUSTER_IP:7005 $REDIS_CLUSTER_IP:7006 --cluster-replicas 1
}

function s:nats(){
  sudo docker compose -f $HOME/projects/paradise/fortest/service_template/wayne/nats/docker-compose.yml up -d
}

function s:arango(){
  sudo docker compose -f $HOME/projects/paradise/fortest/service_template/tools/arangodb/docker-compose.yml up -d
}

# run while arangodb upgrade to new version ex. 3.7.0 => 3.8.0
function u:arango(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/wayne/arangodb/docker-compose.yml run --rm arangodb --database.auto-upgrade
}

function d:redis(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/wayne/redis/docker-compose.yml down -v
}

function d:nats(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/wayne/nats/docker-compose.yml down -v
}

function d:arango(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/wayne/arangodb/docker-compose.yml down -v
}

function commit(){
  git commit -m "feat($1): $2"
  git push
}

function gpa(){
  git add .
  git commit -m "$1"
  git push
}

function ts(){
  date --date=@$1
}

# xmodmap ~/.Xmodmap
git config --global user.name "wayne_shen"
git config --global user.email "wayne_shen@tengyuntech.com"
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc'; fi
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"