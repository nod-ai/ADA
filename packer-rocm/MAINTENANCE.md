# Maintenance

Both _Packer_ and _Ansible_ influence variables in the `packer build` or `ansible-playbook` scopes _(interchangeably)_. Any defaults users _may_ change _should_ be defined in `./${dist}/${dist}-rocm.variables.pkr.hcl` _as well_ as the playbook or role.

The build playbook [build.yml](./playbooks/build.yml) will respect variables defined in the distribution HCL for the _'personalization playbooks'_ in '[playbooks](./playbooks)'. These are run as _'provisioners'_ by _Packer_ on a _QEMU_ VM _(Virtual Machine)_ to create the image.

Ansible variables that are constants _(unlikely to be changed by a typical user)_ or don't pass through _Packer_ can be left out of the HCL. Mind conversion: Ansible supports a wider collection of types than Packer.
