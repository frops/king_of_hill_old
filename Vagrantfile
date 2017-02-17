VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

    config.vm.provider "virtualbox" do |v|
      v.memory = 1024
    end

	config.vm.box = "deb/jessie-amd64"
	# config.vm.box = "debian/jessie64"

	# expose ports
	config.vm.network :forwarded_port, host: 8083, guest: 80
	config.vm.network :forwarded_port, host: 8084, guest: 1111
	config.vm.network :forwarded_port, host: 3310, guest: 3306
	config.vm.network :forwarded_port, host: 27017, guest: 27017

	# Setup shared dir with www-data owner
	config.vm.synced_folder ".", "/vagrant", :user => "www-data", :group => "www-data"

	## Run provision setup script
	config.vm.provision "shell", :path => "vagrant/provision.sh"
end
