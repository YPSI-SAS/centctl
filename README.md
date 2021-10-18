![](asserts/img/LOGO.png)

## Overview
Centctl is a CLI which allows to manage Centreon servers. Centctl use a file named `centctl.yaml` for manage the connection at differents servers.

Developped in [Go](https://golang.org/), it allows to operate a [Centreon](https://www.centreon.com/) platform remotely from a PC under Windows, Linux or Macos without any particular installation.

## Installation

Get source code:
```bash
git clone https://github.com/YPSI-SAS/centctl.git
```
### Docker
Build docker container:
```bash
make docker
```
Add alias to your *bash* configuration:
```bash
echo "alias centctl=docker run --rm -i centctl" >> ~/.bashrc
source ~/.bashrc
```
And run `centctl` command.

### GNU/Linux
Build from source:
```bash
make linux
```
Copy binary to `/usr/local/bin`:
```bash
sudo cp centctl_linux_amd64 /usr/local/bin/centctl
sudo chmod a+x /usr/local/bin/centctl
```
And run `centctl` command.
## Usage
### Config File
Create a file called `centctl.yaml` in your PC and put the complete path of this file in an environment variable named `CENTCTL_CONF`.<span style="color: #FF0000"> WARNING </span> : This variable is required. For example:

```bash
export CENTCTL_CONF=/home/mister/centctl/centctl.yaml
```

<br/>For declare a server use the configuration below by replacing capitalized words

* `NAMESERVER` it's the name of the server that you will use in the flag --server. <span style="color: #FF0000"> WARNING </span> : The name of the server must be unique
* `URL` it's the url at use for access to your centreon server. Its format is **https://monserveurcentreon.fr/centreon** and you replace **monserveurcentreon.fr** by the server. <span style="color: #FF0000"> WARNING </span> : remember to put **http** if your url is not a secure url
* `LOGIN` it's the login that you use for the connection at the server
* `PASSWORD` it's the password that you use for the connection at the server
* `VERSION` it's the version of the API that you want to use (possibles values : v1 or v2)

```yaml
servers:
   - server: "NAMESERVER"
     url: "URL" 
     login: "LOGIN"
     password: "PASSWORD"
     version: "VERSION"
   - server: "NAMESERVER2"
     url: "URL" 
     login: "LOGIN"
     password: "PASSWORD"
     version: "VERSION"

```

## CSV File
Create a file with the extension `.csv`.
<br/>You can create and modify all objects available in centctl. You can find some examples in the examples folder.
Each example can be executed directly with command : `centctl import -f pathOfFileCSV --server nameSERVER`

## Enabling shell autocompletion
### Install bash-completion
You can install it with `apt-get install bash-completion`

The above commands create `/usr/share/bash-completion/bash_completion`, which is the main script of bash-completion. Depending on your package manager, you have to manually source this file in your `~/.bashrc file`.

To find out, reload your shell with command `source ~/.bashrc` and run `type _init_completion`. If the command succeeds, you’re already set, otherwise add the following to your `~/.bashrc` file:

```bash
source /usr/share/bash-completion/bash_completion
```

Reload your shell and verify that bash-completion is correctly installed by typing `type _init_completion`

### Enable centctl autocompletion
Add the completion script to the `/etc/bash_completion.d` directory:

```bash
centctl completion >/etc/bash_completion.d/centctl
``` 

After reloading your shell, centctl autocompletion should be working.

## Developer

### Environment setup
#### Prerequisites via apt
Due to dependencies (for example **make** and *build-essential* package):
```bash
sudo apt update
sudo apt install build-essential golang-go docker-ce -y
```

#### Common environment setup
1. Clone this repository `git clone https://github.com/YPSI-SAS/centctl.git`
2. Change directory into the repository root dir: `cd centctl`
3. Get all *go* dependencies with `go mod tidy` && `go mod download`
4. Use `make` or `make help` command to get more informations about **Makefile** targets.


#### Build Centctl
Build for a specific platform:
```bash
# Build for linux
make linux

# Build for Windows
make windows

# Build for MacOS
make darwin
```