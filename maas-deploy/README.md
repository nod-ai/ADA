# MaaS deployment

## Purpose
Canonical's MaaS is used to provision bare metal systems for ADA deployments. This section is meant to provide the necessary resources for deploying a MaaS provisioning system. For more information about MaaS see https://maas.io

### Requirements

* `ansible-core`, examples: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* [maas-ansible-playbook](https://github.com/canonical/maas-ansible-playbook)
* 2 physical or virtual machines running Ubuntu 22.04 

### Installation

1.  Clone the maas-ansible-playbook repository on the machine with ansible-core installed.

    ```shell
    git clone https://github.com/canonical/maas-ansible-playbook.git
    ```
    * NOTE:
        
        There is an issue with MaaS 3.4 and 3.5 which requires the following code to be placed at the bottom of site.yaml

        ```yaml
        #
        # added this as a bug fix for database structure in 3.4 and 3.5 
        #
        - hosts: maas_region_controller
          become: true
          gather_facts: true
          tasks:
            - name: "BUG FIX sudo maas-region dbupgrade"
              ansible.builtin.command: sudo maas-region dbupgrade
        ```

2.  Create an inventory file for the machines for deployment. One machine will run the postgres backend for Maas, and the other will be a combined region/rack controller.
    ```shell
    ---
    all:
    children:
        maas_postgres:
        hosts:
            db01.example.com:
        maas_region_controller:
        hosts:
            region01.example.com:
        maas_rack_controller:
        hosts:
            region01.example.com:

    vars:
        # required vars
        maas_version: 3.5 
        maas_postgres_password: <password>
        maas_installation_type: deb
        maas_url: http://region01.example.com:5240/MAAS
        # optional vars
        # admin creds
        admin_username: admin
        admin_password: <password>
        admin_email: email@example.com
        admin_id: # ssh keys to be added to the admin user
        maas_proxy_postgres_proxy_enabled: false
        enable_tls: false
        enable_firewall: false
    ```


3.  Install the two vm's with Ubuntu 22.04

4.  Execute the playbook

    ```shell
    ansible-playbook -i inventory.yaml site.yaml
    ```
### Configuration
Once the installation is complete the rest of the configuration is done from the MaaS web UI which is accessible at 

    maas_url: http://region01.example.com:5240/MAAS