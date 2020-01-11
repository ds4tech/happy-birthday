#!/bin/bash

# volume setup
vgchange -ay

DEVICE_FS=`blkid -o value -s TYPE ${DEVICE}`
if [ "`echo -n $DEVICE_FS`" == "" ] ; then
  # wait for the device to be attached
  DEVICENAME=`echo "${DEVICE}" | awk -F '/' '{print $3}'`
  DEVICEEXISTS=''
  while [[ -z $DEVICEEXISTS ]]; do
    echo "checking $DEVICENAME"
    DEVICEEXISTS=`lsblk |grep "$DEVICENAME" |wc -l`
    if [[ $DEVICEEXISTS != "1" ]]; then
      sleep 15
    fi
  done
  pvcreate ${DEVICE}
  vgcreate data ${DEVICE}
  lvcreate --name volume1 -l 100%FREE data
  mkfs.ext4 /dev/data/volume1
fi
mkdir -p /var/lib/jenkins
echo '/dev/data/volume1 /var/lib/jenkins ext4 defaults 0 0' >> /etc/fstab
mount /var/lib/jenkins

# install jenkins and docker
wget -q -O - https://pkg.jenkins.io/debian-stable/jenkins.io.key | sudo apt-key add -
echo "deb http://pkg.jenkins.io/debian-stable binary/" >> /etc/apt/sources.list
add-apt-repository ppa:longsleep/golang-backports -y
apt-get update
apt-get install -y jenkins=${JENKINS_VERSION} unzip docker.io

# install dependencies
apt-get install -y python3 openjdk-8-jre
update-java-alternatives --set java-1.8.0-openjdk-amd64

# enable docker and add perms
usermod -G docker jenkins
systemctl enable docker
service docker start
service jenkins restart

# install pip
wget -q https://bootstrap.pypa.io/get-pip.py
python get-pip.py
#python3 get-pip.py
rm -f get-pip.py
# install awscli
pip install awscli

# install terraform
wget -q https://releases.hashicorp.com/terraform/$${TERRAFORM_VERSION}/terraform_$${TERRAFORM_VERSION}_linux_amd64.zip \
&& unzip -o terraform_$${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/local/bin \
&& rm terraform_$${TERRAFORM_VERSION}_linux_amd64.zip

# install go-lang
apt-get install -y golang-go
echo "export GOPATH=/var/lib/jenkins/go" >> /etc/profile
mkdir /var/lib/jenkins/go
chown jenkins.jenkins /var/lib/jenkins/go
go get github.com/gorilla/mux

# clean up
apt-get clean
