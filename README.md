# Collaborative Core API (Golang)
> The **Collaborative Core API** is supported with **Golang** and contains the following services:
> * `Company Services`
> * `Campus Services`
> * `Talent Services`
> * `Community Services`

## Table of Contents
* [Requirements](#requirements)
* [Package Dependencies](#package-dependencies)
* [Run Program](#run-program)
* [Unit Testing](#unit-testing)
* [Submission to Devcode](#submission-to-devcode)
* [Development Guide](#development-guide)
* [Project Status](#project-status)
* [Author](#author)

## Requirements
* [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
* [Go](https://go.dev/doc/install)
* Docker
    * [On Windows](https://docs.docker.com/desktop/install/windows-install/)
    * [On Mac](https://docs.docker.com/desktop/install/mac-install/)
    * [On Linux](https://docs.docker.com/desktop/install/linux-install/)

## Package Dependencies
* github.com/gin-gonic/gin
* gorm.io/gorm
* gorm.io/driver/mysql
* github.com/joho/godotenv
* github.com/cosmtrek/air `Dev`

## Run Program
* Using Local Machine (Windows)
  * Create new database (on MYSQL) as `<database_name>`
  * Copy `.env.example` to `.env` 
    * Update `MYSQL_DBNAME` configuration as `<database_name>`
    * Update `MYSQL_PASSWORD` configuration as `<your_mysql_password>`
  * Download dependencies from `go.mod` and `go.sum`
    
    ```
    go mod download
    ```
  * Start the program
    
    ```
    go run main.go
    ```
  * Open the path on your local machine
    
    ```
    http://localhost:3030/api/
    ```

* Using Docker 
  * Copy `.env.example` to `.env` 
  * Build the Backend API Service docker image. If you don't specify the `<tag>`, it will be tagged as `latest` by default

    ```
    docker build -t <image_name>:<tag> .
    ```
  * Configure `docker-compose.yaml`, adjust the script below according to your built docker image

    ```
    ...
    backend-api-service: 
      image: <image_name>:<tag>
      restart: always
      ports:
        - 8080:3030
    ...
    ```
  * Run `docker-compose.yaml` file, it may take a few minutes and re-attempts. It works fine, solely wait for the `[SERVER] Server to be run on http://0.0.0.0:3030/api` comes out

    ```
    docker-compose -f docker-compose.yaml up
    ```
  * Open the path on your local machine
      
    ```
    http://localhost:8080/api/
    ```

## Unit Testing
* _TODO HERE_

## Submission to Devcode
* _TODO HERE_

## Development Guide
* Download dependencies from `go.mod` and `go.sum`

    ```
    go mod download
    ```
* Install `github.com/cosmtrek/air` module

    ```
    go install github.com/cosmtrek/air@latest
    ```
* Run the program using `github.com/cosmtrek/air` module

    ```
    air
    ```

## Project Status
* Project is: _in progress_

## Author
<table>
    <tr>
      <td><b>Name</b></td>
      <td><b>GitHub</b></td>
    </tr>
    <tr>
      <td>Rava Naufal Attar</td>
      <td><a href="https://github.com/sivaren">sivaren</a></td>
    </tr>
    <tr>
      <td>Suryanto Tan</td>
      <td><a href="https://github.com/SurTan02">SurTan02</a></td>
    </tr>
    <tr>
      <td>Steven Alexander Wen</td>
      <td><a href="https://github.com/loopfree">loopfree</a></td>
    </tr>
</table>
