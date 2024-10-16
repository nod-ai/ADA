variable "ubuntu_release" {
  type = string
  default = "22.04.5"
}

variable "rocm_filename" {
  type        = string
  default     = "ubuntu-rocm.tar.gz"
  description = "The name of the output file/artifact (tarball)"
}

variable "rocm_kernel" {
  type = string
  default = "linux-image-generic-hwe-22.04=5.15.*"
  description = "The kernel to include with the image. May include version specifier. Software will be compiled against this; define headers/others in 'rocm_extras'"
}

variable "rocm_releases" {
  type = string
  default = "6.2.2"
  description = "Comma-separated string of ROCm release(s) for the image; latest is selected for the 'amdgpu' driver"
}

variable "rocm_installed" {
  type = string
  default = "false"
  description = "If ROCm packages should be installed after 'amdgpu-dkms'; accepts 'truthy' or 'falsy' values (1/0/y/n/'inapplicable')"
}

variable "rocm_extras" {
  type = string
  default = "mesa-amdgpu-va-drivers"
  description = "Comma-separated string of extra packages to install [after 'amdgpu-dkms' and ROCm releases]"
}

variable "rocm_builder_cpus" {
  type = number
  default = 4
  description = "Number of CPU threads given to the builder Virtual Machine. More may help compilation speed"
}

variable "rocm_builder_memory" {
  type = number
  default = 4096
  description = "RAM given to the builder VM, measured in MB. Out-of-memory conditions were found with 2G during DKMS builds"
}

variable "rocm_builder_disk" {
  type = string
  default = "70G"
  description = "amdgpu and ROCm releases demand considerable space. Layout in 'user-data-rocm' will claim all of this"
}

variable "niccli_wanted" {
  type = string
  default = "true"
  description = "If 'niccli' (Broadcom) is included in the image"
}

variable "niccli_url" {
  type = string
  default = "https://docs.broadcom.com/docs-and-downloads/ethernet-network-adapters/NXE/Thor2/GCA1/bcm5760x_230.2.52.0a.zip"
  description = "The URL for the `niccli` archive."
}

variable "niccli_sum" {
  type = string
  default = "sha256:1dd2a7c6978c978febfc08ef2f416b2745573848d45700737d4b1405d8e1ac35"
  description = "Checksum for validating `niccli_url`."
}

variable "niccli_driver" {
  type = string
  default = "true"
  description = "If the `bnxt_{en,re}` NIC drivers are included."
}

packer {
  required_plugins {
    ansible = {
      version = "~> 1"
      source = "github.com/hashicorp/ansible"
    }
  }
}
