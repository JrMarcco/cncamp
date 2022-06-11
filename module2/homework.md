```
cd cncamp/module2/

docker login
make push 

docker images
# REPOSITORY     TAG                IMAGE ID       CREATED         SIZE
# jrmarcco/jrx   simple-web-1.0.0   7e42157825fb   5 minutes ago   84.3MB
# ubuntu         latest             27941809078c   4 days ago      77.8MB
# nginx          latest             0e901e68141f   2 weeks ago     142MB

docker run -d 7e42157825fb

docker ps 
# CONTAINER ID   IMAGE          COMMAND                  CREATED          STATUS          PORTS     NAMES
# 8de1d4455563   7e42157825fb   "/bin/sh -c /simple-…"   17 seconds ago   Up 16 seconds             xenodochial_chaum

docker inspect 8de1d4455563|grep Pid
#             "Pid": 8406,
#             "PidMode": "",
#             "PidsLimit": null,

nsenter -t 8406 -n ip a
# 1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
#     link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
#     inet 127.0.0.1/8 scope host lo
#        valid_lft forever preferred_lft forever
# 10: eth0@if11: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
#     link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
#     inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
#        valid_lft forever preferred_lft forever

curl http://172.17.0.2:8080/healthz 
# activating

docker stop 8de1d4455563
```