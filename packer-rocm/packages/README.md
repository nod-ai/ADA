# packages

Files placed here will be copied into the _Packer_ builder VM. The _intended_ use includes `.deb` or `.rpm` files.

An [Ansible playbook](../playbooks/package-globber.yml) used as a [provisioner](https://developer.hashicorp.com/packer/docs/provisioners/) will install those which are distribution-appropriate.
