niccli
=========

Installs _Broadcom_ `niccli`, a utility for configuring and flashing NICs. Includes the `bnxt_{en,re}` drivers _(optional)_.

Requirements
------------

On the _Ansible Controller_: network access to the web server hosting `niccli_url`.

Role Variables
--------------

| Variable | Description |
|-----|-----|
| `niccli_url` | Required. The download URL for the `niccli` package from Broadcom.<br/>Changes too regularly for templating version-strings. Find [updates here](https://www.broadcom.com/support/download-search?pg=&pf=Ethernet+Connectivity,+Switching,+and+PHYs&pn=&pa=&po=&dk=niccli&pl=&l=false), _'current'_ is in [defaults](./defaults/main.yml). |
| `niccli_sum` | Optional, string. Checksum to use for validating `niccli_url`. Example: `sha256:abcd1234` |
| `niccli_wanted` | Optional, boolean. If `niccli` should be installed. Default `True`. Intended for `host_vars` or `group_vars`. |
| `niccli_driver` | Optional, boolean. If the _'bnxt'_ drivers should be installed through OS-relevant packages and DKMS. Default `True` |
