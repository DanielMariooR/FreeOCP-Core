# Documentation

FreeOCP is a free for all open source template for educational website written in Golang and React. FreeOCP is designed for ease of use and resource efficient deployment. It contains all the important features for educational site such as course enrollment and course and problem contribution.

## Technology 
- Golang
- React
- MySQL

## Architecture
![arch-diagram](https://raw.githubusercontent.com/DanielMariooR/FreeOCP-Core/main/docs/assets/arch.png)

FreeOCP is built with Golang with echo to handle backend request using REST API and React as the client. 

## Structural View
![structure](https://raw.githubusercontent.com/DanielMariooR/FreeOCP-Core/main/docs/assets/component-diagram.png)

The Components are divided into services with each services handling different aspects of the Freeocp website.
- User Service: User register and login and authentication
- Assignment Service: Handle assignment related API and business logic
- Problem Service: Handle problem curation and problem creation logic
- Course Service: Handle course enrollment and course creation logic

## Sequence Diagram

## API Documentation
To See the api documentation please refer to this link [here](https://docs.google.com/document/d/1IIhpZ0rWWhNsvcN9MmN-fkC1oy0tTYWPRU5oQhFBlWA/edit)
