str-replace
=============================
Simple tools to handle string and subdomain permutations

## Install

```bash
go install github.com/j3ssie/str-replace@latest
```

## Usage

```bash
# Simple tools to handle string and subdomain permutations
cat list-of-subdomain.txt | str-replace -d '.' -j ','

# build the wordlist
cat list-of-subdomain.txt | str-replace -d '.' -n

# append the wordlist to existing subdomain
cat list-of-subdomain.txt | str-replace -W wordlists.txt -tld example.com
cat list-of-subdomain.txt | str-replace -W wordlists.txt -j '.' -s

```

