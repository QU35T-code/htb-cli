#!/usr/bin/env python3

import requests
import sys
import configparser

config = configparser.ConfigParser()
config.read('config.ini')

def submit_flag():
    url = "https://www.hackthebox.com/api/v4/machine/own"
    flag = sys.argv[1]
    machine_id = config.get("HTB", "machine_id")
    difficulty = int(sys.argv[2]) * 10
    jwt_token = config.get("HTB", "token")

    headers = {
        "User-Agent": "HTBTool",
        "Authorization": f"Bearer {jwt_token}"
    }

    r = requests.post(url, headers=headers, json={
        "flag": flag,
        "id": machine_id,
        "difficulty": difficulty
    })
    print(r.json()["message"])

if __name__ == '__main__':
    if len(sys.argv) != 3:
        print("Parameters : [FLAG] [DIFFICULTY]")
    else:
        submit_flag()