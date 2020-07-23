# CENTREON CTL
## Config File
Create a file called `config.yaml` in the root of the project 
<br/>For declare a server use the configuration below by replacing capitalized words

* `NAMESERVER` it's the name of the server that you will use in the flag --server. <span style="color: #FF0000"> WARNING </span> : The name of the server must be unique
* `URL` it's the url at use for access to your centreon server. Its format is **https://monserveurcentreon.fr/centreon** and you replace **monserveurcentreon.fr** by the server. <span style="color: #FF0000"> WARNING </span> : remember to put **http** if your url is not a secure url
* `LOGIN` it's the login that you use for the connection at the server
* `PASSWORD` it's the password that you use for the connection at the server

```yaml
servers:
   - server: "NAMESERVER"
     url: "URL" 
     login: "LOGIN"
     password: "PASSWORD"
   - server: "NAMESERVER2"
     url: "URL" 
     login: "LOGIN"
     password: "PASSWORD"
```

## CSV File
Create a file with the extension `.csv` in the root of the project
<br/>You can create contact, host and service. The first column is to specified the object. The syntax is :

```csv
contact,full name,login,email,password,admin
host,name,alias,IPaddress,templateHost,pollerName,hostGroup
service,hostName,description,templateService
```

The following table shows you an example for each case:

|            Example             |                         Syntaxe in the CSV                              |
| :----------------------------- | :---------------------------------------------------------------------  |
| A service                      | service,HostTest,ServiceDescription,TemplateService                     |
| A host with hostGroup          | host,HostTest,HostTest,127.0.0.1,TemplateHost,Central,group_test        |
| A host without hostGroup       | host,HostTest2,HostTest2,127.0.0.1,TemplateHost,Central,                |
| A contact with option admin    | contact,Pierre Allain,pierre,pierreallain@mail.com,pierrepassword,admin |
| A contact without option admin | contact,Paul Allain,paul,paulallain@mail.com,paulpassword,              |

<span style="color: #FF0000"> WARNING </span>: when a host has not host group and when a contact is not admin, think to put the comma in the end of line. 
You can write comments in your csv file using a # at the beginning of the line. Example: `#It's a comment and this line is not analysed by the program`<br/>
For export the CSV file use the command next : `centctl export -f nameFile.csv --server nameServer` 
<br/>When exporting all objects is complete, the `nameFile.csv` is deleted and a file named `donenameFile.csv` is created and contains all object created. Otherwise, an error which specifie the problem appear in the terminal.

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