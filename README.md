A simple caesar cipher algorithm for generating monoalphabetic substitution cryptograms using the English alphabet.

Ciphertext is converted to caps and non-alpha characters are stripped.

Usage:

```bash
./cryptogram-generator --shift 3 --msg "Hello, World!"
```

```bash
echo "Hello, World!" | ./cryptogram-generator --shift 3
```

Expected Output

```bash
KHOORZRUOG
```

### Generate a random cryptogram

![Generate a random cryptogram to solve](images/cryptocow.png?raw=true "Generate a random cryptogram to solve")
