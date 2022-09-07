#!/usr/bin/env python3

import requests
import sys
import configparser

config = configparser.ConfigParser()
config.read('config.ini')

def reset_machine():
    url = "https://www.hackthebox.com/api/v4/vm/reset"
    machine_id = config.get("HTB", "machine_id")
    jwt_token = config.get("HTB", "token")

    headers = {
        "User-Agent": "HTBTool",
        "Authorization": f"Bearer {jwt_token}"
    }

    r = requests.post(url, headers=headers, json={
        "machine_id": machine_id
    })
    print(r.json()["message"])

if __name__ == '__main__':
    reset_machine()