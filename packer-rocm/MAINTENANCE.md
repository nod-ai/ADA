# Maintenance

Both _Packer_ and _Ansible_ influence variable in the `packer build` or `ansible-playbook` scopes interchangeably. Any defaults users may change should be defined in `./${dist}/${dist}-rocm.variables.pkr.hcl` _as well_ as the playbook or role.

The build playbook [build.yml](./playbooks/build.yml) will respect variables defined in the distribution HCL, passing them to the _'personalization playbooks'_ in '[./playbooks](./playbooks)' that run against the builder VM to customize the image.

Ansible variables that are constants _(unlikely to be changed by a typical user)_ or don't pass through _Packer_ can be left out of the HCL. Mind conversion: Ansible supports a wider collection of types than Packer.
