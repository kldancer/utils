#!/bin/bash

iptables -P FORWARD ACCEPT
sysctl -w net.ipv4.ip_forward=1
sysctl -w net.ipv6.conf.all.forwarding=1

function add_network() {
  br=$1
  cidr=$2
  v6=$3

  brctl addbr $br
  ip a a $cidr dev $br
  ip link set $br up
}


function add_vlan_network() {
    br=$1
    cidr=$2
    v6=$3
    vid=$4

    ip link add link $br name $br.$vid type vlan id $vid
    ip link set $br.$vid up
    ip address add $cidr dev $br.$vid
}

add_network vmbr0 192.110.0.1/24 ipv4
add_network vmbr0 2002:1e::1/64 ipv6

add_network vmbr1 192.120.0.1/24 ipv4
add_network vmbr1 192.120.1.1/24 ipv4
add_network vmbr1 192.120.2.1/24 ipv4

add_network vmbr2 192.130.0.1/24 ipv4
add_network vmbr2 2002:1a::1/64 ipv6
add_vlan_network vmbr2 192.130.40.1/24 ipv4 40
add_vlan_network vmbr2 192.130.41.1/24 ipv4 41

add_vlan_network vmbr1 192.120.31.1/24 ipv4 31
add_vlan_network vmbr1 192.120.32.1/24 ipv4 32
add_vlan_network vmbr1 192.120.33.1/24 ipv4 33
add_vlan_network vmbr1 2002:33::1/120 ipv4 33
add_vlan_network vmbr1 2002:60::1/120 ipv4 60
add_vlan_network vmbr1 2002:61::1/120 ipv4 61

add_vlan_network vmbr0 192.120.50.1/24 ipv4 50
add_vlan_network vmbr0 192.120.51.1/24 ipv4 51

iptables -t nat -I POSTROUTING  -o  vmbr0 -j ACCEPT
iptables -t nat -I POSTROUTING  -o  vmbr1 -j ACCEPT
iptables -t nat -I POSTROUTING  -o  vmbr2 -j ACCEPT
iptables -t nat -A POSTROUTING  -o  wlp44s0 -j MASQUERADE
iptables -t nat -A POSTROUTING  -o  tun0 -j MASQUERADE
iptables -t nat -A POSTROUTING  -o  tun1 -j MASQUERADE
