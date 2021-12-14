# CENTCTL

## Presentation
Centctl is a CLI which allows to manage Centreon servers. Centctl use a file named `centctl.yaml` for manage the connection at differents servers. </br>
Developped in Go, it allows to operate a Centreon platform remotely from a PC under Windows, Linux or Macos without any particular installation.

## Config File
Create a file called `centctl.yaml` in your PC and put the complete path of this file in an environment variable named `CENTCTL_CONF`.<span style="color: #FF0000"> WARNING </span> : This variable is required. For example:

`export CENTCTL_CONF=/home/mister/centctl/centctl.yaml`

<br/>For declare a server use the configuration below by replacing capitalized words

* `NAMESERVER` it's the name of the server that you will use in the flag --server. <span style="color: #FF0000"> WARNING </span> : The name of the server must be unique
* `URL` it's the url at use for access to your centreon server. Its format is **https://monserveurcentreon.fr/centreon** and you replace **monserveurcentreon.fr** by the server. <span style="color: #FF0000"> WARNING </span> : remember to put **http** if your url is not a secure url
* `insecure` key is use when the URL server is in https. It is same the flag (--insecure) to forced the connection.
* `LOGIN` it's the login that you use for the connection at the server
* `PASSWORD` it's the password that you use for the connection at the server
* `VERSION` it's the version of the API that you want to use (possibles values : v1 or v2)
* `default` define the default server to connected 
</br>To define the proxy if your server use it, you have three ways:
  - Define an environment variable named `http_proxy`. Its format is **http://IPAddress:Port** or **http://USERNAME:Password@IPAddress:Port** (In this case all servers use the same proxy) or an environment variable named `https_proxy. Its format is the same of http_proxy
  - Define proxy in the top of the yaml file (In this case all servers use the same proxy)
  ```yaml
  proxy:
    - httpURL: "URLProxy"
    - httpsURL: "URLSProxy"
    - user: "USERNAME"
    - password: "PASSWORD" 
  servers:
    - server: "NAMESERVER2"
      url: "URL" 
      login: "LOGIN"
      password: "PASSWORD"
      version: "VERSION"
  ``` 
  - Define differents proxies for each server which used it like below.
</br>It is possible to have servers without a proxy like the NAMESERVER2 below. It is possible to have servers without httpProxy or without httpsProxy, for this case keep the field and use blank value (example: - httpsURL: "").
* `URLProxy` it's the ipAddress and port at use for access to your http proxy. Its format is **IPAddress:Port**
* `URLSProxy` it's the ipAddress and port at use for access to your https proxy. Its format is **IPAddress:Port**
* `USERNAME` it's the user that you use for connection at the proxy (If you don't use it, leave the field blank )
* `PASSWORD` it's the password that you use for connection at the proxy (If you don't use it, leave the field blank )

```yaml
servers:
   - server: "NAMESERVER"
     url: "URL" 
     insecure: true
     login: "LOGIN"
     password: "PASSWORD"
     version: "VERSION"
     default: true
     proxy:
      - httpURL: "URLProxy"
      - httpsURL: "URLSProxy"
      - user: "USERNAME"
      - password: "PASSWORD" 
   - server: "NAMESERVER2"
     url: "URL" 
     login: "LOGIN"
     password: "PASSWORD"
     version: "VERSION"

```

## To manage passwords
To manage passwords and for don't have the password written clearly in config file you have two possibilities.
### First solution
The first solution consists in don't write the password into the yaml file and the CLI ask you to give the password in the terminal before a command.
<br/> In this case, your yaml file look like this (the parameter password no longer appears):

```yaml
servers:
   - server: "NAMESERVER"
     url: "URL" 
     login: "LOGIN"
     version: "VERSION"
     default: true
   - server: "NAMESERVER2"
     url: "URL" 
     login: "LOGIN"
     version: "VERSION"
```

### Second solution 
The second solution consist in encrypt the passwords present in the config file. 
<br/> To encrypt the password enter this command:
```sh
centctl encrypt
``` 

When this command succeed, the passwords in file are encrypt and the result of this command give you the `encryption key`. To decode your passwords after that, centctl **NEEDS** this key, consequently you **MUST** save this key in an environment variable named `CENTCTL_DECRYPT_KEY`. 
```sh
export CENTCTL_DECRYPT_KEY="THE_KEY_RETURNED_BY_THE_COMMAND"
``` 

If you modify passwords or add new server in config file, when you use the command `encrypt` the key used is the same that the first time.

If you loose this key, you must rewrite the passwords in the config file and reexecute the above command.

## CSV File
Create a file with the extension `.csv`.
<br/>You can create and modify all objects available in centctl. You can find some examples in the examples folder.
Each example can be executed directly with command : 
```sh
centctl import -f pathOfFileCSV --server nameSERVER
``` 

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
centctl completion [bash|zsh|fish|powershell] >/etc/bash_completion.d/centctl
``` 

After reloading your shell, centctl autocompletion should be working.

Note :
Centreon, the Centreon Logo, are trademarks, servicemarks, registered trademarks or registered servicemarks owned by Centreon Software. All other trademarks, servicemarks, registered trademarks, and registered servicemarks are the property of their respective owner(s). CentCtl is not endorsed nore supported by Centreon Software and use only external API based on Centreon Plateform. 
