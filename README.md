# ADA
AMD Deployment Assistant


## Build a deployable image with [Packer](https://github.com/nod-ai/ADA/tree/main/packer-rocm) 

## Deploy the image with MAAS

### If you use NFS please update the /opt/rocm symlink to the NFS mount
```
ln -s /nfsshare/rocm-<version> /opt/rocm
```


## Deploy SLURM
