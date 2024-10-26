# amdgpu\_install

Ansible role for [amdgpu-install](https://rocm.docs.amd.com/projects/install-on-linux/en/latest/install/amdgpu-install.html#) _Linux_ installation method

## Role Variables

| `amdgpu_install_rel` | The _release_ portion of the `amdgpu-install` URL.<br/>**Default:** _6.2.3_, [Reference](https://rocm.docs.amd.com/projects/install-on-linux/en/latest/install/amdgpu-install.html) |
| `amdgpu_install_build` | The _build_ portion of the URL.<br/>**Default:** _6.2.60203-1_ |
| `amdgpu_install_pkg` | Optional _URL_ for the `amdgpu-install` package. Replaces _`amdgpu\_install\_rel`_ and `amdgpu\_install\_build`_ above.<br/>**Default:** _templated_ |
| `amdgpu_install_usecases` | Package groups to request from `amdgpu-install` that match typical workflows and runtimes.<br/>**Default:** _dkms_, suggested: _rocm_ |
| `amdgpu_install_args` | Optional dictionary of arguments to pass to `amdgpu-install` |
| `amdgpu_install_branch` | Optional development branch of the `amdgpu` driver with `amdgpu-install-internal`<br/>**Default:** _Skipped_ |
| `amdgpu_install_rocm_branch` | Optional development branch for `rocm` software with `amdgpu-install-internal`<br/>**Default:** _Skipped_ |

Please see [defaults](./defaults/main.yml) or [vars](./vars/main.yml) for the most current references.

## Requirements

1. `community.general`

## Example Playbook

```yaml
    ---
    - name: Manage 'amdgpu-install'
      hosts: default
      environment:  # may be superfluous for your environment; mapped through Packer HCL with 'ansible_env_vars'
        http_proxy: "{{ lookup('ansible.builtin.env', 'http_proxy') | default(omit) }}"
        https_proxy: "{{ lookup('ansible.builtin.env', 'https_proxy') | default(omit) }}"
        no_proxy: "{{ lookup('ansible.builtin.env', 'no_proxy') | default(omit) }}"
      roles:
        - { role: amdgpu_install }
```

## License

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
