# ACIIFY

ACIIFY is a tool used to convert images into ASCII art and print them onto the console.

**Input Formats currently Supported:**

- PNG
- JPEG

## Index



## Installation Instructions

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

## Options

### -f FILENAME

Specify path to file (jpeg, png) you want to convert into ASCII art.

```bash
go run ./main.go -f ./data/<image>
```

### -r RATIO (default 1)

Ratio to Scale the given Image into

```bash
go run ./main.go -f ./data/<image> -r <ratio to scale into>
```

### -c COLOR

This option is used to display the ASCII art in color (256).

```bash
go run ./main.go -f ./data/<image> -c
```

### -p PIXELATED

This option is used to display Pixelated art of the given image.

```bash
go run ./main.go -f ./data/<image> -p
```

## Sample Outputs

## Contributing

You can fork the project and implement the changes you want for a pull request. 

Contributions are Welcome! :)