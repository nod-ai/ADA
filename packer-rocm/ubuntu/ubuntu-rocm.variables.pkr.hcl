variable "hidden" {
  type = bool
  default = true
  description = "If the VNC display for the Virtual Machine is hidden while building the image. Adds display requirements"
}

variable "qemu_binary" {
  type = string
  default = "qemu-system-x86_64"
  description = "The name _or_ path for the QEMU binary."
}

variable "ubuntu_release" {
  type = string
  default = "22.04.5"
}

variable "rocm_filename" {
  type        = string
  default     = "ubuntu-rocm.tar.gz"
  description = "The name of the output file/artifact (tarball)"
}

variable "rocm_amdgpu_pkgs" {
  type = string
  default = "amdgpu-dkms"
  description = "Comma-separated string of 'amdgpu' driver/firmware packages to install. May specify releases."
}

variable "rocm_kernel" {
  type = string
  default = "linux-image-generic-hwe-22.04"
  description = "The kernel to include with the image. May include version specifier. Software will be compiled against this; define headers/extra-modules/others in 'rocm_extras'"
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
  default = "linux-headers-generic-hwe-22.04,linux-image-extra-virtual-hwe-22.04,mesa-amdgpu-va-drivers"
  description = "Comma-separated string of extra packages to install [before 'amdgpu-dkms' and ROCm releases]. For headers, extra-modules, and any other packages. May include release specifiers, '=1.2.3' or globbed."
}

variable "rocm_repos" {
  type = string
  default = "true"
  description = "If the 'rocm' and 'amdgpu' repositories are created by the 'rocm' role. Used to opt out when overrides are in the 'repositories' directory"
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
  default = "https://docs.broadcom.com/docs-and-downloads/ethernet-network-adapters/NXE/Thor2/GCA2/bcm5760x_231.2.63.0a.zip"
  description = "The URL for the `niccli` archive."
}

variable "niccli_sum" {
  type = string
  default = "sha256:5c46de9addf9284fb48fef1c505c470c85fd4c129045bdd8ee706447bc1bd025"
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
