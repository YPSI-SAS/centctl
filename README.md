# CENTCTL

## Presentation
Centctl is a CLI which allows to manage Centreon servers. Centctl use a file named `config.yaml` for manage the connection at differents servers. </br>
Developped in Go, it allows to operate a Centreon platform remotely from a PC under Windows, Linux or Macos without any particular installation.

## Config File
Create a file called `config.yaml` in your PC and put the path of this file in an environment variable named `CENTCTL_CONF`.<span style="color: #FF0000"> WARNING </span> : This variable is required.
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

```sh
source /usr/share/bash-completion/bash_completion
```

Reload your shell and verify that bash-completion is correctly installed by typing `type _init_completion`

### Enable centctl autocompletion
Add the completion script to the `/etc/bash_completion.d` directory:

```sh
centctl completion >/etc/bash_completion.d/centctl
``` 

After reloading your shell, centctl autocompletion should be working.

Note :
Centreon, the Centreon Logo, are trademarks, servicemarks, registered trademarks or registered servicemarks owned by Centreon Software. All other trademarks, servicemarks, registered trademarks, and registered servicemarks are the property of their respective owner(s). CentCtl is not endorsed nore supported by Centreon Software and use only external API based on Centreon Plateform. 