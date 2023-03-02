# Lemurian Labs device introspection

Golang application to work with a Lemurian Lab perception accelerator device.

It provides a CLI to interact with the device and to query its capabilities and content.


# How to build

There is a Taskfile and a VERSION file that is used to automate building an executable
with a consistent semver and build hash.

```bash
> task
task: Available tasks for this project: 
* compile:      Compile the Lemurian Labs device CLI 
* container:    Create a container with the Lemurian Labs device services 
* install:      Compile and install the Lemurian Labs device CLI 
* test:         Run all the tests
```

# How to run

After building the root directory will contain an 'lld' executable:

```bash
> ./lld 
CLI for the Lemurian Labs device introspection 
 
Usage:
   lld [command]  

Available Commands:   
   completion  Generate the autocompletion script for the specified shell
   help        Help about any command   
   version     display lld CLI version  

Flags:
   -h, --help   help for lld  

Use "lld [command] --help" for more information about a command. 
```
