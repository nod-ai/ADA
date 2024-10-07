variable "ubuntu_release" {
  type = string
  default = "22.04.5"
}

variable "rocm_releases" {
  type = string
  default = "6.2.2"
  description = "Comma-separated string of ROCm release(s) for the image; latest is selected for the 'amdgpu' driver"
}

variable "rocm_extras" {
  type = string
  default = "mesa-amdgpu-va-drivers"
  description = "Comma-separated string of extra packages to install [after 'amdgpu-dkms' and ROCm releases]"
}

packer {
  required_plugins {
    ansible = {
      version = "~> 1"
      source = "github.com/hashicorp/ansible"
    }
  }
}
