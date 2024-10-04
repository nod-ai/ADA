# Slurm REST API Generator

The Slurm REST API Generator allows you to easily generate REST APIs based on the OpenAPI specifications provided by your Slurm server.

## Prerequisites

Ensure Docker is running on your machine.

## Getting Started

Follow these steps to generate REST APIs:

1. **Fetch OpenAPI Specification**
Use the following command to retrieve the `slurm_openapi.json` file. And, this file needs to be saved/copied under `api/` directory.
```bash
unset SLURM_JWT; export $(scontrol token)
curl -H "X-SLURM-USER-TOKEN:$SLURM_JWT" -k http://<slurm-rest-server-ip>:<port>/openapi/v3 -o slurm_openapi.json
```

Note: Make sure to replace `<slurm-rest-server-ip>` and `<port>` with the actual IP address and port of your Slurm REST server.

2. **Generate REST APIs**
After obtaining the `slurm_openapi.json`, run the following command to generate the REST APIs for your specific Slurm version:
```bash
cd api; make 
```

Alternatively, the REST APIs are automatically generated during the build process once the slurm_openapi.json file is in place.