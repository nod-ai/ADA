variable "hidden" {
  type = bool
  default = true
  description = "If the VNC display for the Virtual Machine is hidden while building the image. Adds display requirements"
}

variable "amdgpu_install_rel" {
  type = string
  default = "6.2.2"
  description = "The _release_ portion of the `amdgpu-install` URL. May be overridden by 'amdgpu_install_pkg'"
}

variable "amdgpu_install_build" {
  type = string
  default = "6.2.60202-1"
  description = "The _build_ portion of the `amdgpu-install` URL. May be overridden by 'amdgpu_install_pkg'"
}

variable "amdgpu_install_pkg" {
  type = string
  default = "templated"
  description = "Optional URL for the `amdgpu-install` package. Typically determined by Ansible Facts; declaration replaces `amdgpu_install_rel` and `amdgpu_install_build`."
}

variable "amdgpu_install_usecases" {
  type = string
  default = "dkms,graphics,multimedia"
  description = "Package groups to request from `amdgpu-install` that match typical workflows and runtimes."
}

variable "amdgpu_install_branch" {
  type = string
  default = "2065452"  # passing from Ansible -> Packer -> Ansible may complicate type; simplify w/ str
  description = "Placeholder. Optional development branch for the `amdgpu` driver with `amdgpu-install-internal` (through `amdgpu_install_pkg`)"
}

variable "amdgpu_install_rocm_branch" {
  type = string
  default = "compute-rocm-dkms-no-npi-hipclang/14986"
  description = "Placeholder. Optional development branch for `rocm` software with `amdgpu-install-internal` (through `amdgpu_install_pkg`)"
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

variable "rocm_kernel" {
  type = string
  default = "linux-image-generic-hwe-22.04"
  description = "The kernel to include with the image. May include version specifier. Software will be compiled against this; define headers/extra-modules/others in 'rocm_extras'"
}

variable "rocm_extras" {
  type = string
  default = "linux-headers-generic-hwe-22.04,linux-image-extra-virtual-hwe-22.04"
  description = "Comma-separated string of extra packages to install [before 'amdgpu-dkms' and ROCm releases]. For headers, extra-modules, and any other packages. May include release specifiers, '=1.2.3' or globbed."
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
