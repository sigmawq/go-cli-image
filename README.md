# Cli image

![Alt text](gopher.png?raw=true)

Display PNG and JPG images in terminal using the 8 basic ANSI color codes.

## Installation
`go build .`

## Usage
`cli-image <path-to-image> <draw-unit>`
`draw-unit` is `#` by default. You can change it to any characher that is supported in your terminal.

Please note that you need to extend the horizontal (or disable line wrapping completely) to fit the "image".
Also you need to make sure that your terminal supports ANSI color escape sequences and they are enabled. 

## License
[MIT](https://choosealicense.com/licenses/mit/)