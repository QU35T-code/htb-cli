#!/bin/bash

ifconfig tun0 | grep 'inet ' | cut -d' ' -f10 | tr -d '\n'