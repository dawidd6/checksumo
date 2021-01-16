# checksumo 

<img src="./data/checksumo.svg" width="128" alt="logo">

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

## Requirements

- **GTK** `>=3.22`
- **Go** `>=1.11`

## Building

```shell script
make
sudo make install
```
