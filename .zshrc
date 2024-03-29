# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH="/home/shen/.oh-my-zsh"

export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin:$HOME/bin

# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
ZSH_THEME="robbyrussell"

# Set list of themes to pick from when loading at random
# Setting this variable when ZSH_THEME=random will cause zsh to load
# a theme from this variable instead of looking in $ZSH/themes/
# If set to an empty array, this variable will have no effect.
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion.
# Case-sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to automatically update without prompting.
# DISABLE_UPDATE_PROMPT="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line if pasting URLs and other text is messed up.
# DISABLE_MAGIC_FUNCTIONS="true"

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# You can set one of the optional three formats:
# "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# or set a custom format using the strftime function format specifications,
# see 'man strftime' for details.
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load?
# Standard plugins can be found in $ZSH/plugins/
# Custom plugins may be added to $ZSH_CUSTOM/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8
setopt no_nomatch
# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"
function s:audio(){
  systemctl --user restart pulseaudio
}

alias prime-run="__NV_PRIME_RENDER_OFFLOAD=1 __GLX_VENDOR_LIBRARY_NAME=nvidia __VK_LAYER_NV_optimus=NVIDIA_only"
alias zsh="source ~/.zshrc"
alias pid="ps aux | awk '{print \$2 \"\\t\" \$11}' | grep  $1"
alias s:docker="sudo systemctl start docker"
alias docker:cls="docker image prune -a"

alias wayne="cd ~/projects/servermodule"
alias c:wayne="code ~/projects/servermodule"

alias c:egame="code ~/projects/paradise/egame_deploy.code-workspace"

alias paradise="cd ~/projects/paradise"
alias c:paradise="code ~/projects/paradise/kkgame.code-workspace"
alias notes="cd ~/projects/paradise/notes"
alias c:notes="code ~/projects/paradise/notes"

alias s:vpn="echo zxc123 | sudo -S openvpn ~/work/pfSense-UDP4-1194/pfSense-UDP4-1194-config.ovpn"
# main project
alias pegame="cd ~/projects/paradise/fortest/egame-service" 
alias c:pegame="code ~/projects/paradise/fortest/egame-service" 
alias psvc="cd ~/projects/paradise/fortest/service_template" 
alias c:psvc="code ~/projects/paradise/fortest/service_template"
alias psvci="docker images | grep demeter- | awk '{print \$3}'"
function s:psvc(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/local/service/docker-compose.yml up -d
}
function u:psvc(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/local/service/docker-compose.yml pull
}
function d:psvc(){
  sudo docker-compose -f $HOME/projects/paradise/fortest/service_template/local/service/docker-compose.yml down -v
}
alias pgc="cd ~/projects/paradise/gamecontroller"
alias c:pgc="code ~/projects/paradise/gamecontroller"
alias pgc2="cd ~/projects/paradise/gameconnector"
alias c:pgc2="code ~/projects/paradise/gameconnector"
alias pgm="cd ~/projects/paradise/gamemaster"
alias c:pgm="code ~/projects/paradise/gamemaster"
function s:pgm(){
  docker start gamemaster
}
function d:pgm(){
  docker stop gamemaster
}
alias prs="cd ~/projects/paradise/roomservice"
alias c:prs="code ~/projects/paradise/roomservice"
alias pbe="cd ~/projects/paradise/backendmodules"
alias c:pbe="code ~/projects/paradise/backendmodules"
function u:pbe(){
  go get gitlab.geax.io/demeter/backendmodules@$1
  go mod tidy
  go mod vendor
}
alias pgs="cd ~/projects/paradise/gameservice"
alias c:pgs="code ~/projects/paradise/gameservice"
function s:pgs(){
  docker start gameservice
}
function d:pgs(){
  docker stop gameservice
}
alias pgl="cd ~/projects/paradise/gologger"
alias c:pgl="code ~/projects/paradise/gologger"
function u:pgl(){
  go get gitlab.geax.io/demeter/gologger@$1
  go mod tidy
  go mod vendor
}
alias pgo="cd ~/projects/paradise/gamecore"
alias c:pgo="code ~/projects/paradise/gamecore"
function u:pgo(){
  go get gitlab.geax.io/demeter/gamecore@$1
  go mod tidy
  go mod vendor
}
alias ppo="cd ~/projects/paradise/protocol"
alias c:ppo="code ~/projects/paradise/protocol"
function u:ppo(){
  go get gitlab.geax.io/demeter/protocol@$1
  go mod tidy
  go mod vendor
}

alias pdoc="cd ~/projects/paradise/fortest/documentation"
alias c:pdoc="code ~/projects/paradise/fortest/documentation"

alias psch="cd ~/projects/paradise/scheduler"
alias c:psch="code ~/projects/paradise/scheduler"

alias pit="cd ~/projects/paradise/fortest/integrationtesting"
alias c:pit="code ~/projects/paradise/fortest/integrationtesting"

alias psdk="cd ~/projects/paradise/game-sdk"
alias c:psdk="code ~/projects/paradise/game-sdk"

alias psw="cd ~/projects/paradise/streamingworker"
alias c:psw="code ~/projects/paradise/streamingworker"

alias pvc="cd ~/projects/paradise/versioncontroller"
alias c:pvc="code ~/projects/paradise/versioncontroller"

alias pad="cd ~/projects/paradise/game-management/admin"
alias c:pad="code ~/projects/paradise/game-management/admin"

alias pre="cd ~/projects/paradise/game-management/report"
alias c:pre="code ~/projects/paradise/game-management/report"

alias pnpa="cd ~/projects/paradise/game-management/node-public-api"
alias c:pnpa="code ~/projects/paradise/game-management/node-public-api"

alias pgpa="cd ~/projects/paradise/game-management/go-public-api"
alias c:pgpa="code ~/projects/paradise/game-management/go-public-api"

alias pghb="cd ~/projects/paradise/game-management/game-hall-backend"
alias c:pghb="code ~/projects/paradise/game-management/game-hall-backend"

alias pghw="cd ~/projects/paradise/game-management/game-hall-web"
alias c:pghw="code ~/projects/paradise/game-management/game-hall-web"

alias pgrw="cd ~/projects/paradise/game-management/game-records-web"
alias c:pgrw="code ~/projects/paradise/game-management/game-records-web"

alias pgrb="cd ~/projects/paradise/game-management/game-records-backend"
alias c:pgrb="code ~/projects/paradise/game-management/game-records-backend"

alias pmi="cd ~/projects/paradise/fortest/migration"
alias c:pmi="cd ~/projects/paradise/fortest/migration"
function pmi_tar() {
  cd ~/projects/paradise/fortest
  tar zcvf db_migration.tar.gz ./migration
  mv db_migration.tar.gz ~/projects/paradise/sql/kkgame/db_migration.tar.gz
}

alias pmi2="cd ~/projects/paradise/migration"
alias c:pmi2="code ~/projects/paradise/migration"
function u:pmi() {
  cd ~/projects/paradise/migration
  CGO_ENABLED=0 go build -o migration
  cp ./migration ~/projects/paradise/sql/kkgame/migration_$1
  file ~/projects/paradise/sql/kkgame/migration_$1
}

alias pds="cd ~/projects/paradise/game-management/demo-site"
alias c:pds="code ~/projects/paradise/game-management/demo-site"

alias pwt="cd ~/projects/paradise/game-management/webdevtool"
alias c:pwt="code ~/projects/paradise/game-management/webdevtool"

alias psql="cd ~/projects/paradise/sql"
alias c:psql="code ~/projects/paradise/sql"

alias argocd="cd ~/projects/paradise/argocd-kkgame"
alias c:argocd="code ~/projects/paradise/argocd-kkgame"

alias pts="cd ~/projects/paradise/transfer-service"
alias c:pts="code ~/projects/paradise/transfer-service"

alias d:nginx="systemctl stop nginx.service"

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

function versions(){
  echo "gamemaster:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/122/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "gameservice:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/257/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "gamecontroller:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/121/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "streamingworker:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/432/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "scheduler:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/160/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
}

function versions_web(){
  echo "game-hall-web:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/107/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "game-records-web:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/218/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "demo-site:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/459/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
}

function versions_api(){
  echo "game-hall-backend:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/108/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "game-records-backend:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/219/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "node-public-api:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/127/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "Admin:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/371/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "report:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/363/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
  echo "transfer-service:"
  curl --header "PRIVATE-TOKEN: A1GkKDf9uW8TD2wWzZZh" "https://gitlab.geax.io/api/v4/projects/348/repository/commits/master"
  echo "\n----------------------------------------------------------------------"
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

function login(){
  if [ -z "$1" ];then
    echo "must input one of [27, 36, 37, 41, 61, 75, 151]"
    return
  fi
  if [ $1 = '27' ];then
    sshpass -p "password" ssh -o StrictHostKeyChecking=no root@10.200.6.27
    return
  fi
  if [ $1 = '36' ];then
    sshpass -p "password" ssh -o StrictHostKeyChecking=no root@10.200.6.36
    return
  fi
  if [ $1 = '37' ];then
    sshpass -p "password" ssh -o StrictHostKeyChecking=no root@10.200.6.37
    return
  fi
  if [ $1 = '75' ];then
    sshpass -p "password" ssh -o StrictHostKeyChecking=no root@10.200.6.75
    return
  fi
  if [ $1 = '41' ];then
    sshpass -p "12345678" ssh -o StrictHostKeyChecking=no thai@10.200.6.41
    return
  fi
  if [ $1 = '61' ];then
    sshpass -p "12345678" ssh -o StrictHostKeyChecking=no thai@10.200.6.61
    return
  fi
  if [ $1 = '151' ];then
    sshpass -p "12345678" ssh -o StrictHostKeyChecking=no thai_wu@10.200.6.151
    return
  fi
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
git config --global user.email "wayne_shen@paradise-soft.com.tw"
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc'; fi
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"
