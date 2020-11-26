![decipiat logo](https://i.imgur.com/zQYA4fD.png)

`decipiat` is the offensive Phishing-as-a-Service environment.

**Usage**
---

```
Usage: ./decipiat up [OPTIONS]

  Offensive Phishing-as-a-Service environment.
  Developed by: Erhad Husovic -> (Github: XdaemonX)
		Halis Duraki  -> (Github: duraki)


Options:
  -host	     Set hostname
  -port      Set port
  -usessl    Append SSL/TLS config
  -v         Display version and author information and exit.
  -help, -h  Show this message and exit.
```

**Installation Options**
---

1. Install with [`go`](https://golang.org/doc/install)
    + `$ go install decipiat`
    + `$ decipiat up`

2. Download the `decipiat` binary from Releases tab.


**Configuration Options**
---

1. Add more info 

    + Turn on Firewall?
        - This helps protect your Mac from being attacked over the internet.
    + Turn on logging?
        - If there IS an infection, logs are useful for determining the source.
    + Turn on stealth mode?
        - Your Mac will not respond to ICMP ping requests or connection attempts from closed TCP and UDP networks.

**How to Contribute**
---

1. Clone repo and create a new branch: `$ git checkout https://github.com/duraki/decipiat -b name_for_new_branch`.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes

**Project Structure**
---

```
galaxy@devil. decipiat master × tree
.
├── README.md // This file
├── go.mod
├── main.go
├── res/ -> resource dir
├── integration/ -> integration dir (lateraus, modlishka ...)
├── core/ -> anything related to kernel/core
└── web/ -> backend and frontend
```

**Acknowledgements**
---

Just the two of us :hearts:
