# ADA
AMD Deployment Assistant


## Build a deployable image with [Packer](https://github.com/nod-ai/ADA/tree/main/packer-rocm) 

## Deploy the image with MAAS

### If you use NFS please update the /opt/rocm symlink to the NFS mount
```
ln -s /nfsshare/rocm-<version> /opt/rocm
```

## Deploy SLURM

## smci_mi300x_platform_amdgpu_alllogs_collection.py

###  Overview

The script pull the AllLog from the Supermicro H13 platform.  Script can run from Windows or Linux.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
    - If it is installed, then please type quit to exit python3.exe.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3.exe smci_mi300x_platform_amdgpu_alllogs_collection.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
2.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_mi300x_platform_amdgpu_alllogs_collection.py
```

### Example of running on RHEL

1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
2.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_mi300x_platform_amdgpu_alllogs_collection.py
```

### Parameters

| Parameter | Description                                                                          |  
|-----------|--------------------------------------------------------------------------------------|  
| --debug   | Adding the --debug option will provide debug information and show the BMC password.  |  

## smci_mi300x_platform_amdgpu_clear_all_logs.py

###  Overview

The script deletes logs from the MI300X DCGPU assembly followed by a power cycle.  This does not
delete the BMC logs.  Running this script is strongly suggested after updating the BKC (MI300X
firmware) or replace an OAM module on the MI300X assembly.  Script can run from Windows or Linux.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
    - If it is installed, then please type quit to exit python3.exe.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3.exe smci_mi300x_platform_amdgpu_clear_all_logs.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
2.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_mi300x_platform_amdgpu_clear_all_logs.py
```

### Example of running on RHEL

1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
2.	Run the script to collect the logs
    - You will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_mi300x_platform_amdgpu_clear_all_logs.py
```

### Parameters

| Parameter | Description                                                                          |  
|-----------|--------------------------------------------------------------------------------------|  
| --debug   | Adding the --debug option will provide debug information and show the BMC password.  |  

