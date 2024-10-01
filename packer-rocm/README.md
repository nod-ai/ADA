# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with [ROCm](https://www.amd.com/en/products/software/rocm.html) installed.
Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.

## One-time Setup

Please see [installation instructions](https://developer.hashicorp.com/packer/docs/install)
for the `packer` CLI.

1. Clone repositories:

    ```shell
    git clone git@github.com:canonical/packer-maas.git
    git clone git@github.com:nod-ai/ADA.git
    ```

2. Copy assets from _ADA_ `packer-rocm` to the _Canonical_ `packer-maas` source:

    ```shell
    # Repeat '--exclude' with shell expansion, slashes are significant for 'rsync'
    rsync -avP --exclude={'*.md','LICENSE','NOTICE'} ADA/packer-rocm/ packer-maas/
    ```

3. Install plugins:

    ```shell
    cd packer-maas/ubuntu
    packer init .
    ```

## Building

```shell
PACKER_LOG=1 packer build -only=qemu.rocm .
```
