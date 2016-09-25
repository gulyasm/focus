# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"


  config.vm.synced_folder ".", "/home/vagrant/gocode/src/github.com/gulyasm/focus"

  config.vm.provision "shell", path: "vagrant-init.sh"
end
