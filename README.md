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

### Verify your download
Make sure you have the binary and it's `sha256sum.txt` downloaded. \
Then, run the following from the same directory:
```bash
sha256sum -c sha256sum.txt
```

Here's what the output should look like:
```
mnstrsay-v0.1.0-linux-amd64: OK
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
