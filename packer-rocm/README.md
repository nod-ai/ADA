# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with `amdgpu-dkms` and _(optional)_ [ROCm](https://www.amd.com/en/products/software/rocm.html). Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.


## Building

### Prerequisites

* `git`
* `ansible`: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* _Linux_ host

### Setup

```shell
git clone https://github.com/nod-ai/ADA.git
ansible-galaxy collection install -r ADA/packer-rocm/requirements.yml
```

Place any `.deb` packages to include with the image in `ADA/packer-rocm/ubuntu/packages/`

### Build

```shell
ansible-playbook ADA/packer-rocm/playbooks/build.yml \
    -e rocm_releases="6.2.2,6.2.1" \
    -e rocm_kernel="linux-image-generic-hwe-22.04" \
    -e rocm_extras="linux-headers-generic-hwe-22.04,mesa-amdgpu-va-drivers" \
    -e rocm_builder_cpus=16 \
    -e rocm_builder_disk="70G" \
    -e qemu_binary="qemu-kvm" \
    -K
```

Remove `-K` if your account does _not_ require a passphrase for `sudo`. This is used to prepare the host _(repositories and packages)_.

**All** of these variables are _optional_. Please see [I/O](#io) for more.

### I/O

| Variable | Description | Default |
|:----------:|-------------|:---------:|
| `qemu_binary` | The name _or_ path for the _QEMU_ binary. | `qemu-system-x86_64` |
| `packer_binary` | The name _or_ path for the _Packer_ binary.<br/>Installation skipped if overridden. | `/usr/bin/packer` |
| `rocm_releases` | One or more versions to include _[comma-separated]_.<br/>Newest selects the `amdgpu` driver. | `6.2.2` |
| `rocm_kernel` | The kernel package with an optional release specifier. | `linux-image-generic-hwe-22.04` |
| `rocm_extras` | Packages to install _before_ `amdgpu-dkms` and _ROCm_.<br/>Comma-separated. May include releases with `=x.y.z` or globbing. | `linux-headers-generic-hwe-22.04`<br/>`mesa-amdgpu-va-drivers` |
| `rocm_filename` | The name of the output file/artifact _(tarball)_ | `ubuntu-rocm.tar.gz` |
| `rocm_installed` | If _ROCm_ multi-release packages are installed.<br/>The `amdgpu` _driver/extras_ are, always. | `False` |
| `rocm_builder_cpus` | Number of virtual CPUs given to the builder VM. | _4_ |
| `rocm_builder_disk` | Space given to the builder; releases compound quickly. | _70G_ |
| `rocm_builder_memory` | Megabytes of memory given to the builder.<br/>Reduction may cause out-of-memory conditions. | _4096_ |
| `niccli_wanted` | If [niccli](https://techdocs.broadcom.com/us/en/storage-and-ethernet-connectivity/ethernet-nic-controllers/bcm957xxx/adapters/Configuration-adapter/nic-cli-configuration-utility.html) is included in the image. | `True` |
| `niccli_url` | The URL for the _Broadcom_ `niccli` installation archive. | [Link](https://docs.broadcom.com/docs-and-downloads/ethernet-network-adapters/NXE/Thor2/GCA1/bcm5760x_230.2.52.0a.zip) |
| `niccli_sum` | _Optional_. Checksum to validate `niccli_url` downloads.<br/>Example: `sha256:abcd1234` | _Undefined_ |
| `niccli_driver` | If the `bnxt_{en,re}` NIC drivers are included. | `True` |
| `headless` | If the VNC window for the VM is _hidden_ during build. | `True` |

#### MaaS

This image is built _(and uploaded)_ in `tgz` format to respect customized disk layouts:

```shell
maas admin boot-resources create \
     name='custom/packer-rocm-ubuntu-22.04.5' \
     title='packer-rocm (Ubuntu 22.04.5)' \
     architecture='amd64/generic' \
     filetype='tgz' \
     content@=ubuntu-rocm.tar.gz
```

The artifact will be in `ADA/packer-rocm/packer-maas/ubuntu/`, named `ubuntu-rocm.tar.gz`.

#### Proxy

These _environment variables_ are respected:

* `http_proxy`
* `https_proxy`
* `no_proxy`
