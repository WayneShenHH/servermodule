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
alias zsh="source ~/.zshrc"
alias toolsgo="cd ~/projects/toolsgo"
alias pid="ps aux | awk '{print \$2 \"\\t\" \$11}' | grep  $1"
alias paradise="cd ~/projects/paradise"
alias notes="cd ~/projects/paradise/notes"
alias c:notes="code ~/projects/paradise/notes"
alias wayne="cd ~/projects/servermodule"
alias c:wayne="code ~/projects/servermodule"
alias c:service="code ~/projects/paradise/service-template"
alias s:docker="sudo systemctl start docker"

alias gcpredis="gcloud compute ssh fsbs-forwarder --zone asia-east1-b -- -N -L 6386:10.0.0.3:6379"
alias gcpredisl1="kubectl -n=lab1 port-forward service/redis 6386:6379"
alias gcpnsq="kubectl -n=fsbs port-forward nsqlookupd-4k5qt 4161:4161"
alias gcplog="kubectl port-forward svc/kibana 5601:443 -n=logging"
alias nsqlook="cd ~/nsqlog && nsqlookupd"
alias nsq="cd ~/nsqlog && nsqd --lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1"
alias nsqui="cd ~/nsqlog && nsqadmin --lookupd-http-address=127.0.0.1:4161"

alias lintfix="golangci-lint run --fix"

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

function doloop(){
  for i in {1..$1};
do
    # echo "doloop"
    # sh -c $2
    bddtest odds2
done
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
  go clean -testcache && go test -v -timeout 50s $1 -run $2
}

function bentest(){
  go test -v $1 -bench=$2
}

function unitest(){
  go test -cover $(go list ./... | grep -E -v "vendor|integration|wayne")
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

function s:rediscmd(){
  sudo docker run --rm --name redis-commander -d \
  -p 18081:8081 \
  rediscommander/redis-commander:latest
}

function s:redis(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/redis/docker-compose.yml up -d
}

function s:nats(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/nats/docker-compose.yml up -d
}

function s:arango(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/arangodb/docker-compose.yml up -d
}

function s:server(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/service/docker-compose.yml up -d
}

function d:redis(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/redis/docker-compose.yml down -v
}

function d:nats(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/nats/docker-compose.yml down -v
}

function d:arango(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/arangodb/docker-compose.yml down -v
}

function d:server(){
  sudo docker-compose -f $HOME/projects/paradise/service-template/local/service/docker-compose.yml down -v
}

function commit(){
  git commit -m "feat($1): $2"
  git push
}

function ts(){
  date --date=@$1
}

# xmodmap ~/.Xmodmap
git config --global user.name "wayne_shen"
git config --global user.email "wayne_shen@paradise-soft.com.tw"