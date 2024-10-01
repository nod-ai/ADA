variable "ubuntu_release" {
  type = string
  default = "22.04.5"
}

variable "rocm_release" {
  type = string
  default = "6.2.2"
}

variable "rocm_release_build" {
  type = string
  default = "6.2.60202-1"
}

packer {
  required_plugins {
    ansible = {
      version = "~> 1"
      source = "github.com/hashicorp/ansible"
    }
  }
}
