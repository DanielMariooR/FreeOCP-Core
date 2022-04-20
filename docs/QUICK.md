# Quick Start Guide
This quick start guide is for helping those that want to quickly see the features of the freeocp website without going into production grade deployment.

## Requirement
- Golang 1.17 or higher
- NPM 16 or higher
- MySQL
- Git

## Golang Backend
- Clone the FreeOCP-Core backend from [here](https://github.com/DanielMariooR/FreeOCP-Core)
- Copy and .env.example file content into .env file and fill in with your MySQL credentials and desired port
- Execute the following command to start backend service `go run main.go &`
- Check to see if the service is up and running

## React Frontend
- Clone the FreeOCP-Client from [here](https://github.com/DanielMariooR/FreeOCP-Client)
- Run the following command `npm install` to install all the required dependencies
- Run the following command `npm start` to start the npm client
- Open the client in `http://localhost:3000` to start exploring the freeocp website.
