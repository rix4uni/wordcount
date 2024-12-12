## wordcount

Generate custom wordlist according to the word use count.

## Installation
```
go install github.com/rix4uni/wordcount@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/wordcount/releases/download/v0.0.1/wordcount-linux-amd64-0.0.1.tgz
tar -xvzf wordcount-linux-amd64-0.0.1.tgz
rm -rf wordcount-linux-amd64-0.0.1.tgz
mv wordcount ~/go/bin/wordcount
```
Or download [binary release](https://github.com/rix4uni/wordcount/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/wordcount.git
cd wordcount; go install
```

## Usage
```
Usage:
  wordcount [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dnswordlist Process subdomains and generate a word count list
  help        Help about any command

Flags:
  -h, --help     help for wordcount
  -t, --toggle   Help message for toggle
```

## Usage Examples

#### dnswordlist
```bash
┌──(root㉿kali)-[/root/wordcount]
└─# cat subs.txt | wordcount dnswordlist -o best-dns-wordlist.json
```