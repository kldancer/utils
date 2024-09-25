filter
```shell
[root@server4 src]# iptables -t filter -S 
# Warning: iptables-legacy tables present, use iptables-legacy to see them
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-N DOCKER
-N DOCKER-ISOLATION-STAGE-1
-N DOCKER-ISOLATION-STAGE-2
-N DOCKER-USER
-N KUBE-FIREWALL
-N KUBE-KUBELET-CANARY
-A INPUT -j KUBE-FIREWALL
-A FORWARD -j DOCKER-USER
-A FORWARD -j DOCKER-ISOLATION-STAGE-1
-A FORWARD -o docker0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A FORWARD -o docker0 -j DOCKER
-A FORWARD -i docker0 ! -o docker0 -j ACCEPT
-A FORWARD -i docker0 -o docker0 -j ACCEPT
-A OUTPUT -j KUBE-FIREWALL
-A DOCKER-ISOLATION-STAGE-1 -i docker0 ! -o docker0 -j DOCKER-ISOLATION-STAGE-2
-A DOCKER-ISOLATION-STAGE-1 -j RETURN
-A DOCKER-ISOLATION-STAGE-2 -o docker0 -j DROP
-A DOCKER-ISOLATION-STAGE-2 -j RETURN
-A DOCKER-USER -j RETURN
-A KUBE-FIREWALL -m comment --comment "kubernetes firewall for dropping marked packets" -m mark --mark 0x8000/0x8000 -j DROP
-A KUBE-FIREWALL ! -s 127.0.0.0/8 -d 127.0.0.0/8 -m comment --comment "block incoming localnet connections" -m conntrack ! --ctstate RELATED,ESTABLISHED,DNAT -j DROP
```

nat
```shell
[root@server4 src]# iptables -t nat -S 
# Warning: iptables-legacy tables present, use iptables-legacy to see them
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT
-N DOCKER
-N KUBE-KUBELET-CANARY
-N KUBE-MARK-DROP
-N KUBE-MARK-MASQ
-N KUBE-POSTROUTING
-A PREROUTING -m addrtype --dst-type LOCAL -j DOCKER
-A OUTPUT ! -d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j DOCKER
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A POSTROUTING -s 172.17.0.0/16 ! -o docker0 -j MASQUERADE
-A DOCKER -i docker0 -j RETURN
-A KUBE-MARK-DROP -j MARK --set-xmark 0x8000/0x8000
-A KUBE-MARK-MASQ -j MARK --set-xmark 0x4000/0x4000
-A KUBE-POSTROUTING -m mark ! --mark 0x4000/0x4000 -j RETURN
-A KUBE-POSTROUTING -j MARK --set-xmark 0x4000/0x0
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -j MASQUERADE --random-fully
```

mangle、raw
```shell
[root@server4 src]# iptables -t mangle -S
# Warning: iptables-legacy tables present, use iptables-legacy to see them
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT
-N KUBE-KUBELET-CANARY
[root@server4 src]# iptables -t raw -S
# Warning: iptables-legacy tables present, use iptables-legacy to see them
-P PREROUTING ACCEPT
-P OUTPUT ACCEPT
```

nat
NodePortSVC
1. 对目标地址是本地的流量跳转到FAB-NODEPORTS
2. 将目的端口为31908的TCP流量跳转到FAB-SVC-2-TCP9090链进行处理。
3. FAB-SVC-2-TCP9090链跳转到FAB-SEP-2-0-TCP9090链
4. 在FAB-SEP-2-0-TCP9090链中添加一条规则，将流量跳转到FAB-MARK-MASQ链进行处理
5. 在FAB-SEP-2-0-TCP9090链中添加一条规则，将目的地址和端口DNAT为10.250.2.85:9090，并且是TCP协议的流量。
6. 在FAB-MARK-MASQ链中添加一条规则，将数据包标记为0x40/0x40，用于后续的网络地址转换。
```shell
-A PREROUTING -m comment --comment "fabric service nodeports" -m addrtype --dst-type LOCAL -j FAB-NODEPORTS
-A OUTPUT -m comment --comment "fabric service nodeports" -m addrtype --dst-type LOCAL -j FAB-NODEPORTS

-N FAB-SVC-2-TCP9090
-N FAB-SEP-2-0-TCP9090

-A FAB-NODEPORTS -p tcp -m tcp --dport 31908 -j FAB-SVC-2-TCP9090
-A FAB-SVC-2-TCP9090 -j FAB-SEP-2-0-TCP9090
-A FAB-SEP-2-0-TCP9090 -m comment --comment "fabric masq endpoint 10.250.2.85:9090" -j FAB-MARK-MASQ
-A FAB-SEP-2-0-TCP9090 -p tcp -m comment --comment "fabric dnat endpoint 10.250.2.85:9090" -j DNAT --to-destination 10.250.2.85:9090

-N FAB-MARK-MASQ
-A FAB-MARK-MASQ -m comment --comment "fabric mark packet to masquerade" -j MARK --set-xmark 0x40/0x40
```
