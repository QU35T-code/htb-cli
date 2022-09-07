#!/bin/bash

echo "WIP"
#WIP ffuf -c -w /opt/SecLists/Discovery/DNS/subdomains-top1million-110000.txt -u http://$1 -H "Host: FUZZ.$2"