# repositories

Any `.repo` or `.list` files placed here will be automatically copied into the _Packer_ builder VM.

Use `rocm_extras` to define any packages they may provide. If replacing the `amdgpu` or `rocm` repositories, `-e rocm_repos=false` is available.
