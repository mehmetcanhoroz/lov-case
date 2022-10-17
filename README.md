<a name="readme-top"></a>
<!-- PROJECT LOGO -->
<br />
<div align="center">
<h3 align="center">Test Case with gRPC</h3>

  <p align="center">
    This repo is created for test case!
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#build">Build</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#features">Features</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

In this repo, I developed an application with desired technologies. It is basically 4 basic math operation calculator.

Features:
* API/Service/Repository Layer
* 4 basic math operation support and default method calculated 4 operations at the same time
* You should implement DRY principles to the rest of your life :smile:


<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

This section contains list major frameworks/libraries used to bootstrap the project.

* [![Golang][Golang]][Golang-url]
* [![Golang][gRPC]][gRPC-url]
* [![Makefile][makefile]][makefile-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

After you clone this repository, you will find here required and useful commands.

### Prerequisites

Before starting to use the app or build, we need to install go mod.
* go.mod
  ```sh
  make install
  ```
* .env file = After executing install command, you will have .env file in your folder. You can set there a few variables such as:
  ```sh
  SERVER_PORT=8000
  SERVER_HOST=localhost
  ```

### Build

_After development or to use this application directly we need to build Server and Client application_

1. Build Server
   ```sh
   make build-server
   ```
2. Build Client
   ```sh
   make build-client
   ```
3. Start server
   ```sh
   ./bin/server
   ```
4. Send Request to Server via Client - 5 Different options we have
   ```sh
   ./bin/client -method add -a 20 -b 20
   ```
   ```sh
   ./bin/client -method sub -a 50 -b 20
   ```
   ```sh
   ./bin/client -method div -a 10 -b 2
   ```
   ```sh
   ./bin/client -method multi -a 50 -b 20
   ```
   ```sh
   ./bin/client -method all -a 20 -b 10
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- FEATURES -->
## Features

- [x] Addition, subtraction, multiplication and division.
- [x] Subtraction, multiplication and division.
- [x] Multiplication and division.
- [x] Division.
- [x] All calculation in an api
- [x] Build and manage bins with Makefile

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Mehmet Canhoroz - [@mehmetcanhoroz](https://twitter.com/mehmetcanhoroz)



[![Makefile][linkedin-shield]][linkedin-url]  [https://www.linkedin.com/in/mehmet-canhoroz/](https://www.linkedin.com/in/mehmet-canhoroz/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Golang]: https://img.shields.io/badge/golang-000000?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://go.dev/
[gRPC]: https://img.shields.io/badge/grpc-0769AD?style=for-the-badge&logo=grpc&logoColor=white
[gRPC-url]: https://grpc.io/
[makefile]: https://img.shields.io/badge/Makefile-DD0031?style=for-the-badge&logo=makefile&logoColor=white
[makefile-url]: https://www.gnu.org/software/make/manual/make.html
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/mehmet-canhoroz