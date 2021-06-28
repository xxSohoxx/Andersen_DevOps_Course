# About the task #
In this task we need to create and deploy our own service. It is simple Flask app which do the following:  
* the service listens at least on port 80 (443 as an option)
* the service accepts GET and POST methods
* the service should receive `JSON` object and return strings in the following manner:
```sh
curl -XPOST -d'{"animal":"cow", "sound":"moooo", "count": 3}' http://myvm.localhost/
cow says moooo
cow says moooo
cow says moooo
Made with ❤️ by %your_name
```
The operating stage:
* create an ansible playbook that deploys the service to the VM
* make sure all the components you need are installed and all the directories for the app are present
* configure systemd so that the application starts after reboot
* secure the VM so that our product is not stolen: allow connections only to the ports 22,80,443. Disable root login. Disable all authentication methods except 'public keys'.

## How to start ##
### Prerequisites ###
You need to have Debian 10 server, where we are going to deploy the service, and be able to access it via SSH + RSA key. Your user account needs to have SUDO privileges.  
Here are some links regarding prerequisites configuration for your remote VM:  
* [How to install SSH server](https://phoenixnap.com/kb/how-to-enable-ssh-on-debian)
* [How to setup a Public Key authentication](https://www.ssh.com/academy/ssh/copy-id)  

This is an Ansible task, so you also need to have Ansible installed on your working station.  
* [How to install Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-ansible-on-debian)  


### Deploy the service ###
* You need to clone current git repository by sending a command: *git clone https://github.com/xxSohoxx/Andersen_DevOps_Course* 
* Then you need to change directory: *cd Andersen_DevOps_Course/ansible_task/*
* And run a command: *ansible-playbook  -i ./inventory/inventory -K ./ansible/deploy.yml*
