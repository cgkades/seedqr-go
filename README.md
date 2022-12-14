# seedqr-go


This project was created to generate a SeedQR on an air-gapped system for use with.... well anything that can take a SeedQR (i.e. Blockstream Jade).

The idea was to generate a SeedQR on an air-gapped system and save the png into [KeePassXC](https://keepassxc.org/). I also didn't want to load a bunch of python libraries, so having a single binary is also ideal.


```shell
seedqr -- SeedQR Generator for air-gapped systems
Flags:
  -d    disable QR Code border
  -i    invert black and white
  -m int
        Size of Mnemonic (12 or 24) (default 24)
  -o string
        out PNG file prefix, empty for stdout
  -s int
        image size (pixel) (default 256)
  -t    print as text-art on stdout
Example: seedqr -o myseed -t
```
