# `dilawar`
> tiny (cli based) finance manager

## Install
```bash
# install with go get and make sure you have $GOPATH/bin in your path 
λ go get github.com/umayr/dilawar/cmd/dilawar

# make with source
λ git clone github.com/umayr/dilawar
λ make build

# create cross platform binaries
λ git clone github.com/umayr/dilawar
λ make cross
```
or you can download the pre-compiled version of binaries from [here](https://github.com/umayr/dilawar/releases/tag/0.1.0). Moreover, If you're using Windows, just run the file `install.bat` as an administrator, and it'll download and install the pre-compiled version to your system on one click.

## Usage

Global Help:
```
λ dilawar --help
NAME:
   dilawar - tiny finance manager

USAGE:
   dilawar [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     debit    adds amount in debit
     credit   adds amount in credit
     history  shows all history
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```
`credit` subcommand help:
```
λ dilawar credit -h
NAME:
   dilawar credit - adds amount in credit

USAGE:
   dilawar credit [command options] [arguments...]

OPTIONS:
   --amount value, -a value   amount to be added (default: 0)
   --message value, -m value  message for the transaction
```
`debit` subcommand help:
```
λ dilawar debit -h
NAME:
   dilawar debit - adds amount in debit

USAGE:
   dilawar debit [command options] [arguments...]

OPTIONS:
   --amount value, -a value   amount to be added (default: 0)
   --message value, -m value  message for the transaction
```
`history` subcommand help:
```
λ dilawar history -h
NAME:
   dilawar history - shows all history

USAGE:
   dilawar history [command options] [arguments...]

OPTIONS:
   --csv  outputs as comma seperate values
```
no output if there's not record:
```
λ dilawar
```
add a credit:
```
λ dilawar credit -a 1000 -m "chips"
```
check balance:
```
λ dilawar
-1000
```
more transactions:
```bash
# pay half debt
λ dilawar debit -a 500 -m "half"
# check balance again
λ dilawar
-500
# pay more than what you owe
λ dilawar debit -a 1000 -m "new month"
# check balance again
λ dilawar
500
```
show history:
```
λ dilawar history
  ID | DESCRIPTION |      TIME      | TYPE  |  AMOUNT    
+----+-------------+----------------+-------+-----------+
   1 | chips       | 49 seconds ago | CR    | 1,000 PKR  
   2 | half        | 30 seconds ago | DR    | 500 PKR    
   3 | new month   | 9 seconds ago  | DR    | 1,000 PKR  
+----+-------------+----------------+-------+-----------+
                                      TOTAL |    500     
                                    +-------+-----------+
```                                 
save history as CSV:
```
λ dilawar history -csv > records.csv
```
remove all data:
```
λ rm -rf ~/.dilawar
```

## Lisence
MIT