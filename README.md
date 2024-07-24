# Social Media Backend Using Golang

## Overview

This project is a comprehensive social media backend developed using Golang. It provides the essential features and functionalities required for a modern social media platform, including user authentication, post management, comments, likes, follow system, and trending posts.

## Key Features

- **User Authentication and Authorization**
   - Secure user registration and login with JWT.
   - Password hashing and validation.

- **Post Management**
   - Create, read, update, and delete posts.
   - Add, update, and delete comments on posts.
   - Like and unlike posts.

- **Follow System**
   - Follow and unfollow users to build connections.

- **Trending Posts**
   - Algorithm to determine and display trending posts based on engagement scores and activity.

## Technology Stack

- **Programming Language**: Golang
- **Web Framework**: Fiber
- **Database**: MongoDB
- **Authentication**: JWT
- **Configuration**: Environment variables

## Architecture

- **Models**
   - Define the structure of various entities such as users, posts, comments, likes, and follow relationships.

- **Handlers**
   - Implement the business logic for handling HTTP requests.

- **Routes**
   - Define API endpoints for user interactions, post management, follow system, and trending posts.

- **Middleware**
   - JWT middleware for secure access to protected routes.
   - CORS middleware for cross-origin request handling.

- **Database**
   - MongoDB for scalable and flexible data storage.
   - MongoDB collections for storing users, posts, comments, likes, and follow relationships.

## Getting Started

### Prerequisites

- Go 1.16 or later
- MongoDB
- Git

## Run Locally

Clone the repository
```bash
    git clone https://github.com/Srivastava-samarth/Social-Media-Backend.git
```

Start the server

```bash
  go run main.go
```


### Setup Environment Variables

Create a `.env` file in the root directory and add the following variables:

```env
PORT=5757
MONGO_URI=<your-mongodb-uri>
JWT_SECRET=<your-jwt-secret>
```

## Contributing

Contributions are always welcome!Please fork the repository and submit a pull request.

