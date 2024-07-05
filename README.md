# Ascii-Art-Web

Ascii-Art-Web is a web application that converts text into ASCII art using different banner styles. This project aims to provide an engaging, interactive, and intuitive user experience.

## Project Files
### asciiArtFunctions/

- **asciiwebfiles.go:** Contains the main logic for ASCII art conversion.
- **file_check.go:** Handles file existence and validation.
- **loadbanner.go:** Loads banner files into the application.
- **printbanner.go:** Prints the loaded text in ASCII art format.

### bannerfiles/

- **shadow.txt:** Contains the Shadow banner style.
- **standard.txt:** Contains the Standard banner style.
- **thinkertoy.txt:** Contains the Thinkertoy banner style.

### static/css/

- **style.css:** Stylesheet for the web application.

### template/

- **index.html:** Main HTML template for the input page.
- **result.html:** HTML template for displaying the conversion result.

### main.go
- The main entry point of the application. It initializes the server, sets up the routes, and starts listening for incoming requests.

## Features

* Multiple Banner Styles: Choose from different banner styles such as Shadow, Standard, and Thinkertoy.
* Real-time Conversion: Convert text to ASCII art instantly.
* Responsive Design: A user-friendly interface that works on all devices.

### Installation

* Clone the repository:

```sh
git clone https://github.com/ombima56/ascii-art-web.git
cd ascii-art-web
```
### Install dependencies:

```sh
go mod tidy
```
Run the application:

```sh
go run `main.go`
```
or
```sh
go run .
```
## Usage

* Open your web browser and go to `http://localhost:8080`.
* Enter the text you want to convert in the input field.
* Select a banner style from the options.
* Click the "generate" button to see the ASCII art.

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for any improvements.
## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more details.