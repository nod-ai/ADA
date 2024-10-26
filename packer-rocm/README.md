# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with `amdgpu-dkms` and [ROCm](https://www.amd.com/en/products/software/rocm.html). Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.

## Building

### Prerequisites

* _Linux_ host
* `ansible`: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* `git`

### Setup

```shell
git clone --recurse-submodules https://github.com/nod-ai/ADA.git
ansible-galaxy collection install -r ADA/packer-rocm/requirements.yml
```

#### Drop-ins

* [./ubuntu/packages/](./ubuntu/packages/): `.deb` packages to include with the image, processed by [curtin](https://curtin.readthedocs.io/en/latest/topics/overview.html).
* [./repositories/](./repositories/): `.repo` or `.list` package sources for _Apt_, _Yum_, or _DNF_.
  * File permissions are retained _(eg: credentials)_.
  * Select packages they offer with `-e rocm_extras=pkg1,pkg2`; accepts releases.

### Build

```shell
ansible-playbook ADA/packer-rocm/playbooks/build.yml \
    -e amdgpu_install_usecases=rocm \
    -e rocm_builder_cpus=8 \
    -e qemu_binary="qemu-kvm" \
    -K
```

Remove `-K` if your account does _not_ require a passphrase for `sudo`. This is used to prepare the host, skip with `-t build`.

When _changing the kernel:_ include `extra-modules`, `headers`, and other packaged build requirements with `-e rocm_extras`. Versions are suggested.

**All** variables are _optional_. Please see [I/O](#io) for more.

### I/O

| Variable | Description |
|:----------:|-------------|
| `hidden` | If the VNC window for the VM is _hidden_ during build. Adjustment brings _display_ requirements.<br/>**Default:** `True` |
| `amdgpu_install_rel` | The _release_ portion of the `amdgpu-install` URL.<br/>**Default:** _6.2.2_, [Reference](https://rocm.docs.amd.com/projects/install-on-linux/en/latest/install/amdgpu-install.html) |
| `amdgpu_install_build` | The _build_ portion of the URL.<br/>**Default:** _6.2.60202-1_ |
| `amdgpu_install_pkg` | Override for the `amdgpu-install` package. URL or file path _(not copied)_.<br/>**Default:** _templated_ from `amdgpu_install_rel` and `amdgpu_install_build` |
| `amdgpu_install_usecases` | Package groups to request from `amdgpu-install` and the distribution package manager.<br/>**Default:** `dkms,graphics,multimedia` _(kernel driver and Mesa)_<br/>**Reference:** see `amdgpu-install --list-usecase` |
| `amdgpu_install_args` | Optional dictionary of arguments to pass to `amdgpu-install`.<br/>**Default:** _Skipped_ |
| `amdgpu_install_branch` | Optional development branch of the `amdgpu` driver with `amdgpu-install-internal`<br/>**Default:** _Skipped_ |
| `amdgpu_install_rocm_branch` | Optional development branch for `rocm` software with `amdgpu-install-internal`<br/>**Default:** _Skipped_ |
| `packer_binary` | The name _or_ path for the _Packer_ binary. Installation skipped when changed.<br/>**Default:** `/usr/bin/packer` |
| `qemu_binary` | The name _or_ path for the _QEMU_ binary.<br/>**Default:** `qemu-system-x86_64` |
| `rocm_kernel` | The kernel package with an optional release specifier.<br/>**Default:** `linux-image-generic-hwe-22.04` |
| `rocm_extras` | Packages installed before `amdgpu-install` _'usecases'_, comma-separated string with optional releases.<br/>**Default:** _linux-headers-generic-hwe-22.04,linux-image-extra-virtual-hwe-22.04_ |
| `rocm_filename` | The name of the output file/artifact _(tarball)_<br/>**Default:** `ubuntu-rocm.tar.gz` |
| `rocm_builder_cpus` | Number of virtual CPUs given to the builder VM.<br/>**Default:** _4_ |
| `rocm_builder_disk` | Space given to the builder; releases compound quickly.<br/>**Default:** _60G_ |
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
