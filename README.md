<img align="left" width="64" height="64" src="data/com.github.dawidd6.checksumo.svg">
<h1>Checksumo</h1>

<a href='https://flathub.org/apps/details/com.github.dawidd6.checksumo'><img width='240' alt='Download on Flathub' src='https://flathub.org/assets/badges/flathub-badge-en.png'/></a>

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

![](data/screenshots/check5.png)

## Installation

### Flatpak

Click the Flathub badge above for more information.

### Source

To build and run this software, one needs to have:

- **GTK** `>=3.24`
- **Go** `>=1.11`
- **Meson** `>=0.48`

then execute below commands:

```shell script
meson build
sudo meson install -C build
```
