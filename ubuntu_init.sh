sudo snap install slack
sudo snap install --classic code
sudo apt-get install build-essential libssl-dev
sudo apt-get install git curl zsh docker.io htop

# oh-my-zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
echo  "\n# default set zsh\n/usr/bin/zsh" >> ~/.bashrc

# docker 
sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker

# https://github.com/nvm-sh/nvm/releases
curl https://raw.githubusercontent.com/creationix/nvm/v0.40.3/install.sh | bash
nvm ls-remote
nvm install v25.5.0


# git
git config --global credential.helper store

ssh-keygen -t rsa -C "wayne_shen@tengyuntech.com"

# golang
# https://go.dev/dl/

mkdir ~/projects
mkdir ~/projects/fgw
mkdir ~/projects/fgw/deploy
mkdir ~/projects/fgw/web
mkdir ~/projects/fgw/cocos
cd ~/projects
git clone https://github.com/WayneShenHH/servermodule.git

# hyper js
cd ~/projects/servermodule
sudo dpkg -i ./hyper_3.4.1_amd.deb
# hyper issue
sudo chmod 4755 /opt/Hyper/chrome-sandbox

# sync zshrc
rm ~/.zshrc
cp .zshrc ~/.zshrc
