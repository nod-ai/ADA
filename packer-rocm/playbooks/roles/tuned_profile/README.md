tuned\_profile
=========

Role to manage [tuned](http://www.tuned-project.org/), profiles, and related Ansible facts.


Role Variables
--------------

| Variable | Description |
|-----|-----|
| `tuned_profile_name` | Name of the custom `tuned` profile to create. |
| `tuned_profile_scripts` | Optional list of script files to copy from the controller/repository, supporting the profile. |
| `tuned_profile_reboot` | If activating the profile requires _(and gets)_ a reboot to become effective. Default is False.<br/>Intended for those with 'cmdline' options. |
| `tuned_profile_params` | Required, dictionary of `tuned.conf` sections with each key being a plugin. |


Preview
--------------

Profile variables _(minified)_, see [our playbook](../../playbook.yml) for the most current example:

```yaml
tuned_profile_name: amd_custom
tuned_profile_scripts:
  - acs_disable.sh
tuned_profile_params:
  main:
    description: "Ansible managed; do not edit. Please 'include'/extend in other profiles instead."
    summary: 'AMD site customizations'
    include: 'network-throughput'  # base profile to build upon
  bootloader:  # one of many plugins; shouldn't be repeated. the 'cmdline_...' keys should be globally unique
    cmdline_amdsre_debug: 'earlyprintk=ttyS1,keep'
    cmdline_amdsre_iommu: 'amd_iommu=on iommu=pt nokaslr'
    cmdline_amdsre_pci: 'pci=realloc=off pcie_aspm=off'
  disableacs:  # scripts may get their own section names to avoid repeat limitations
    type: script
    script: '${i:PROFILE_DIR}/acs_disable.sh'
  sysctl:
    kernel.numa_balancing: 0
    dev.raid.speed_limit_min: 100000
    dev.raid.speed_limit_max: 20000000
```

Rendered:

```ini
$ cat /etc/tuned/profiles/amd_custom/tuned.conf 
#
# Ansible managed
#

[main]
description=Ansible managed; do not edit. Please 'include'/extend in other profiles instead.
summary=AMD site customizations
include=latency-performance

[bootloader]
cmdline_amdsre_debug=earlyprintk=ttyS1,keep
cmdline_amdsre_iommu=amd_iommu=on iommu=pt nokaslr
cmdline_amdsre_pci=pci=realloc=off pcie_aspm=off

[disableacs]
type=script
script=${i:PROFILE_DIR}/acs_disable.sh

[sysctl]
kernel.numa_balancing=0
dev.raid.speed_limit_min=100000
dev.raid.speed_limit_max=20000000

[variables]
include=/etc/tuned/amd-vars.conf
```

