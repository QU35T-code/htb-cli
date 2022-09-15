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

Flags:
  -h, --help   help for htb

Use "htb [command] --help" for more information about a command.
```

## Installation

`go install https://github.com/QU35T-code/HTB-pentest`

## Configuration

Add `config.yml` file at root project with this sample :
API Token can be find here : https://app.hackthebox.com/profile/settings => Create App Token

```
token: eyJ...
```

## Config
This step is important and allows you to identify the target machine for other commands.
```
└─$ ./htb config Active
The machine is correctly configured
```

## Blood

By default the command shows the active machine.
```
└─$ ./htb blood        
Machine : Vessel

--- USER ---
Name : Coaran
Time : 3H 23M 16S

--- ROOT ---
Name : irogir
Time : 4H 20M 10S


└─$ ./htb blood Opensource
Machine : OpenSource

--- USER ---
Name : jazzpizazz
Time : 0H 57M 3S

--- ROOT ---
Name : jazzpizazz
Time : 1H 39M 17S
```

## Flag

```
└─$ ./htb flag flag4testing 3

Machine : SteamCloud

SteamCloud user is now owned.
```

## Info

By default the command shows the active machine.
```
└─$ ./htb info

--- INFO ---
Name : Vessel
OS : Linux
Points : 40
Difficulty : Hard
Is Completed ? false


└─$ ./htb info Opensource
--- INFO ---
Name : OpenSource
OS : Linux
Points : 20
Difficulty : EasyIs Completed ? true  
```

## IP

```
└─$ ./htb ip
Machine : Vessel

10.10.11.178  
```

## Start

```
└─$ ./htb start
Retired : Machine deployed to lab.
```

## Stop

```
└─$ ./htb stop
Retired : Machine terminated.
```
