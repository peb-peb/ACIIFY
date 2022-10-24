# ACIIFY

ACIIFY is a tool used to convert images into ASCII art and print them onto the console.

**Input Formats currently Supported:**

- PNG
- JPEG

## Index

- [Installation Instructions](#installation-instructions)
- [CLI Usage](#cli-usage)
- [Sample Outputs](#sample-outputs)
- [Contributing](#contributing)

## Installation Instructions

1. Install `go`. Installation instructions for it are [here](https://go.dev/dl/)
2. Fork the Repository.
3. Clone the repository in your local environment.
4. Go to the directory and run `go build ./main.go`
5. Then, run `go run ./main.go -f ./data/golang.png -r 0.1 -c`

## CLI usage

```text
usage: TODO

options:
  -help
        show this help message and exit
  -c    Display ASCII art in color     
  -f string
        Image Path to Convert
  -p    Display Pixelated art
  -r float
        Ratio to Scale the given Image (default 1)
```

### -f FILENAME

Specify path to file (jpeg, png) you want to convert into ASCII art.

```bash
go run ./main.go -f ./data/<image>
```

Example:
```bash
go run ./main.go -f ./data/gradient.png
```

> NOTE: Most of the time, running without ratio results in very large ASCII arts. So, always try with `-r 0.1` flag.

### -r RATIO (default 1)

Ratio to Scale the given Image into

```bash
go run ./main.go -f ./data/<image> -r <ratio to scale into>
```

Example:
```bash
go run ./main.go -f ./data/golang.png -r 0.1
```

### -c COLOR

This option is used to display the ASCII art in color (256).

```bash
go run ./main.go -f ./data/<image> -c
```

Example:
```bash
go run ./main.go -f ./data/golang.png -r 0.1 -c
```

### -p PIXELATED

This option is used to display Pixelated art of the given image.

```bash
go run ./main.go -f ./data/<image> -p
```

To print colored Pixelated art, use `-c` with `-p`.

```bash
go run ./main.go -f ./data/<image> -r 0.1 -p -c
```

Example:
```bash
go run ./main.go -f ./data/gradient.png -r 0.1 -c -p
```

## Sample Outputs

| Image | ASCII Art |
|:-----:|:---------:|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/golang.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/golang-ascii-bw.png)|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/golang.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/golang-ascii-color.png)|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/gradient.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/gradient-ascii-color.png)|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/gradient.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/gradient-pixelated.png)|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/anya-forger.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/anya-forger-ascii-bw.png)|
|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/anya-forger.png)|![](https://github.com/peb-peb/ACIIFY/blob/776899d8f856d73e198964ba90ec7a3024039958/data/anya-forger-ascii-color.png)|

## Contributing

You can fork the project and implement the changes you want for a pull request. 

Contributions are Welcome! :)
