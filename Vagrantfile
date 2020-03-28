# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  # Every Vagrant virtual environment requires a box to build off of.
  #config.vm.box = "ubuntu/xenial64"
  #config.vm.box = "ubuntu/bionic64"
  config.vm.box = "minimal/xenial64"
  
  config.vm.hostname = "node-1"

  # Create a private network, which allows host-only access to the machine using a specific IP.
  config.vm.network "private_network", ip: "192.168.100.15"

  # Share an additional folder to the guest VM. The first argument is the path on the host to the actual folder.
  # The second argument is the path on the guest to mount the folder.
  config.vm.synced_folder "./", "/files"

  # Define the bootstrap file: A (shell) script that runs after first setup of your box (= provisioning)
  config.vm.provision "fix-no-tty", type: "shell" do |s|
    s.privileged = false
    s.inline = "sudo sed -i '/tty/!s/mesg n/tty -s \\&\\& mesg n/' /root/.profile"
  end
  config.vm.provision :shell, path: "bootstrap.sh"
  config.vm.provider "virtualbox" do |v|
    v.memory = 4096
    v.cpus = 2
  end
end
