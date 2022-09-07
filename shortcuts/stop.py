#!/usr/bin/env python3

import requests
import sys
import configparser

config = configparser.ConfigParser()
config.read('config.ini')

def stop_machine():
    url = "https://www.hackthebox.com/api/v4/machine/stop"
    jwt_token = config.get("HTB", "token")

    headers = {
        "User-Agent": "HTBTool",
        "Authorization": f"Bearer {jwt_token}"
    }

    r = requests.post(url, headers=headers)
    print(r.json()["message"])

if __name__ == '__main__':
    stop_machine()