# Introduction
VidAnnotate is a simple RESTful API for managing videos and annotations. It allows users to create, update, delete, and list videos and annotations. The API uses SQLite3 for persistent storage, and provides basic security through API key and JWT token validation.

## Features
* Create, update, and delete videos
* Create, update, and delete annotations
* List annotations for a specific video
* Validate annotation start and end time against video duration
* Authorize requests through API key and JWT token

## Requirements
* Go (version 1.14 or later)
* SQLite3
* Docker (optional)

## Installation
* Clone the repository: `git clone https://github.com/yourusername/VidAnnotate.git`
* Navigate to the project directory: `cd VidAnnotate`
* Install dependencies: `go mod download`
* Create a new SQLite3 database and update the database configuration in `main.go`
* Run the migration to create tables: `go run main.go migrate`
* Run the application: `go run main.go`

## Configuration
You can configure the following environment variables in the main.go file:

`API_KEY`: API key for authentication
`JWT_TOKEN`: JWT token for authentication

## Usage

The API will be running on http://localhost:8000

You can use the following endpoints to interact with the API:

* `POST /videos`: Create a new video
* `DELETE /videos/{id}`: Delete a video and all related annotations
* `POST /annotations`: Create a new annotation
* `PUT /annotations/{id}`: Update an annotation
* `DELETE /annotations/{id}`: Delete an annotation
* `GET /videos/{id}/annotations`: Get all annotations for a specific video

## Running with Docker
You can use Docker to run the application by building the image and running the container.

* Build the image: `docker build -t vidannotate .`
* Run the container: `docker run -p 8000:8000 vidannotate`
The API will be running on http://localhost:8000

Make sure that you have Docker installed and running on your machine.

## Licensing
VidAnnotate is released under the MIT License. See the LICENSE file for more information.

## Contribution
Any contributions to VidAnnotate are welcome. Please fork the repository and submit a pull request.