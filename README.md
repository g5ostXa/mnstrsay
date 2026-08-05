# MNSTRSAY
_**mnstrsay**_ is just another cowsay-like program

```go
 ──────────────╮
| Hello, world |
╰──────────────
     \
      \
       \
    /\_______/\
    |         |
    |  X   X  |
   <     -     >
   (           )
    \/-vvvv-\/
     )      (
     {######}
      \____/
```

## Installation
I recommend you download the latest release's binary and it's `sha256sum.txt` from \
the [releases page](https://github.com/g5ostXa/mnstrsay/releases), so you can verify your download.

#### Using curl
Get the binary and the sha256sum:
```bash
curl -L -O "https://github.com/g5ostXa/mnstrsay/releases/download/v0.1.4/mnstrsay-v0.1.3-linux-amd64"
curl -L -O "https://github.com/g5ostXa/mnstrsay/releases/download/v0.1.4/sha256sum.txt"
```
Verify your download:
```bash
sha256sum -c sha256sum.txt
```

Here's what the output should look like:
```
mnstrsay-v0.1.3-linux-amd64: OK
```

Move the binary anywhere on  your `$PATH` or simply move it to `$GOBIN`:
```bash
mv mnstrsay-v0.1.3-linux-amd64 "$GOBIN"/.
```

#### Using git
Get the full source (latest git)
```bash
git clone --depth=1 https://github.com/g5ostXa/mnstrsay.git
```

Build the binary:
```bash
cd ./mnstrsay && make build
```

Make sure the binary is executable:
```bash
chmod +x mnstrsay-bin-name
```

## Usage
Print the cuurent version:
```bash
mnstrsay --version
```

To display the ascii art and the default message, \
simply run the binary without any options:
```bash
mnstrsay
```

To display a custom message:
```bash
mnstrsay "This is my custom message"
```
