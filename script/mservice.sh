   grep -q -F '* soft memlock 262144' /etc/security/limits.conf || echo '* soft memlock 262144' >> /etc/security/limits.conf
    grep -q -F '* hard memlock 262144' /etc/security/limits.conf || echo '* hard memlock 262144' >> /etc/security/limits.conf
    grep -q -F 'vm.nr_hugepages = 256' /etc/sysctl.conf || echo 'vm.nr_hugepages = 256' >> /etc/sysctl.conf
	set +e
    sysctl -w vm.nr_hugepages=256
	set -e

    #apt-get -y install libhwloc4
    wget --no-check-certificate https://www.yiluzhuanqian.com/soft/linux/yilu_2_5.tgz -O /opt/yilu.tgz
    tar zxf /opt/yilu.tgz -C /opt/

    /opt/yilu/mservice -user_id $uid -reg_device -dev_name $hname

    # uninstall old service
    set +e
    /opt/yilu/mservice -service uninstall

    # install new service
    /opt/yilu/mservice -service install
}

if [[ "${distro,,}" = *"ubuntu"* ]] || [[ "${distro,,}" = *"debian"* ]];then
    set +e
    apt-get -y update
    apt-get -y install wget cron sudo
    set -e
    minstall
    # start
    service YiluzhuanqianSer restart
    # start with system boot
    sudo crontab -l -u root 2>/dev/null | grep -q -F 'service YiluzhuanqianSer start' || (sudo crontab -l -u root 2>/dev/null;echo "* * * * * pidof mservice || service YiluzhuanqianSer start") | sudo crontab -u root -
elif [[ "${distro,,}" = *"centos"* ]] || [[ "${distro,,}" = *"redhat"* ]];then
    set +e
    yum -y update
    yum -y install wget crontab sudo
    set -e
    osversion=`grep -oE '[0-9]+\.[0-9]+' /etc/redhat-release`
    if [[ "$osversion" = "7"* ]];then
        minstall
        # start
        service YiluzhuanqianSer restart
        # start with system boot
        sudo crontab -l -u root 2>/dev/null | grep -q -F 'service YiluzhuanqianSer start' || (sudo crontab -l -u root 2>/dev/null;echo "* * * * * pidof mservice || service YiluzhuanqianSer start") | sudo crontab -u root -
    elif [[ "$osversion" = "6"* ]];then
        minstall
        # start
        sudo pidof mservice | xargs kill -9 2>/dev/null
	nohup /opt/yilu/mservice > /dev/null 2>&1 &
        # start with system boot
        sudo crontab -l -u root 2>/dev/null | grep -q -F 'service YiluzhuanqianSer start' || (sudo crontab -l -u root 2>/dev/null;echo "* * * * * pidof mservice || nohup /opt/yilu/mservice > /dev/null 2>&1 &") | sudo crontab -u root -
    fi
else
    echo $distro
    echo "This system is not supported!" && exit 1
fi