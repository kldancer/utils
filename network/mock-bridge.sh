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
add_vlan_network vmbr0 192.120.50.1/24 ipv4 50
add_vlan_network vmbr0 192.120.51.1/24 ipv4 51

iptables -t nat -I POSTROUTING  -o  vmbr0 -j ACCEPT


