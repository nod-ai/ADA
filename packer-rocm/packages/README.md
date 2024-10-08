# packages

Files placed here will be copied into the _Packer_ builder VM. The _intended_ use includes `.deb` or `.rpm` files.

Packages appropriate for the image will be installed by an [Ansible playbook](../playbooks/package-globber.yml), used by _Packer_ as a [provisioner](https://developer.hashicorp.com/packer/docs/provisioners/).
