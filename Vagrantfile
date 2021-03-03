# -*- mode: ruby -*-
# # vi: set ft=ruby :

## Version constraint
Vagrant.require_version ">= 1.3.5"

## Plugin constraint
unless Vagrant.has_plugin?("vagrant-hostsupdater")
  raise 'Missing plugin `vagrant-hostsupdater`! Install it by `vagrant plugin install vagrant-hostsupdater` command before vagrant up.'
end

# VM config
Vagrant.configure("2") do |config|
  config.ssh.insert_key = false
  config.vm.box = "centos/7"
  config.vm.box_version = "2004.01"
  # config.vm.box_url = "https://cloud.centos.org/centos/7/vagrant/x86_64/images/CentOS-7-x86_64-Vagrant-2004_01.VirtualBox.box"
  config.vm.hostname = "vagrant-test.localhost.com"
  config.vm.network :forwarded_port, guest: 5000, host: 8080, id: "http", protocol: "tcp"
  config.vm.network :private_network, ip: "192.168.10.10"
  config.vm.provider :virtualbox do |v|
      v.name = "centos7_vagrant_test"
      v.customize ["modifyvm", :id, "--memory", 2048]
      v.cpus = 1
      v.check_guest_additions = false
      v.functional_vboxsf     = false
      v.gui                   = false
  end

  # VM NFS
  config.vm.synced_folder ".", "/home/cwd/src", id: "home", type: "nfs", :nfs => true, :mount_options => ['nolock,vers=3,udp,actimeo=2']

  # VM provisioning
  config.vm.provision "shell", inline: <<-SHELL
    echo Hello, World
  SHELL
  config.vm.provision "shell", :path => "provision/local/scripts/docker-init.sh"
  # config.vm.provision "shell", inline: <<-SHELL
  #   /home/init.sh
  # SHELL
  # $DOCKER_COMPOSE_VERSION="1.28.2"
  # config.vm.provision "shell", :path => "home/docker.sh", :args => [$DOCKER_COMPOSE_VERSION]
  # config.ssh.forward_agent = true


  # Avoid plugin conflicts
  if Vagrant.has_plugin?("vagrant-vbguest") then
    config.vbguest.auto_update = false
  end

end