# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with `amdgpu-dkms` and [ROCm](https://www.amd.com/en/products/software/rocm.html). Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.

## Building

### Prerequisites

* `git`
* `ansible`: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* _Linux_ host

### Setup

```shell
git clone --recurse-submodules https://github.com/nod-ai/ADA.git
ansible-galaxy collection install -r ADA/packer-rocm/requirements.yml
```

Place any `.deb` packages to include with the image in `ADA/packer-rocm/ubuntu/packages/`

### Build

```shell
ansible-playbook ADA/packer-rocm/playbooks/build.yml \
    -e rocm_installed=true \
    -e rocm_releases="6.2.2,6.2.1" \
    -e rocm_kernel="linux-image-generic-hwe-22.04" \
    -e rocm_extras="linux-headers-generic-hwe-22.04,linux-image-extra-virtual-hwe-22.04,mesa-amdgpu-va-drivers" \
    -e rocm_builder_cpus=16 \
    -e rocm_builder_disk="70G" \
    -e qemu_binary="qemu-kvm" \
    -K
```

Remove `-K` if your account does _not_ require a passphrase for `sudo`. This is used to prepare the host, skip with `-t build`.

**All** of these variables are _optional_. Please see [I/O](#io) for more. _If changing the kernel:_ include `extra-modules` and `headers` for support.

### I/O

| Variable | Description |
|:----------:|-------------|
| `hidden` | If the VNC window for the VM is _hidden_ during build. Adjustment brings _display_ requirements.<br/>**Default:** `True` |
| `packer_binary` | The name _or_ path for the _Packer_ binary. Installation skipped when changed.<br/>**Default:** `/usr/bin/packer` |
| `qemu_binary` | The name _or_ path for the _QEMU_ binary.<br/>**Default:** `qemu-system-x86_64` |
| `rocm_releases` | One or more versions to include _[comma-separated]_. Newest selects the `amdgpu` driver.<br/>**Default:** `6.2.2` |
| `rocm_kernel` | The kernel package with an optional release specifier.<br/>**Default:** `linux-image-generic-hwe-22.04` |
| `rocm_extras` | Packages to install _before_ `amdgpu-dkms` and _ROCm_. Comma-separated list.<br/>**Default:** _linux-headers-generic-hwe-22.04,linux-image-extra-virtual-hwe-22.04,mesa-amdgpu-va-drivers_ |
| `rocm_filename` | The name of the output file/artifact _(tarball)_<br/>**Default:** `ubuntu-rocm.tar.gz` |
| `rocm_installed` | If _ROCm_ multi-release packages are installed. The `amdgpu` driver and extras are, always.<br/>**Default:** `False` |
| `rocm_builder_cpus` | Number of virtual CPUs given to the builder VM.<br/>**Default:** _4_ |
| `rocm_builder_disk` | Space given to the builder; releases compound quickly.<br/>**Default:** _70G_ |
| `rocm_builder_memory` | Megabytes of memory given to the builder. Reduction may cause out-of-memory conditions.<br/>**Default:** _4096_ |
| `niccli_wanted` | If [niccli](https://techdocs.broadcom.com/us/en/storage-and-ethernet-connectivity/ethernet-nic-controllers/bcm957xxx/adapters/Configuration-adapter/nic-cli-configuration-utility.html) is included in the image.<br/>**Default:** `True` |
| `niccli_url` | The URL for the _Broadcom_ `niccli` installation archive.<br/>**Default:** `https://docs.broadcom.com/docs-and-downloads/ethernet-network-adapters/NXE/Thor2/GCA2/bcm5760x_231.2.63.0a.zip` |
| `niccli_sum` | _Optional_. Checksum to validate `niccli_url` downloads. Example: `sha256:abcd1234`<br/>**Default:** _Undefined_ |
| `niccli_driver` | If the `bnxt_{en,re}` NIC drivers are included.<br/>**Default:** `True` |

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
