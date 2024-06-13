
# PDF Encrypter

This is a simple tool to encrypt PDF files using the [pdfcpu](https://github.com/pdfcpu/pdfcpu) library. It supports both individual files and directories containing multiple PDF files.

## Usage

```sh
pdfencrypter [the path for input file or directory] [optional: owner password]
```

If the owner password is not provided, the tool will read the password from a `pass.txt` file located in the same directory.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/pdfencrypter.git
    ```
2. Change to the repository directory:
    ```sh
    cd pdfencrypter
    ```
3. Build the project:
    ```sh
    go build -o pdfencrypter
    ```

## Examples

### Encrypt a Single File

To encrypt a single PDF file, provide the path to the file and the owner password:
```sh
./pdfencrypter /path/to/your/file.pdf ownerpassword
```

### Encrypt All Files in a Directory

To encrypt all PDF files in a directory, provide the path to the directory:
```sh
./pdfencrypter /path/to/your/directory ownerpassword
```

### Using a Password File

If you do not provide an owner password as a command line argument, the tool will read the password from a file named `pass.txt`:
1. Create a file named `pass.txt` in the same directory as the executable.
2. Add your password to the `pass.txt` file.
3. Run the tool without the owner password argument:
    ```sh
    ./pdfencrypter /path/to/your/file_or_directory
    ```

### Display Help

To display the usage information, use the `-h` flag:
```sh
./pdfencrypter -h
```

## Dependencies

This project uses the [pdfcpu](https://github.com/pdfcpu/pdfcpu) library. Make sure you have Go installed and properly set up on your machine.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## Author

[Your Name](https://github.com/yourusername)
