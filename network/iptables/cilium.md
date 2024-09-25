filter
```shell
[root@f1 ~]# iptables -t filter -S
-P INPUT ACCEPT
-P FORWARD DROP
-P OUTPUT ACCEPT
-N CILIUM_FORWARD
-N CILIUM_INPUT
-N CILIUM_OUTPUT
-N DOCKER
-N DOCKER-ISOLATION-STAGE-1
-N DOCKER-ISOLATION-STAGE-2
-N DOCKER-USER
-N KUBE-EXTERNAL-SERVICES
-N KUBE-FIREWALL
-N KUBE-FORWARD
-N KUBE-KUBELET-CANARY
-N KUBE-NODEPORTS
-N KUBE-PROXY-CANARY
-N KUBE-SERVICES
-A INPUT -m comment --comment "cilium-feeder: CILIUM_INPUT" -j CILIUM_INPUT
-A INPUT -m comment --comment "kubernetes health check service ports" -j KUBE-NODEPORTS
-A INPUT -m conntrack --ctstate NEW -m comment --comment "kubernetes externally-visible service portals" -j KUBE-EXTERNAL-SERVICES
-A INPUT -j KUBE-FIREWALL
-A FORWARD -m comment --comment "cilium-feeder: CILIUM_FORWARD" -j CILIUM_FORWARD
-A FORWARD -j DOCKER-USER
-A FORWARD -j DOCKER-ISOLATION-STAGE-1
-A FORWARD -o docker0 -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
-A FORWARD -o docker0 -j DOCKER
-A FORWARD -i docker0 ! -o docker0 -j ACCEPT
-A FORWARD -i docker0 -o docker0 -j ACCEPT
-A FORWARD -m comment --comment "kubernetes forwarding rules" -j KUBE-FORWARD
-A FORWARD -m conntrack --ctstate NEW -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A FORWARD -m conntrack --ctstate NEW -m comment --comment "kubernetes externally-visible service portals" -j KUBE-EXTERNAL-SERVICES
-A OUTPUT -m comment --comment "cilium-feeder: CILIUM_OUTPUT" -j CILIUM_OUTPUT
-A OUTPUT -m conntrack --ctstate NEW -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A OUTPUT -j KUBE-FIREWALL
-A CILIUM_FORWARD -o cilium_host -m comment --comment "cilium: any->cluster on cilium_host forward accept" -j ACCEPT
-A CILIUM_FORWARD -i cilium_host -m comment --comment "cilium: cluster->any on cilium_host forward accept (nodeport)" -j ACCEPT
-A CILIUM_FORWARD -i lxc+ -m comment --comment "cilium: cluster->any on lxc+ forward accept" -j ACCEPT
-A CILIUM_FORWARD -i cilium_net -m comment --comment "cilium: cluster->any on cilium_net forward accept (nodeport)" -j ACCEPT
-A CILIUM_INPUT -m mark --mark 0x200/0xf00 -m comment --comment "cilium: ACCEPT for proxy traffic" -j ACCEPT
-A CILIUM_OUTPUT -m mark --mark 0xa00/0xe00 -m comment --comment "cilium: ACCEPT for proxy traffic" -j ACCEPT
-A CILIUM_OUTPUT -m mark --mark 0x800/0xe00 -m comment --comment "cilium: ACCEPT for l7 proxy upstream traffic" -j ACCEPT
-A CILIUM_OUTPUT -m mark ! --mark 0xe00/0xf00 -m mark ! --mark 0xd00/0xf00 -m mark ! --mark 0x400/0xf00 -m mark ! --mark 0xa00/0xe00 -m mark ! --mark 0x800/0xe00 -m mark ! --mark 0xf00/0xf00 -m comment --comment "cilium: host->any mark as from host" -j MARK --set-xmark 0xc00/0xf00
-A DOCKER-ISOLATION-STAGE-1 -i docker0 ! -o docker0 -j DOCKER-ISOLATION-STAGE-2
-A DOCKER-ISOLATION-STAGE-1 -j RETURN
-A DOCKER-ISOLATION-STAGE-2 -o docker0 -j DROP
-A DOCKER-ISOLATION-STAGE-2 -j RETURN
-A DOCKER-USER -j RETURN
-A KUBE-FIREWALL -m comment --comment "kubernetes firewall for dropping marked packets" -m mark --mark 0x8000/0x8000 -j DROP
-A KUBE-FIREWALL ! -s 127.0.0.0/8 -d 127.0.0.0/8 -m comment --comment "block incoming localnet connections" -m conntrack ! --ctstate RELATED,ESTABLISHED,DNAT -j DROP
-A KUBE-FORWARD -m conntrack --ctstate INVALID -j DROP
-A KUBE-FORWARD -m comment --comment "kubernetes forwarding rules" -m mark --mark 0x4000/0x4000 -j ACCEPT
-A KUBE-FORWARD -m comment --comment "kubernetes forwarding conntrack rule" -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
```

```shell
# 创建新的filter链
-N CILIUM_FORWARD
-N CILIUM_INPUT
-N CILIUM_OUTPUT
# 在INPUT链中添加新规则，将匹配该规则的数据包交给CILIUM_INPUT链处理
-A INPUT -m comment --comment "cilium-feeder: CILIUM_INPUT" -j CILIUM_INPUT
-A FORWARD -m comment --comment "cilium-feeder: CILIUM_FORWARD" -j CILIUM_FORWARD
-A OUTPUT -m comment --comment "cilium-feeder: CILIUM_OUTPUT" -j CILIUM_OUTPUT
# 在CILIUM_FORWARD链中添加新规则:
# 允许从任何接口到cilium_host接口的转发，并接受这些数据包
# 允许从cilium_host接口到任何接口的转发，并接受这些数据包
# 允许从lxc+接口到任何接口的转发，并接受这些数据包
# 允许从cilium_net接口到任何接口的转发，并接受这些数据包
-A CILIUM_FORWARD -o cilium_host -m comment --comment "cilium: any->cluster on cilium_host forward accept" -j ACCEPT
-A CILIUM_FORWARD -i cilium_host -m comment --comment "cilium: cluster->any on cilium_host forward accept (nodeport)" -j ACCEPT
-A CILIUM_FORWARD -i lxc+ -m comment --comment "cilium: cluster->any on lxc+ forward accept" -j ACCEPT
-A CILIUM_FORWARD -i cilium_net -m comment --comment "cilium: cluster->any on cilium_net forward accept (nodeport)" -j ACCEPT
# 在CILIUM_INPUT链中添加新规则:
# 允许标记为0x200/0xf00的数据包通过，并接受这些数据包
-A CILIUM_INPUT -m mark --mark 0x200/0xf00 -m comment --comment "cilium: ACCEPT for proxy traffic" -j ACCEPT
# 在CILIUM_OUTPUT链中添加新规则
# 允许标记为0xa00/0xe00的数据包通过，并接受这些数据包
# 允许标记为0x800/0xe00的数据包通过，并接受这些数据包
# 将不匹配任何之前规则的数据包标记为0xc00/0xf00。
-A CILIUM_OUTPUT -m mark --mark 0xa00/0xe00 -m comment --comment "cilium: ACCEPT for proxy traffic" -j ACCEPT
-A CILIUM_OUTPUT -m mark --mark 0x800/0xe00 -m comment --comment "cilium: ACCEPT for l7 proxy upstream traffic" -j ACCEPT
-A CILIUM_OUTPUT -m mark ! --mark 0xe00/0xf00 -m mark ! --mark 0xd00/0xf00 -m mark ! --mark 0x400/0xf00 -m mark ! --mark 0xa00/0xe00 -m mark ! --mark 0x800/0xe00 -m mark ! --mark 0xf00/0xf00 -m comment --comment "cilium: host->any mark as from host" -j MARK --set-xmark 0xc00/0xf00
```


nat
```shell
[root@f1 ~]# iptables -t nat -S 
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT
-N CILIUM_OUTPUT_nat
-N CILIUM_POST_nat
-N CILIUM_PRE_nat
-N DOCKER
-N KUBE-KUBELET-CANARY
-N KUBE-MARK-DROP
-N KUBE-MARK-MASQ
-N KUBE-NODEPORTS
-N KUBE-POSTROUTING
-N KUBE-PROXY-CANARY
-N KUBE-SEP-4NJI4U4OVD7ASSTN
-N KUBE-SEP-4W4DMGANR5QBPYNW
-N KUBE-SEP-FUY4JREFWM7DZPDS
-N KUBE-SEP-IMPQAXS4DSXOT6OW
-N KUBE-SEP-JFY262ZQRD3MRSBY
-N KUBE-SEP-JTHDLAECD3YWPOED
-N KUBE-SEP-LNM6HHUZOJM3CKW7
-N KUBE-SEP-M5CX275RXKVJC7F7
-N KUBE-SEP-MV2OXEWOCKPWUWNN
-N KUBE-SEP-N27GNE72QNUULDFK
-N KUBE-SEP-Q5DFUOU55OKTF7SL
-N KUBE-SEP-SZOBJIZ4TCVR33ZB
-N KUBE-SEP-Y5LSJLRXX3OJOFLV
-N KUBE-SERVICES
-N KUBE-SVC-DZSQUIY5BXN7HSA3
-N KUBE-SVC-ERIFXISQEP7F7OF4
-N KUBE-SVC-JD5MR3NA4I4DYORP
-N KUBE-SVC-NPX46M4PTMTKRN6Y
-N KUBE-SVC-NZTS37XDTDNXGCKJ
-N KUBE-SVC-PDS5MBI2OEFEKTY7
-N KUBE-SVC-TCOU7JCQXEZGVUNU
-N KUBE-SVC-XXQFMBX2TZG6LPU6
-N KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A PREROUTING -m comment --comment "cilium-feeder: CILIUM_PRE_nat" -j CILIUM_PRE_nat
-A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A PREROUTING -m addrtype --dst-type LOCAL -j DOCKER
-A OUTPUT -m comment --comment "cilium-feeder: CILIUM_OUTPUT_nat" -j CILIUM_OUTPUT_nat
-A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A OUTPUT ! -d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j DOCKER
-A POSTROUTING -m comment --comment "cilium-feeder: CILIUM_POST_nat" -j CILIUM_POST_nat
-A POSTROUTING -s 172.17.0.0/16 ! -o docker0 -j MASQUERADE
-A POSTROUTING -m comment --comment "kubernetes postrouting rules" -j KUBE-POSTROUTING
-A CILIUM_POST_nat -s 10.0.0.0/24 ! -d 10.0.0.0/24 ! -o cilium_+ -m comment --comment "cilium masquerade non-cluster" -j MASQUERADE
-A CILIUM_POST_nat -m mark --mark 0xa00/0xe00 -m comment --comment "exclude proxy return traffic from masquerade" -j ACCEPT
-A CILIUM_POST_nat ! -s 10.0.0.0/24 ! -d 10.0.0.0/24 -o cilium_host -m comment --comment "cilium host->cluster masquerade" -j SNAT --to-source 10.0.0.96
-A CILIUM_POST_nat -s 127.0.0.1/32 -o cilium_host -m comment --comment "cilium host->cluster from 127.0.0.1 masquerade" -j SNAT --to-source 10.0.0.96
-A CILIUM_POST_nat -o cilium_host -m mark --mark 0xf00/0xf00 -m conntrack --ctstate DNAT -m comment --comment "hairpin traffic that originated from a local pod" -j SNAT --to-source 10.0.0.96
-A DOCKER -i docker0 -j RETURN
-A KUBE-MARK-DROP -j MARK --set-xmark 0x8000/0x8000
-A KUBE-MARK-MASQ -j MARK --set-xmark 0x4000/0x4000
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp --dport 30080 -j KUBE-SVC-XXQFMBX2TZG6LPU6
-A KUBE-NODEPORTS -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp --dport 30486 -j KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A KUBE-POSTROUTING -m mark ! --mark 0x4000/0x4000 -j RETURN
-A KUBE-POSTROUTING -j MARK --set-xmark 0x4000/0x0
-A KUBE-POSTROUTING -m comment --comment "kubernetes service traffic requiring SNAT" -j MASQUERADE --random-fully
-A KUBE-SEP-4NJI4U4OVD7ASSTN -s 10.0.0.52/32 -m comment --comment "default/x-tools-svc-1" -j KUBE-MARK-MASQ
-A KUBE-SEP-4NJI4U4OVD7ASSTN -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp -j DNAT --to-destination 10.0.0.52:80
-A KUBE-SEP-4W4DMGANR5QBPYNW -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-MARK-MASQ
-A KUBE-SEP-4W4DMGANR5QBPYNW -p tcp -m comment --comment "kube-system/kube-dns:metrics" -m tcp -j DNAT --to-destination 10.0.0.61:9153
-A KUBE-SEP-FUY4JREFWM7DZPDS -s 10.0.0.39/32 -m comment --comment "default/deathstar" -j KUBE-MARK-MASQ
-A KUBE-SEP-FUY4JREFWM7DZPDS -p tcp -m comment --comment "default/deathstar" -m tcp -j DNAT --to-destination 10.0.0.39:80
-A KUBE-SEP-IMPQAXS4DSXOT6OW -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-MARK-MASQ
-A KUBE-SEP-IMPQAXS4DSXOT6OW -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp" -m tcp -j DNAT --to-destination 10.0.0.61:53
-A KUBE-SEP-JFY262ZQRD3MRSBY -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:dns" -j KUBE-MARK-MASQ
-A KUBE-SEP-JFY262ZQRD3MRSBY -p udp -m comment --comment "kube-system/kube-dns:dns" -m udp -j DNAT --to-destination 10.0.0.61:53
-A KUBE-SEP-JTHDLAECD3YWPOED -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:dns" -j KUBE-MARK-MASQ
-A KUBE-SEP-JTHDLAECD3YWPOED -p udp -m comment --comment "kube-system/kube-dns:dns" -m udp -j DNAT --to-destination 10.0.1.248:53
-A KUBE-SEP-LNM6HHUZOJM3CKW7 -s 10.0.1.227/32 -m comment --comment "kube-system/hubble-ui:http" -j KUBE-MARK-MASQ
-A KUBE-SEP-LNM6HHUZOJM3CKW7 -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp -j DNAT --to-destination 10.0.1.227:8081
-A KUBE-SEP-M5CX275RXKVJC7F7 -s 10.0.1.168/32 -m comment --comment "default/deathstar" -j KUBE-MARK-MASQ
-A KUBE-SEP-M5CX275RXKVJC7F7 -p tcp -m comment --comment "default/deathstar" -m tcp -j DNAT --to-destination 10.0.1.168:80
-A KUBE-SEP-MV2OXEWOCKPWUWNN -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-MARK-MASQ
-A KUBE-SEP-MV2OXEWOCKPWUWNN -p tcp -m comment --comment "kube-system/kube-dns:metrics" -m tcp -j DNAT --to-destination 10.0.1.248:9153
-A KUBE-SEP-N27GNE72QNUULDFK -s 172.28.8.138/32 -m comment --comment "kube-system/hubble-peer:peer-service" -j KUBE-MARK-MASQ
-A KUBE-SEP-N27GNE72QNUULDFK -p tcp -m comment --comment "kube-system/hubble-peer:peer-service" -m tcp -j DNAT --to-destination 172.28.8.138:4244
-A KUBE-SEP-Q5DFUOU55OKTF7SL -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-MARK-MASQ
-A KUBE-SEP-Q5DFUOU55OKTF7SL -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp" -m tcp -j DNAT --to-destination 10.0.1.248:53
-A KUBE-SEP-SZOBJIZ4TCVR33ZB -s 10.0.1.32/32 -m comment --comment "kube-system/hubble-relay" -j KUBE-MARK-MASQ
-A KUBE-SEP-SZOBJIZ4TCVR33ZB -p tcp -m comment --comment "kube-system/hubble-relay" -m tcp -j DNAT --to-destination 10.0.1.32:4245
-A KUBE-SEP-Y5LSJLRXX3OJOFLV -s 172.28.8.138/32 -m comment --comment "default/kubernetes:https" -j KUBE-MARK-MASQ
-A KUBE-SEP-Y5LSJLRXX3OJOFLV -p tcp -m comment --comment "default/kubernetes:https" -m tcp -j DNAT --to-destination 172.28.8.138:6443
-A KUBE-SERVICES -d 10.106.170.104/32 -p tcp -m comment --comment "kube-system/hubble-peer:peer-service cluster IP" -m tcp --dport 443 -j KUBE-SVC-NZTS37XDTDNXGCKJ
-A KUBE-SERVICES -d 10.104.103.145/32 -p tcp -m comment --comment "kube-system/hubble-relay cluster IP" -m tcp --dport 80 -j KUBE-SVC-DZSQUIY5BXN7HSA3
-A KUBE-SERVICES -d 10.105.169.103/32 -p tcp -m comment --comment "default/x-tools-svc-1 cluster IP" -m tcp --dport 80 -j KUBE-SVC-XXQFMBX2TZG6LPU6
-A KUBE-SERVICES -d 10.96.0.10/32 -p udp -m comment --comment "kube-system/kube-dns:dns cluster IP" -m udp --dport 53 -j KUBE-SVC-TCOU7JCQXEZGVUNU
-A KUBE-SERVICES -d 10.96.0.10/32 -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp cluster IP" -m tcp --dport 53 -j KUBE-SVC-ERIFXISQEP7F7OF4
-A KUBE-SERVICES -d 10.96.0.10/32 -p tcp -m comment --comment "kube-system/kube-dns:metrics cluster IP" -m tcp --dport 9153 -j KUBE-SVC-JD5MR3NA4I4DYORP
-A KUBE-SERVICES -d 10.107.246.98/32 -p tcp -m comment --comment "default/deathstar cluster IP" -m tcp --dport 80 -j KUBE-SVC-PDS5MBI2OEFEKTY7
-A KUBE-SERVICES -d 10.96.0.1/32 -p tcp -m comment --comment "default/kubernetes:https cluster IP" -m tcp --dport 443 -j KUBE-SVC-NPX46M4PTMTKRN6Y
-A KUBE-SERVICES -d 10.99.53.195/32 -p tcp -m comment --comment "kube-system/hubble-ui:http cluster IP" -m tcp --dport 80 -j KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A KUBE-SERVICES -m comment --comment "kubernetes service nodeports; NOTE: this must be the last rule in this chain" -m addrtype --dst-type LOCAL -j KUBE-NODEPORTS
-A KUBE-SVC-DZSQUIY5BXN7HSA3 -m comment --comment "kube-system/hubble-relay" -j KUBE-SEP-SZOBJIZ4TCVR33ZB
-A KUBE-SVC-ERIFXISQEP7F7OF4 -m comment --comment "kube-system/kube-dns:dns-tcp" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-IMPQAXS4DSXOT6OW
-A KUBE-SVC-ERIFXISQEP7F7OF4 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-SEP-Q5DFUOU55OKTF7SL
-A KUBE-SVC-JD5MR3NA4I4DYORP -m comment --comment "kube-system/kube-dns:metrics" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-4W4DMGANR5QBPYNW
-A KUBE-SVC-JD5MR3NA4I4DYORP -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-SEP-MV2OXEWOCKPWUWNN
-A KUBE-SVC-NPX46M4PTMTKRN6Y -m comment --comment "default/kubernetes:https" -j KUBE-SEP-Y5LSJLRXX3OJOFLV
-A KUBE-SVC-NZTS37XDTDNXGCKJ -m comment --comment "kube-system/hubble-peer:peer-service" -j KUBE-SEP-N27GNE72QNUULDFK
-A KUBE-SVC-PDS5MBI2OEFEKTY7 -m comment --comment "default/deathstar" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-FUY4JREFWM7DZPDS
-A KUBE-SVC-PDS5MBI2OEFEKTY7 -m comment --comment "default/deathstar" -j KUBE-SEP-M5CX275RXKVJC7F7
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m comment --comment "kube-system/kube-dns:dns" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-JFY262ZQRD3MRSBY
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m comment --comment "kube-system/kube-dns:dns" -j KUBE-SEP-JTHDLAECD3YWPOED
-A KUBE-SVC-XXQFMBX2TZG6LPU6 -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp --dport 30080 -j KUBE-MARK-MASQ
-A KUBE-SVC-XXQFMBX2TZG6LPU6 -m comment --comment "default/x-tools-svc-1" -j KUBE-SEP-4NJI4U4OVD7ASSTN
-A KUBE-SVC-ZGWW2L4XLRSDZ3EF -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp --dport 30486 -j KUBE-MARK-MASQ
-A KUBE-SVC-ZGWW2L4XLRSDZ3EF -m comment --comment "kube-system/hubble-ui:http" -j KUBE-SEP-LNM6HHUZOJM3CKW7
```
svc nat的链规则调用秩序：
* KUBE-SERVICES (匹配svc clusterIP、端口、协议，送往kube-svc-xxx链) ->
* KUBE-SVC-ZGWW2L4XLRSDZ3EF
  * 匹配tcp协议，目标端口为NodePort端口，指定协议为tcp，送到KUBE-MARK-MASQ链 ->
  * 送到KUBE-SEP-LNM6HHUZOJM3CKW7链 ->
* KUBE-MARK-MASQ(打上0x4000/0x4000标记) ...
* KUBE-SEP-LNM6HHUZOJM3CKW7 
  * 匹配源ip为svc epIP，送到KUBE-MARK-MASQ链 ->
  * 匹配协议为tcp，指定协议，然后进行DNat操作
```shell
-N CILIUM_OUTPUT_nat
-N CILIUM_POST_nat
-N CILIUM_PRE_nat
-N KUBE-NODEPORTS
-N KUBE-PROXY-CANARY
-N KUBE-SERVICES
-N KUBE-SEP-4NJI4U4OVD7ASSTN
-N KUBE-SEP-4W4DMGANR5QBPYNW
-N KUBE-SEP-FUY4JREFWM7DZPDS
-N KUBE-SEP-IMPQAXS4DSXOT6OW
-N KUBE-SEP-JFY262ZQRD3MRSBY
-N KUBE-SEP-JTHDLAECD3YWPOED
-N KUBE-SEP-LNM6HHUZOJM3CKW7
-N KUBE-SEP-M5CX275RXKVJC7F7
-N KUBE-SEP-MV2OXEWOCKPWUWNN
-N KUBE-SEP-N27GNE72QNUULDFK
-N KUBE-SEP-Q5DFUOU55OKTF7SL
-N KUBE-SEP-SZOBJIZ4TCVR33ZB
-N KUBE-SEP-Y5LSJLRXX3OJOFLV
-N KUBE-SVC-DZSQUIY5BXN7HSA3
-N KUBE-SVC-ERIFXISQEP7F7OF4
-N KUBE-SVC-JD5MR3NA4I4DYORP
-N KUBE-SVC-NPX46M4PTMTKRN6Y
-N KUBE-SVC-NZTS37XDTDNXGCKJ
-N KUBE-SVC-PDS5MBI2OEFEKTY7
-N KUBE-SVC-TCOU7JCQXEZGVUNU
-N KUBE-SVC-XXQFMBX2TZG6LPU6
-N KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A PREROUTING -m comment --comment "cilium-feeder: CILIUM_PRE_nat" -j CILIUM_PRE_nat
-A PREROUTING -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A OUTPUT -m comment --comment "cilium-feeder: CILIUM_OUTPUT_nat" -j CILIUM_OUTPUT_nat
-A OUTPUT -m comment --comment "kubernetes service portals" -j KUBE-SERVICES
-A POSTROUTING -m comment --comment "cilium-feeder: CILIUM_POST_nat" -j CILIUM_POST_nat
# 在CILIUM_POST_nat链中添加规则
# 对源IP地址在10.0.0.0/24范围内，目的IP地址不在10.0.0.0/24范围内，且出接口以cilium_开头的流量进行MASQUERADE（源地址转换）处理。
# 对标记为0xa00/0xe00的流量进行ACCEPT（接受）处理。
# 对源IP地址不在10.0.0.0/24范围内，目的IP地址不在10.0.0.0/24范围内，且出接口为cilium_host的流量进行SNAT（目的地址转换）处理，转换后的源IP地址为10.0.0.96。
# 对源IP地址为127.0.0.1，出接口为cilium_host的流量进行SNAT（目的地址转换）处理，转换后的源IP地址为10.0.0.96。
# 对出接口为cilium_host，标记为0xf00/0xf00，且连接状态为DNAT的流量进行SNAT（目的地址转换）处理，转换后的源IP地址为10.0.0.96。
-A CILIUM_POST_nat -s 10.0.0.0/24 ! -d 10.0.0.0/24 ! -o cilium_+ -m comment --comment "cilium masquerade non-cluster" -j MASQUERADE
-A CILIUM_POST_nat -m mark --mark 0xa00/0xe00 -m comment --comment "exclude proxy return traffic from masquerade" -j ACCEPT
-A CILIUM_POST_nat ! -s 10.0.0.0/24 ! -d 10.0.0.0/24 -o cilium_host -m comment --comment "cilium host->cluster masquerade" -j SNAT --to-source 10.0.0.96
-A CILIUM_POST_nat -s 127.0.0.1/32 -o cilium_host -m comment --comment "cilium host->cluster from 127.0.0.1 masquerade" -j SNAT --to-source 10.0.0.96
-A CILIUM_POST_nat -o cilium_host -m mark --mark 0xf00/0xf00 -m conntrack --ctstate DNAT -m comment --comment "hairpin traffic that originated from a local pod" -j SNAT --to-source 10.0.0.96
-A KUBE-NODEPORTS -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp --dport 30080 -j KUBE-SVC-XXQFMBX2TZG6LPU6
-A KUBE-NODEPORTS -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp --dport 30486 -j KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A KUBE-SEP-4NJI4U4OVD7ASSTN -s 10.0.0.52/32 -m comment --comment "default/x-tools-svc-1" -j KUBE-MARK-MASQ
-A KUBE-SEP-4NJI4U4OVD7ASSTN -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp -j DNAT --to-destination 10.0.0.52:80
-A KUBE-SEP-4W4DMGANR5QBPYNW -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-MARK-MASQ
-A KUBE-SEP-4W4DMGANR5QBPYNW -p tcp -m comment --comment "kube-system/kube-dns:metrics" -m tcp -j DNAT --to-destination 10.0.0.61:9153
-A KUBE-SEP-FUY4JREFWM7DZPDS -s 10.0.0.39/32 -m comment --comment "default/deathstar" -j KUBE-MARK-MASQ
-A KUBE-SEP-FUY4JREFWM7DZPDS -p tcp -m comment --comment "default/deathstar" -m tcp -j DNAT --to-destination 10.0.0.39:80
-A KUBE-SEP-IMPQAXS4DSXOT6OW -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-MARK-MASQ
-A KUBE-SEP-IMPQAXS4DSXOT6OW -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp" -m tcp -j DNAT --to-destination 10.0.0.61:53
-A KUBE-SEP-JFY262ZQRD3MRSBY -s 10.0.0.61/32 -m comment --comment "kube-system/kube-dns:dns" -j KUBE-MARK-MASQ
-A KUBE-SEP-JFY262ZQRD3MRSBY -p udp -m comment --comment "kube-system/kube-dns:dns" -m udp -j DNAT --to-destination 10.0.0.61:53
-A KUBE-SEP-JTHDLAECD3YWPOED -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:dns" -j KUBE-MARK-MASQ
-A KUBE-SEP-JTHDLAECD3YWPOED -p udp -m comment --comment "kube-system/kube-dns:dns" -m udp -j DNAT --to-destination 10.0.1.248:53
-A KUBE-SEP-LNM6HHUZOJM3CKW7 -s 10.0.1.227/32 -m comment --comment "kube-system/hubble-ui:http" -j KUBE-MARK-MASQ
-A KUBE-SEP-LNM6HHUZOJM3CKW7 -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp -j DNAT --to-destination 10.0.1.227:8081
-A KUBE-SEP-M5CX275RXKVJC7F7 -s 10.0.1.168/32 -m comment --comment "default/deathstar" -j KUBE-MARK-MASQ
-A KUBE-SEP-M5CX275RXKVJC7F7 -p tcp -m comment --comment "default/deathstar" -m tcp -j DNAT --to-destination 10.0.1.168:80
-A KUBE-SEP-MV2OXEWOCKPWUWNN -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-MARK-MASQ
-A KUBE-SEP-MV2OXEWOCKPWUWNN -p tcp -m comment --comment "kube-system/kube-dns:metrics" -m tcp -j DNAT --to-destination 10.0.1.248:9153
-A KUBE-SEP-N27GNE72QNUULDFK -s 172.28.8.138/32 -m comment --comment "kube-system/hubble-peer:peer-service" -j KUBE-MARK-MASQ
-A KUBE-SEP-N27GNE72QNUULDFK -p tcp -m comment --comment "kube-system/hubble-peer:peer-service" -m tcp -j DNAT --to-destination 172.28.8.138:4244
-A KUBE-SEP-Q5DFUOU55OKTF7SL -s 10.0.1.248/32 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-MARK-MASQ
-A KUBE-SEP-Q5DFUOU55OKTF7SL -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp" -m tcp -j DNAT --to-destination 10.0.1.248:53
-A KUBE-SEP-SZOBJIZ4TCVR33ZB -s 10.0.1.32/32 -m comment --comment "kube-system/hubble-relay" -j KUBE-MARK-MASQ
-A KUBE-SEP-SZOBJIZ4TCVR33ZB -p tcp -m comment --comment "kube-system/hubble-relay" -m tcp -j DNAT --to-destination 10.0.1.32:4245
-A KUBE-SEP-Y5LSJLRXX3OJOFLV -s 172.28.8.138/32 -m comment --comment "default/kubernetes:https" -j KUBE-MARK-MASQ
-A KUBE-SEP-Y5LSJLRXX3OJOFLV -p tcp -m comment --comment "default/kubernetes:https" -m tcp -j DNAT --to-destination 172.28.8.138:6443
-A KUBE-SERVICES -d 10.106.170.104/32 -p tcp -m comment --comment "kube-system/hubble-peer:peer-service cluster IP" -m tcp --dport 443 -j KUBE-SVC-NZTS37XDTDNXGCKJ
-A KUBE-SERVICES -d 10.104.103.145/32 -p tcp -m comment --comment "kube-system/hubble-relay cluster IP" -m tcp --dport 80 -j KUBE-SVC-DZSQUIY5BXN7HSA3
-A KUBE-SERVICES -d 10.105.169.103/32 -p tcp -m comment --comment "default/x-tools-svc-1 cluster IP" -m tcp --dport 80 -j KUBE-SVC-XXQFMBX2TZG6LPU6
-A KUBE-SERVICES -d 10.96.0.10/32 -p udp -m comment --comment "kube-system/kube-dns:dns cluster IP" -m udp --dport 53 -j KUBE-SVC-TCOU7JCQXEZGVUNU
-A KUBE-SERVICES -d 10.96.0.10/32 -p tcp -m comment --comment "kube-system/kube-dns:dns-tcp cluster IP" -m tcp --dport 53 -j KUBE-SVC-ERIFXISQEP7F7OF4
-A KUBE-SERVICES -d 10.96.0.10/32 -p tcp -m comment --comment "kube-system/kube-dns:metrics cluster IP" -m tcp --dport 9153 -j KUBE-SVC-JD5MR3NA4I4DYORP
-A KUBE-SERVICES -d 10.107.246.98/32 -p tcp -m comment --comment "default/deathstar cluster IP" -m tcp --dport 80 -j KUBE-SVC-PDS5MBI2OEFEKTY7
-A KUBE-SERVICES -d 10.96.0.1/32 -p tcp -m comment --comment "default/kubernetes:https cluster IP" -m tcp --dport 443 -j KUBE-SVC-NPX46M4PTMTKRN6Y
-A KUBE-SERVICES -d 10.99.53.195/32 -p tcp -m comment --comment "kube-system/hubble-ui:http cluster IP" -m tcp --dport 80 -j KUBE-SVC-ZGWW2L4XLRSDZ3EF
-A KUBE-SERVICES -m comment --comment "kubernetes service nodeports; NOTE: this must be the last rule in this chain" -m addrtype --dst-type LOCAL -j KUBE-NODEPORTS
-A KUBE-SVC-DZSQUIY5BXN7HSA3 -m comment --comment "kube-system/hubble-relay" -j KUBE-SEP-SZOBJIZ4TCVR33ZB
-A KUBE-SVC-ERIFXISQEP7F7OF4 -m comment --comment "kube-system/kube-dns:dns-tcp" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-IMPQAXS4DSXOT6OW
-A KUBE-SVC-ERIFXISQEP7F7OF4 -m comment --comment "kube-system/kube-dns:dns-tcp" -j KUBE-SEP-Q5DFUOU55OKTF7SL
-A KUBE-SVC-JD5MR3NA4I4DYORP -m comment --comment "kube-system/kube-dns:metrics" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-4W4DMGANR5QBPYNW
-A KUBE-SVC-JD5MR3NA4I4DYORP -m comment --comment "kube-system/kube-dns:metrics" -j KUBE-SEP-MV2OXEWOCKPWUWNN
-A KUBE-SVC-NPX46M4PTMTKRN6Y -m comment --comment "default/kubernetes:https" -j KUBE-SEP-Y5LSJLRXX3OJOFLV
-A KUBE-SVC-NZTS37XDTDNXGCKJ -m comment --comment "kube-system/hubble-peer:peer-service" -j KUBE-SEP-N27GNE72QNUULDFK
-A KUBE-SVC-PDS5MBI2OEFEKTY7 -m comment --comment "default/deathstar" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-FUY4JREFWM7DZPDS
-A KUBE-SVC-PDS5MBI2OEFEKTY7 -m comment --comment "default/deathstar" -j KUBE-SEP-M5CX275RXKVJC7F7
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m comment --comment "kube-system/kube-dns:dns" -m statistic --mode random --probability 0.50000000000 -j KUBE-SEP-JFY262ZQRD3MRSBY
-A KUBE-SVC-TCOU7JCQXEZGVUNU -m comment --comment "kube-system/kube-dns:dns" -j KUBE-SEP-JTHDLAECD3YWPOED
-A KUBE-SVC-XXQFMBX2TZG6LPU6 -p tcp -m comment --comment "default/x-tools-svc-1" -m tcp --dport 30080 -j KUBE-MARK-MASQ
-A KUBE-SVC-XXQFMBX2TZG6LPU6 -m comment --comment "default/x-tools-svc-1" -j KUBE-SEP-4NJI4U4OVD7ASSTN
-A KUBE-SVC-ZGWW2L4XLRSDZ3EF -p tcp -m comment --comment "kube-system/hubble-ui:http" -m tcp --dport 30486 -j KUBE-MARK-MASQ
-A KUBE-SVC-ZGWW2L4XLRSDZ3EF -m comment --comment "kube-system/hubble-ui:http" -j KUBE-SEP-LNM6HHUZOJM3CKW7
```


mangle
```shell
[root@f1 ~]# iptables -t mangle -S
-P PREROUTING ACCEPT
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
-P POSTROUTING ACCEPT
-N CILIUM_POST_mangle
-N CILIUM_PRE_mangle
-N KUBE-KUBELET-CANARY
-N KUBE-PROXY-CANARY
-A PREROUTING -m comment --comment "cilium-feeder: CILIUM_PRE_mangle" -j CILIUM_PRE_mangle
-A POSTROUTING -m comment --comment "cilium-feeder: CILIUM_POST_mangle" -j CILIUM_POST_mangle
# 对进入的网络流量进行透明socket检查，如果标记不为0xe00/0xf00，则将流量的标记设置为0x200/0xffffffff。
# 对进入的TCP和UDP流量进行标记检查，如果标记为0xe9a40200，则将这些流量重定向到本机的42217端口，通过TPROXY进行代理。
# 对进入的TCP和UDP流量进行标记检查，如果标记为0xe62d0200，则将这些流量重定向到本机的11750端口，通过TPROXY进行代理。
-A CILIUM_PRE_mangle -m socket --transparent -m mark ! --mark 0xe00/0xf00 -m comment --comment "cilium: any->pod redirect proxied traffic to host proxy" -j MARK --set-xmark 0x200/0xffffffff
-A CILIUM_PRE_mangle -p tcp -m mark --mark 0xe9a40200 -m comment --comment "cilium: TPROXY to host cilium-dns-egress proxy" -j TPROXY --on-port 42217 --on-ip 127.0.0.1 --tproxy-mark 0x200/0xffffffff
-A CILIUM_PRE_mangle -p udp -m mark --mark 0xe9a40200 -m comment --comment "cilium: TPROXY to host cilium-dns-egress proxy" -j TPROXY --on-port 42217 --on-ip 127.0.0.1 --tproxy-mark 0x200/0xffffffff
-A CILIUM_PRE_mangle -p tcp -m mark --mark 0xe62d0200 -m comment --comment "cilium: TPROXY to host cilium-http-ingress proxy" -j TPROXY --on-port 11750 --on-ip 127.0.0.1 --tproxy-mark 0x200/0xffffffff
-A CILIUM_PRE_mangle -p udp -m mark --mark 0xe62d0200 -m comment --comment "cilium: TPROXY to host cilium-http-ingress proxy" -j TPROXY --on-port 11750 --on-ip 127.0.0.1 --tproxy-mark 0x200/0xffffffff
```

raw
```shell
[root@f1 ~]# iptables -t raw -S
-P PREROUTING ACCEPT
-P OUTPUT ACCEPT
-N CILIUM_OUTPUT_raw
-N CILIUM_PRE_raw
-A PREROUTING -m comment --comment "cilium-feeder: CILIUM_PRE_raw" -j CILIUM_PRE_raw
-A OUTPUT -m comment --comment "cilium-feeder: CILIUM_OUTPUT_raw" -j CILIUM_OUTPUT_raw
# 
-A CILIUM_OUTPUT_raw -o lxc+ -m mark --mark 0xa00/0xfffffeff -m comment --comment "cilium: NOTRACK for proxy return traffic" -j CT --notrack
-A CILIUM_OUTPUT_raw -o cilium_host -m mark --mark 0xa00/0xfffffeff -m comment --comment "cilium: NOTRACK for proxy return traffic" -j CT --notrack
-A CILIUM_OUTPUT_raw -o lxc+ -m mark --mark 0x800/0xe00 -m comment --comment "cilium: NOTRACK for L7 proxy upstream traffic" -j CT --notrack
-A CILIUM_OUTPUT_raw -o cilium_host -m mark --mark 0x800/0xe00 -m comment --comment "cilium: NOTRACK for L7 proxy upstream traffic" -j CT --notrack
-A CILIUM_PRE_raw -m mark --mark 0x200/0xf00 -m comment --comment "cilium: NOTRACK for proxy traffic" -j CT --notrack
```