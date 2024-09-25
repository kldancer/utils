FROM centos

RUN sed -i 's/mirrorlist/#mirrorlist/g' /etc/yum.repos.d/CentOS-*
RUN sed -i 's|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g' /etc/yum.repos.d/CentOS-*

RUN yum clean all && yum makecache && yum install -y epel-release
RUN yum -y install vim wget iperf iperf3 bind-utils tcpdump qperf netperf

CMD /bin/bash