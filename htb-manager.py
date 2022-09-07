#!/usr/bin/env python3

import argparse
import os

parser = argparse.ArgumentParser()

dirname = "/opt/Setup-HTB/shortcuts"

htb_group = parser.add_argument_group("Hackthebox")
utils_group = parser.add_argument_group("Utils")

utils_group.add_argument("-ci", help="Copy Target IP Address", action='store_true')
utils_group.add_argument("-cli", help="Copy Local IP Address", action='store_true')
utils_group.add_argument("-www", help="Run HTTP Server (default port 80)", const='80', nargs='?')
utils_group.add_argument("-vhosts", help="WIP", action='store_true')
utils_group.add_argument("-hosts", nargs='?', help="Add a hostname")
utils_group.add_argument("-vpn", const='lab', nargs='?', choices=["lab", "release", "starting"], help="Start the VPN (default lab)")
htb_group.add_argument("-start", nargs='?', help="Start a machine (WIP)")
htb_group.add_argument("-stop", help="Stop current machine", action='store_true')
htb_group.add_argument("-reset", nargs='?', help="Reset a machine (WIP)")
htb_group.add_argument("-flag", nargs='*', help="Submit a flag")
utils_group.add_argument("-prev", help="Generate a powershell reverse shell (Base64 encoded) (WIP)")

args = parser.parse_args()

if args.ci:
    os.system(f"{dirname}/copy-ip.sh")
elif args.cli:
    os.system(f"{dirname}/copy-local-ip.sh")
elif args.www:
    os.system(f"{dirname}/webserver.sh {args.www}")
elif args.vhosts:
    os.system(f"{dirname}/vhosts.sh")
elif args.hosts:
    os.system(f"{dirname}/hosts.sh {args.hosts}")
elif args.vpn:
    os.system(f"{dirname}/vpn.sh {args.vpn}")
elif args.flag:
    os.system(f"{dirname}/submit-flag.py {args.flag[0]} {args.flag[1]}")
elif args.stop:
    os.system(f"{dirname}/stop.py")