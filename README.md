str-replace
=============================
Simple tools to handle string and generate subdomain permutations

## Install

```bash
go install github.com/j3ssie/str-replace@latest
```

## Usage

```bash
# Build the wordlist by splitting subdomains as '.' string
cat list-of-subdomain.txt | str-replace -d '.' -n

# Build permutation subdomains from the wordlist from the existing subdomains
# This will replace every part of the subdomain except the tld with the wordlist provided
cat list-of-subdomain.txt | str-replace -W wordlists.txt -tld example.com

```

## Don't know how to use it? Well, This is already integrated into the Osmedeus workflow.

<p align="center">
  <img alt="OsmedeusEngine" src="https://raw.githubusercontent.com/osmedeus/assets/main/logo-transparent.png" height="200" />
  <p align="center">
    This project was part of Osmedeus Engine. Check out how it was integrated at <a href="https://twitter.com/OsmedeusEngine">@OsmedeusEngine</a>
  </p>
</p>

