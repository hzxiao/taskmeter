#docker
yum install docker-io -y

service docker start

chkconfig docker on

echo "OPTIONS='--registry-mirror=https://mirror.ccs.tencentyun.com'" >> /etc/sysconfig/docker

systemctl daemon-reload

service docker restart

# jenkins 
yum -y install java-1.8.0-openjdk-devel
sudo wget -O /etc/yum.repos.d/jenkins.repo https://pkg.jenkins.io/redhat-stable/jenkins.repo
sudo rpm --import https://pkg.jenkins.io/redhat-stable/jenkins.io.key
yum -y install jenkins

systemctl start jenkins.service
chkconfig jenkins on

sudo firewall-cmd --zone=public --add-port=8080/tcp --permanent
sudo firewall-cmd --reload

#git
yum install curl-devel expat-devel gettext-devel openssl-devel zlib-devel
yum install gcc perl-ExtUtils-MakeMaker

cd /usr/local/src

wget https://www.kernel.org/pub/software/scm/git/git-2.10.0.tar.gz
tar -zvxf git-2.10.0.tar.gz
cd git-2.10.0
make all prefix=/usr/local/git
echo 'export PATH=$PATH:/usr/local/git/bin' >> /etc/bashrc
source /etc/bashrc


# MongoDB
yum -y install mongodb-org  

sudo service mongod start

sudo chkconfig mongod on