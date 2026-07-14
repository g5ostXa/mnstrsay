## MNSTRSAY
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

### Installation
I recommend you download the latest release's binary and it's `sha256sum.txt` from \
the [releases page](https://github.com/g5ostXa/mnstrsay/releases), so you can verify your download.

#### Using curl
Get the binary and the sha256sum:
```bash
curl -L -O "https://github.com/g5ostXa/mnstrsay/releases/download/v0.1.1/mnstrsay-v0.1.1-linux-amd64"
curl -L -O "https://github.com/g5ostXa/mnstrsay/releases/download/v0.1.1/sha256sum.txt"
```
Verify your download:
```bash
sha256sum -c sha256sum.txt
```

Here's what the output should look like:
```
mnstrsay-v0.1.1-linux-amd64: OK
```

#### Using git
Get the full source (latest git)
```bash
git clone --depth=1 https://github.com/g5ostXa/mnstrsay.git
```

You'll need to manually build the binary:
```bash
go build -o "mnstrsay" ./cmd/mnstrsay
```

### Usage
First, make sure the binary is executable:
```bash
chmod +x mnstrsay-bin-name
```

Then, to display the ascii art and the default message, \
simply run the binary without any options:
```bash
./mnstrsay-bin-name
```

To display a custom message:
```bash
./mnstrsay-bin-name "Your personalized message here"
```
