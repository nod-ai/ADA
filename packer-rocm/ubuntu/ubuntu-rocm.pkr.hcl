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
  cpus            = var.rocm_builder_cpus
  disk_size       = "${var.rocm_builder_disk}"  # Packer seems to have trouble with strings that begin with numbers; explicitly cast
  memory          = var.rocm_builder_memory
  # image/build prefs
  accelerator     = "kvm"  # or 'none' if KVM is unavailable
  boot_command    = ["<wait5>e<wait2>", "<down><down><down><end><wait>", "<bs><bs><bs><bs><wait>autoinstall ---<wait><f10>"]
  boot_wait       = "5s"
  efi_boot        = true
  efi_drop_efivars = true  # don't place efivars.fd in output artifact
  format          = "raw"  # qcow2 may not be converted. if written to drives, can't be read back/won't find 'curtin'
  headless        = var.headless
  http_directory  = var.http_directory
  shutdown_command       = "sudo -S shutdown -P now"
  ssh_handshake_attempts = 500
  ssh_username           = "ubuntu"
  ssh_password           = var.ssh_ubuntu_password
  ssh_wait_timeout       = "1h"
  ssh_timeout            = "1h"
  # debug/discard
  # ssh_pty = true
  # ssh_agent_auth = false
}

build {
  sources = ["source.qemu.rocm"]

  # generate/copy tarball of custom packages; 'packer-maas' will process
  provisioner "shell-local" {
    inline = [
      "tar cvzf ${path.root}/custom-packages.tar.gz -C ${path.root}/packages --overwrite .",
    ]
    inline_shebang = "/bin/bash -e"
  }
  provisioner "file" {
    destination = "/tmp/"
    generated   = true
    sources     = [
      "${path.root}/custom-packages.tar.gz"
    ]
  }

  # copy supporting 'packer-maas' assets
  provisioner "file" {
    destination = "/tmp/"
    sources     = [
      "${path.root}/scripts/curtin-hooks",
      "${path.root}/scripts/setup-bootloader",
      "${path.root}/scripts/install-custom-packages"
    ]
  }

  # remove step/message from 'install-custom-packages' RE: uninstalling existing kernels; DKMS/'cloud-init'
  provisioner "shell" {
    inline_shebang = "/bin/bash"
    inline = [
      "sed -i -e '/remove existing kernels/d' -e '/xargs apt-get -y purge/d' /tmp/install-custom-packages",
    ]
  }

  provisioner "shell" {
    environment_vars  = ["HOME_DIR=/home/ubuntu", "http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}", "CLOUDIMG_CUSTOM_KERNEL=${var.kernel}", "DEBIAN_FRONTEND=noninteractive"]
    execute_command   = "{{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
    expect_disconnect = true
    scripts           = ["${path.root}/scripts/curtin.sh", "${path.root}/scripts/networking.sh", "${path.root}/scripts/cloudimg/install-custom-kernel.sh"]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/os_prep.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [
      "-e", "ansible_python_interpreter=/usr/bin/python3",  # work around Packer/SSH proxy+client limitations
      "--scp-extra-args", "'-O'"
    ]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/rocm.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [
      "-e", "ansible_python_interpreter=/usr/bin/python3",
      "--scp-extra-args", "'-O'",
      "-e", "rocm_releases=${var.rocm_releases}",  # pass ROCm requests [release + packages]
      "-e", "rocm_extras=${var.rocm_extras}",
      "-e", "rocm_installed=${var.rocm_installed}"
    ]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/tuned.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [
      "-e", "ansible_python_interpreter=/usr/bin/python3",
      "--scp-extra-args", "'-O'"
    ]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/niccli.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [
      "-e", "ansible_python_interpreter=/usr/bin/python3",
      "--scp-extra-args", "'-O'",
      "-e", "niccli_wanted=${var.niccli_wanted}",
      "-e", "niccli_url=${var.niccli_url}",
      "-e", "niccli_sum=${var.niccli_sum}",
      "-e", "niccli_driver=${var.niccli_driver}"
    ]
  }

  provisioner "ansible" {
    playbook_file = "${path.root}/../playbooks/tuned.yml"
    user          = "ubuntu"
    ansible_env_vars  = ["http_proxy=${var.http_proxy}", "https_proxy=${var.https_proxy}", "no_proxy=${var.no_proxy}"]
    extra_arguments = [  
      "-e", "ansible_python_interpreter=/usr/bin/python3",
      "--scp-extra-args", "'-O'"
    ]
  }

  provisioner "shell" {
    environment_vars = [
      "DEBIAN_FRONTEND=noninteractive"
    ]
    execute_command   = "{{ .Vars }} sudo -S -E sh -eux '{{ .Path }}'"
    inline_shebang = "/bin/bash"  # not giving '-e'; allow clean-up drift - don't fail if a task fails
    inline = [
      "apt-get autoremove --purge -yq",
      "apt-get clean -yq",
      "cloud-init clean --logs --machine-id --configs all",
      "rm -fv /etc/ssh/ssh_host_* /etc/cloud/cloud.cfg.d/{*installer*,20-disable-cc-dpkg-grub}.cfg /etc/cloud/ds-identify.cfg /var/log/{alternatives,bootstrap}.log",
      "find /tmp/niccli -xdev -delete",
      "find /tmp/libbnxt_re -xdev -delete",
      "find /var/log/installer -xdev -delete",
    ]
  }

  post-processor "shell-local" {
    inline = [
      "SOURCE=${source.name}",
      "OUTPUT=${var.rocm_filename}",
      "IMG_FMT=raw",
      "ROOT_PARTITION=2",
      # expedite compression: use an exported function to test for 'pigz' and remap 'gzip' to it, if found
      "if command -v pigz > /dev/null 2>&1; then",
      "  echo 'Mapping gzip to pigz for parallel compression'",
      "  gzip() { pigz \"$@\"; }",
      "  export -f gzip",
      "else",
      "  echo '*Not* mapping gzip to pigz for parallel compression'",
      "fi",
      # mount the image, prepare the tarball, compress, and finally clean up temporary i/o
      "source ../scripts/fuse-nbd",
      "source ../scripts/fuse-tar-root",
      "rm -rf output-${source.name} ${path.root}/custom-packages.tar.gz"
    ]
    inline_shebang = "/bin/bash -e"
  }
}
