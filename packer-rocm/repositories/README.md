# repositories

Any `.repo` or `.list` files placed here will be automatically copied into the _Packer_ builder VM.

Use `rocm_extras` to request any packages they may provide.

When replacing the `amdgpu` or `rocm` repositories, `-e rocm_repos=false` may skip those for generally-available releases.
