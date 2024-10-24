# repositories

Any `.repo` or `.list` files placed here will be automatically copied into the _Packer_ builder VM.

Use `-e rocm_extras=pkg1,pkg2,...` to request any packages they -- or the distribution -- provide.

When replacing the `amdgpu` or `rocm` repositories, `-e rocm_repos=false` may skip those for generally-available releases.
