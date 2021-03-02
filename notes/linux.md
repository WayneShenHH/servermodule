# Commads

config key binding: alt + left/right = home/end keys

```s
echo "keycode  64 = Mode_switch Meta_L Alt_L Meta_L\nkeycode 113 = Left NoSymbol Home\nkeycode 114 = Right NoSymbol End" >> ~/.Xmodmap

echo "\nxmodmap ~/.Xmodmap" >> ~/.zshrc
```

login with root
>sudo su -

postman
>sudo snap install postman

vscode
>sudo snap install --classic code

directory gui
>nautilus your-path-here

quit process on terminal
>ctrl + shift + c

clear terminal
>ctrl + l

open the port from firewall
>sudo ufw allow ssh

install git
>sudo apt install git

install oh-my-zsh
```
sudo apt-get install zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
echo  "\n# default set zsh\n/usr/bin/zsh" >> ~/.bashrc
```

install java jdk
>sudo apt install openjdk-8-jdk
>java -version

install node.js
>sudo apt-get install -y nodejs
>node -v

install npm
>sudo apt-get install npm
>npm -v

install aglio generator
>sudo npm install -g aglio

install aglio mock server
>sudo npm install -g api-mock

ubuntu kernel info
>dpkg --get-selections | grep linux-image

ubuntu remove kernel version
>sudo apt-get purge linux-image-2.6.38-10-generic

install python alternatives

```shell
# preview all versions of python
ls /usr/bin/python*

# should error due to there has not any version
sudo update-alternatives --list python

# add python2 & python3 to alternatives
sudo update-alternatives --install /usr/bin/python python /usr/bin/python2 1
sudo update-alternatives --install /usr/bin/python python /usr/bin/python3 2

# choose your default python here
sudo update-alternatives --config python

# check result
python --version
```

kernel version manager

```
sudo add-apt-repository ppa:cappelikan/ppa
sudo apt update
sudo apt install mainline
```

nvidia driver

```shell
sudo add-apt-repository ppa:graphics-drivers/ppa
sudo apt update
sudo apt install ubuntu-drivers-common

ubuntu-drivers devices # list versions
sudo apt install nvidia-driver-450
sudo reboot
lsmod|grep nvidia # check result
```

chewing input
```
sudo apt install ibus-chewing
```