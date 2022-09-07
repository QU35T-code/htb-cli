#!/usr/bin/env python3

import requests
import sys
import configparser

config = configparser.ConfigParser()
config.read('config.ini')

def get_ip():
    machine_id = config.get("HTB", "machine_id")
    url = f"https://www.hackthebox.com/api/v4/machine/profile/{machine_id}"
    jwt_token = config.get("HTB", "token")

    headers = {
        "User-Agent": "HTBTool",
        "Authorization": f"Bearer {jwt_token}"
    }

    r = requests.get(url, headers=headers)
    print(r.json()["info"]["ip"])

if __name__ == '__main__':
    get_ip()