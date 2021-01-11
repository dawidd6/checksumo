# checksumo 

<img src="./data/com.github.dawidd6.checksumo.svg" width="128" alt="logo">

A simple application for verifying specified file against given hash, written in Go with GTK+3 graphical interface.

Automatically detects the following hash types:
- SHA-256
- SHA-512
- MD5

Supports cancellation of verification if desired.

## Gallery

![](data/screenshots/check1.png)

![](data/screenshots/check2.png)

![](data/screenshots/check3.png)

![](data/screenshots/check4.png)

## Installation

This application targets Ubuntu 18.04 and later, though the only requirement is to have GTK+ 3.22 or later installed.

### Via apt repository

An apt repository is provided. It is made out of GitHub Releases and packages are served via HTTPS protocol, but they are
not signed.

```shell script
echo "deb [trusted=yes] https://github.com/dawidd6/checksumo/releases/download/repo ./" | sudo tee -a /etc/apt/sources.list
sudo apt update
sudo apt install checksumo
```

### From released binary

Head over to Releases section of this repository, grab the `.deb` file and install it via `dpkg`.

### From source

Needs `libgtk-3-dev` package on Debian-based distributions.

```shell script
make
make install
```