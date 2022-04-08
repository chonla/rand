# Randomizer CLI

Random a string in command-line.

## Installation

### Homebrew/Linuxbrew

```
brew tap chonla/universe
brew install rand
```

### From source

```
go get github.com/chonla/rand
```

### Upgrade

```
brew upgrade
```

## Usage

```
Usage of rand:

  rand [-d <dataset>] [-e <dataset-entries>]


  -d string
    	set data set (default "alphanum")
  -e string
    	set data entries, used only when data set is custom
  -h	show this help
  -l int
    	result length (default 10)
  -v	show rand version

  possible dataset:
    alpha : alphabetic (a-z, lowercase)
    alphacase : alphabetic with case-sensitive (a-z + A-Z)
    num : numeric (0-9)
    alphanum : alphanumeric (a-z + 0-9, lowercase)
    alphanumcase : alphanumeric with case-sensitive (a-z + A-Z + 0-9)
    hex : hexadecimal (0-9 + a-f, lowercase)
    hexcase : hexadecimal with case-sensitive (0-9 + a-f + A-F)
    binary : binary string
    custom : custom data set, use -e to specify entries in data set

  possible alternative output format:
    base64 : encode randomized string into base64 string
    hex : encode randomized string into hexadecimal string
```

## Example

```
$ rand
2q1xcz1z4n

$ rand -d hex
e2ee0ead54

$ rand -d hexcase
f4621C7F0A

$ rand -d alphanumcase
gcL0OmI6tL

$ rand -d alphanumcase -l 40
qqgNqJgstDNwEdYYc49CYGI25UUELH2I6gftSf2n

$ rand -d custom -e 123abc -l 32
accbacb32aa1c1caa12a112c22cbb331

$ rand -d binary -l 32 -o base64
CwRuKV0kXYC7pN1pESrUdewcdfnNw44/xCRys3iKV5g=
```

## License

[MIT](LICENSE.txt)