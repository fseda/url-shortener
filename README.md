# URL Shortener Project

A URL shortener built using Go for the server, and HTML/CSS for the user interface. This project allows you to create shortened versions of long URLs, making them more convenient to share and manage. [Check It Out Live!](https://short-url-08gu.onrender.com/)

 ![Screenshot](/static/imgs/screenshot.png) 

## Table of Contents

- [URL Shortener Project](#url-shortener-project)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technologies](#technologies)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation and Running](#installation-and-running)
    - [Run with Docker](#run-with-docker)

## Features

- Shorten long URLs into user-friendly and compact links.
- Redirect users to the original URLs when they access the shortened links.
- Track the number of times each shortened link is accessed.

## Technologies

- Go: Server-side logic and URL shortening
- HTML: User interface structure
- CSS: Styling and design

## Getting Started

Follow these steps to get the URL shortener up and running on your local machine.

### Prerequisites

- Go installed on your system
- Basic understanding of Go, HTML, and CSS

### Installation and Running

1. Clone the repository:

   ```bash
   $ git clone https://github.com/yourusername/url-shortener.git
   $ cd url-shortener
   ```
   
2. Install dependencies
   
   ```bash
   $ go mod tidy
   $ go get
   ```

3. Compile and Run

    ```bash
    $ go build
    $ ./url-shortener
    ```

### Run with Docker

Requires "docker-compose" to work

1. Build and Run from Docker Compose
   
   ```bash
   $ docker-compose up --build
   ```