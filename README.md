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

1. Add more info 

    + Turn on Firewall?
        - This helps protect your Mac from being attacked over the internet.
    + Turn on logging?
        - If there IS an infection, logs are useful for determining the source.
    + Turn on stealth mode?
        - Your Mac will not respond to ICMP ping requests or connection attempts from closed TCP and UDP networks.

**Configuration Options**
---

**How to Contribute**
---

1. Clone repo and create a new branch: `$ git checkout https://github.com/duraki/decipiat -b name_for_new_branch`.
2. Make changes and test
3. Submit Pull Request with comprehensive description of changes

**Development Environment**

Instructions for MacOS Catalina environment:

```
### MongoDB => https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/
$ brew tap mongodb/brew
$ brew install mongodb-community@4.4

$ mkdir -p /data/db
$ sudo chown -R `id -un` /data/db

-- Service env (startup):
$ brew services start mongodb-community@4.4
$ brew services stop mongodb-community@4.4

-- Config env (once):
$ mongod --config /usr/local/etc/mongod.conf

$ mongo       # to connect via cli
$ mongotop    # db tools/utils

-- tldr: mongo
> show dbs
admin     0.000GB
config    0.000GB
decipiat  0.000GB
local     0.000GB
> use decipiat
switched to db decipiat
> show collections
users
> db.users.find()
  ...
```

---

**Project Structure**

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
