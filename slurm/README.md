# Ansible Slurm Installation

## Purpose

The purpose of these playbooks are to install a slurm cluster extracting information about the cluster makeup from the ansible inventory file.

### Requirements

* `ansible-core`, examples: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)

### Installation

All of the relevant configuration data for the slurm cluster is either provided in the inventory file or it is provided in the files section of the slurm directory.

The config files are:

* cgroups.conf 
* slurm.conf
* slurmdbd.conf

In addition to the config files the munge.key for the slurm nodes is placed into that directory.

#### Example inventory.yml

```yaml
all:
  children:
    slurm_compute_nodes:
      hosts:
        ansible-test01:
        ansible-test02:
        ansible-test03:
        ansible-test04:
    slurm_head_node:
      hosts:
        ansible-test01:
    slurmdbd_node:
      hosts:
        ansible-test01:
  vars:
    ansible_user: ubuntu
    cluster_name: vlan3051
    slurm_partition: test
    slurm_config_file: /etc/slurm/slurm.conf
    cgroup_config_file: /etc/slurm/cgroup.conf
    munge_key_file: /etc/munge/munge.key
```

#### Running the playbook

```shell
ansible-playbook -i inventory.yml install_slurm_all.yml
```