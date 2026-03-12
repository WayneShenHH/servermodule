# If you come from bash you might have to change your $PATH.
export ZSH="/home/shen/.oh-my-zsh"
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin:$HOME/bin
export PATH=$PATH:$HOME/.local/bin

ZSH_THEME="robbyrussell"
export GEMINI_API_KEY=""

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

alias d:nginx="systemctl stop nginx.service"

alias s:ibus="ibus-daemon &"

function add:crt(){
    cp $1 /usr/local/share/ca-certificates
    sudo update-ca-certificates
}

alias gcpredis="gcloud compute ssh fsbs-forwarder --zone asia-east1-b -- -N -L 6386:10.0.0.3:6379"
alias gcpredisl1="kubectl -n=lab1 port-forward service/redis 6386:6379"
alias gcpnsq="kubectl -n=fsbs port-forward nsqlookupd-4k5qt 4161:4161"
alias gcplog="kubectl port-forward svc/kibana 5601:443 -n=logging"
alias nsqlook="cd ~/nsqlog && nsqlookupd"
alias nsq="cd ~/nsqlog && nsqd --lookupd-tcp-address=127.0.0.1:4160 -broadcast-address=127.0.0.1"
alias nsqui="cd ~/nsqlog && nsqadmin --lookupd-http-address=127.0.0.1:4161"
alias c:zsh="code ~/.zshrc"
alias c:host="sudo vi /etc/hosts"
alias l:host="cat /etc/hosts"
alias redisui="nohup redisdm >/dev/null &"
alias kc="kubectl"
alias sshkey="cat ~/.ssh/id_rsa.pub"
alias dockercls="docker system prune"
alias snipasteUI="snipaste >/dev/null &"
alias wu="wayneutil"

function add:dir() {
  if [ ! -d "$1" ]; then
    mkdir "$1"
  fi

  cd "$1"
}

function repair:zsh(){
  cd ~
  mv .zsh_history .zsh_history_bad
  strings -eS .zsh_history_bad > .zsh_history
  fc -R .zsh_history
}

function u:npm(){
  rm -rf ./node_modules
  rm package-lock.json
  npm i
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

function compress(){
  # $1: destination file name
  # $2: source diretory
  tar zcvf $1.tar.gz $2
}

function unzip() {
  tar -xzf $1
}

function sucode(){
  sudo code $1 --user-data-dir
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

function setps(){
  export PS=$1
}
function killps(){
  export PID=$(ps aux | awk '{print $2 "\t" $11}' | grep $PS | awk '{print $1}')
  kill -9 $PID
}

function docker-stop-all(){
  sudo docker stop $(docker ps -a -q)
}

export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion

# The next line updates PATH for the Google Cloud SDK.
if [ -f '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/path.zsh.inc'; fi

# The next line enables shell command completion for gcloud.
if [ -f '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc' ]; then . '/home/shen/Downloads/google-cloud-sdk/completion.zsh.inc'; fi
export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"