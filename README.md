# unikernel-function-runtime

## Manual usage of unikernels
    Below are detailed some basic steps needed to build your own unikernels.

    Next obvious step will be to build an agent to automate below tasks while
    said agent will be agnostic regarding the unikernel of choice.

### Clone & build Firecracker
> git clone https://github.com/firecracker-microvm/firecracker.git

> cd firecracker

> ./tools/devtool build

* Executable binary in `build/cargo_target/${toolchain}/debug/firecracker`

### OSv Clone & Build Kernel
TODO : Automate steps below
* Initial steps
> git clone https://github.com/cloudius-systems/osv.git

> cd osv && git submodule update --init --recursive

* Install dependencies
> ./scripts/setup.py

* Build default image
> ./scripts/build

### OSv Build & Run custom images
* Modify config.json => `osv/config.json`
    
    ! For this project you should use `<THIS_REPO_CLONE_DIR>/images/osv` 
```json
    {
    "modules": {
            "osvinit": {
                "type": "direct-dir",
                "path": "${OSV_BASE}/modules/cloud-init/"
            },
        "repositories": [
            "${OSV_BASE}/apps",
            "${OSV_BASE}/modules",
            "< !!! PATH TO YOUR IMAGES DIRECTORY HERE !!! >" 
        ]
    },
    "default": [ "cloud-init", "httpserver" ]
    }
```
* Create app directory in images directory from above step
* Create required files under app directory (see `images/osv/node_bm` example)
```
    ROOTFS // place relevant files and directories here
    modules.py // specify requirements and command to be executed
    usr.manifest // map directories
```
* Use OSv Firecracker script at `./scripts/firecracker.py` (automatically runs last built image) or run Firecracker yourself
```
    kernel image path : osv/build/last/loader-stripped.elf
    rootfs image path : osv/build/last/usr.img
```
TODO: Explain `manifest_from_host.sh` script => TLDR: uses runtimes installed on host system to build images
### NanOS Clone & Build
TODO: Automate unikernel creation steps && write steps in more detail
* Initial steps
> git clone https://github.com/nanovms/nanos.git

> Grab dependencies, see https://github.com/nanovms/nanos
> Grab `ops` tool, see https://github.com/nanovms/ops

* First build 
> make run-noaccel

### NanOS Building custom images

* Create directory for your app and `cd` to said directory
* Create `config.json`

Example config files 

(nanos/test/e2e/ruby_2.5.1/config.json)
```json
{
    "Args": ["myapp.rb", "-o", "0.0.0.0"],
    "Dirs": [".ruby"],
    "ENV": {
        "GEM_HOME": ".ruby"
    },
    "RunConfig": {
        "Ports": ["4567"]
    },
    "Boot": "../../../output/test/e2e/boot.img",
    "Kernel": "../../../output/test/e2e/kernel.img"
}

```
(nanos/test/e2e/node_v11.5.0/config.json)
```json
{
    "Args": [
        "hello.js"
    ],
    "RunConfig": {
        "Ports": [
            "12345"
        ]
    },
    "Program": "/usr/bin/node",
    "Files": [
        "hello.js"
    ]
}
```

* Build image using ops tool and give it a name
    `ops image create -c config.json -i <IMAGE_NAME>`

* Find image under `$HOME/.ops/images/<IMAGE_NAME>`

## Tested languages/runtimes
Note : relevance of below measurements is uncertain at best since I had no native linux machine to test on.
       tests were run on WSL2 with Ubuntu 20.04.4 LTS, and further test runs on more appropiate environments needed

Note : above mentioned WSL2 instance lives on a hard disk

Note : 5 runs, ignored first run due to caching reasons

Avg. Firecracker main() timestamp <-> Avg. Firecracker start VCPU 
Avg. Firecracker start VCPU  <-> Avg. Start process 
Note : <-> represents time difference
### Node.js
* NanOS

|              | FC_MAIN > START VCPU (ms) | START VCPU > EXEC ENTRY POINT (ms) |
|--------------|---------------------------|------------------------------------|
|              | 3                         | 33.6                               |
|              | 4                         | 30.2                               |
|              | 6                         | 38.2                               |
|              | 6                         | 33                                 |
|              | 5                         | 33.9                               |
| Average (ms) | 4.8                       | 33.78                              |

* OSv

> N/A

### Kotlin
* NanOS

|              | FC_MAIN > START VCPU (ms) | START VCPU > EXEC ENTRY POINT (ms) |
|--------------|---------------------------|------------------------------------|
|              | 5                         | 33.6                               |
|              | 6                         | 34.8                               |
|              | 5                         | 34.6                               |
|              | 5                         | 34.2                               |
|              | 5                         | 34                                 |
| Average (ms) | 5.2                       | 34.24                              |

* OSv

> N/A
### Swift
* NanOS

|              | FC_MAIN > START VCPU (ms) | START VCPU > EXEC ENTRY POINT (ms) |
|--------------|---------------------------|------------------------------------|
|              | 6                         | 30.5                               |
|              | 5                         | 32.1                               |
|              | 5                         | 31.6                               |
|              | 6                         | 30.7                               |
|              | 5                         | 32                                 |
| Average (ms) | 5.4                       | 31.38                              |

* OSv

> N/A


