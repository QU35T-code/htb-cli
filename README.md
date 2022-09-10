```
This is a program developed in Go to facilitate and automate certain tasks for the Hackthebox platform.

Usage:
  htb [command]

Available Commands:
  blood       Displays users who have blood the machine
  config      Save the machine chosen as an argument
  flag        Submit a flag (user and root)
  help        Help about any command
  info        Displays machine information
  ip          Get machine IP address
  reset       Reset a machine
  start       Start a machine
  status      Displays the active machine
  stop        Stop the current machine
  www         Starts an HTTP (WIP) server

Flags:
  -h, --help   help for htb

Use "htb [command] --help" for more information about a command.
```

## Installation

`git clone https://github.com/QU35T-code/HTB-manager.git`

## Configuration

Add `config.yml` file at root project with this sample :

```
token: eyJ...
```
