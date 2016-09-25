sudo apt-get -qq update
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password password root'
sudo debconf-set-selections <<< 'mysql-server mysql-server/root_password_again password root'
sudo apt-get -qq install -y git mysql-server mercurial
[ ! -e go1.7.linux-amd64.tar.gz ] && wget -q https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz
sudo tar -xf go1.7.linux-amd64.tar.gz -C /opt
source /home/vagrant/.profile
if [ -z $GOSET ] 
    then
    echo "export PATH=/opt/go/bin:$PATH" >> /home/vagrant/.profile
    echo "export GOROOT=/opt/go" >> /home/vagrant/.profile
    echo "export GOPATH=/home/vagrant/gocode" >> /home/vagrant/.profile
    echo "export GOSET=True" >> /home/vagrant/.profile
    source /home/vagrant/.profile
    cd /home/vagrant/gocode/src/github.com/gulyasm/focus
    go get
    go get -u github.com/golang/lint/golint
    sudo chown -R vagrant:vagrant /home/vagrant/gocode
    cd -
fi
