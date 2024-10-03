source "qemu" "rocm" {
  # install media
  iso_checksum    = "file:http://releases.ubuntu.com/${var.ubuntu_release}/SHA256SUMS"
  iso_target_path = "packer_cache/ubuntu-${var.ubuntu_release}.iso"
  iso_url         = "https://releases.ubuntu.com/${var.ubuntu_release}/ubuntu-${var.ubuntu_release}-live-server-amd64.iso"
  # cloud-init, note 'cd_files' preserves original names
  cd_files        = ["meta-data"]
  cd_content      = {
    "user-data" = file("${path.root}/user-data-rocm")  # workaround for '-rocm' suffix; cloud-init expects 'user-data'
  }
  cd_label        = "cidata"
  # system resources
  cpus            = 8      # expedite compiling, adjust to machine allowance
  disk_size       = "70G"  # ROCm/amdgpu builds demand a lot of space; layout in 'user-data-rocm' will claim all of this (OotB default is 50%)
  memory          = 4096   # OOM w/ 2G during DKMS builds, 3G *may* suffice
  # image/build prefs
  accelerator     = "kvm"  # or 'none' if KVM is unavailable
  boot_command    = ["<wait5>e<wait2>", "<down><down><down><end><wait>", "<bs><bs><bs><bs><wait>autoinstall ---<wait><f10>"]
  boot_wait       = "5s"
  efi_boot        = true
  efi_drop_efivars = true
  format          = "raw"
  headless        = var.headless
  http_directory  = var.http_directory
  shutdown_command       = "sudo -S shutdown -P now"
  ssh_handshake_attempts = 500
  ssh_username           = "ubuntu"
  ssh_password           = var.ssh_ubuntu_password
  ssh_wait_timeout       = var.timeout
  ssh_timeout            = var.timeout
  # debug/discard
  # ssh_pty = true
  # ssh_agent_auth = false
}

build {
  sources = ["source.qemu.rocm"]

  provisioner "file" {
    destination = "/tmp/curtin-hooks"
    source      = "${path.root}/scripts/curtin-hooks"
  }

  provisioner "shell" {
    environment_vars  = ["HOME_DIR=/home/ubuntu", "http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    execute_command   = "{{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
    expect_disconnect = true
    scripts           = ["${path.root}/scripts/curtin.sh", "${path.root}/scripts/networking.sh"]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/rocm.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [  
      "-e", "ansible_python_interpreter=/usr/bin/python3",  # work around Packer/SSH proxy+client limitations
      "--scp-extra-args", "'-O'",
      "-e", "rocm_release_build=${var.rocm_release_build}",  # pass ROCm requests
      "-e", "rocm_release=${var.rocm_release}",
      "-e", "amdgpu_install=${join(",", var.amdgpu_install)}"
    ]
  }

  provisioner "shell" {
    execute_command   = "{{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
    scripts           = ["${path.root}/scripts/cleanup.sh"]
  }

  post-processor "compress" {
    output = "ubuntu-rocm.dd.gz"
  }
}
