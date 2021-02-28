Disable fast startup on Windows 10

Disable Secure Boot in bios

Boot Ubuntu from USB

Press 'e' in grub

Append nomodeset modprobe.blacklist=nouveau after quiet splash

Press F10 to boot

Install Ubuntu / Mint

Create `/etc/modprobe.d/blacklist-nvidia-nouveau.conf`:
```
blacklist nouveau
options nouveau modeset=0       
```

Reboot

Install kernel 5.11.0-051100-generic or higher from mainline builds:
```
sudo add-apt-repository ppa:cappelikan/ppa
sudo apt update
sudo apt install mainline
```

Reboot

Install latest nvidia-driver-460 from ppa
```shell
sudo add-apt-repository ppa:graphics-drivers/ppa
sudo apt update
sudo apt install ubuntu-drivers-common

ubuntu-drivers devices # list versions
sudo apt install nvidia-driver-450
sudo reboot
lsmod|grep nvidia # check result
```

Reboot

Add option in `/usr/share/X11/xorg.conf.d/10-nvidia.conf`
```
Section "OutputClass"
    Identifier "nvidia"
    MatchDriver "nvidia-drm"
    Driver "nvidia"
    Option "AllowEmptyInitialConfiguration"
    Option "PrimaryGPU" "yes"
    ModulePath "/usr/lib/x86_64-linux-gnu/nvidia/xorg"
EndSection
```

Add option in `/usr/share/X11/xorg.conf.d/10-amdgpu.conf`
```
Section "OutputClass"
	Identifier "AMDgpu"
	MatchDriver "amdgpu"
	Driver "amdgpu"
	Option "PrimaryGPU" "no"
EndSection
```

Edit `/etc/default/grub`
```
GRUB_CMDLINE_LINUX_DEFAULT="quiet splash amdgpu.exp_hw_support=1 modprobe.blacklist=nouveau"
```

Run update-grub
```
sudo update-grub
```

