# OnlineDeck
This repository contains a REST API written in Golang that simulates a deck of cards. It provides three endpoints:

* POST /deck?cards= to create a new deck 
* GET /deck/:id/draw to draw cards from an existing deck
 * GET /deck/:id to open an existing deck

* The API is built using the Gin router and follows the Clean Architecture principles, as described by Uncle Bob. This ensures that the code is modular and easy to maintain.

* The dependencies are organized using a controller-service-dao structure, where the dependencies are faced inwards. 
  The controller layer depends on the service layer, which in turn depends on the dao layer.  
  All dependencies are injected using the Google Wire library at the controller level.

* The package structure follows a similar pattern, where the service, model, and dao layers are separated into their own package folders. 
  Any client, such as a controller or worker, can depend on the implementations provided by these packages.

## Getting Started
To get started with this API, you will need to have Golang installed on your machine. You can then clone this repository and run the following command:

<pre>
go mod tidy
this will install all the depencies
go run main.go
This will start the API on http://localhost:3000.
</pre>

Usage
1. Create a new deck
To create a new deck, send a POST request to /deck. The response will contain a JSON object with the id of the new deck:
<pre>

POST /deck

Response:
{
  "id": "abc123"
}

</pre>

2. Draw cards from an existing deck
To draw cards from an existing deck, send a GET request to /deck/:id/draw, where :id is the ID of the deck you want to draw from. The response will contain a JSON object with the drawn cards:

<pre>
GET /deck/abc123/draw?count=2

Response:
{
  "cards": [
    {
      "value": "10",
      "suit": "spades"
    },
    {
      "value": "Q",
      "suit": "hearts"
    }
  ]
}
</pre>
Open an existing deck
To open an existing deck, send a GET request to /deck/:id, where :id is the ID of the deck you want to open. The response will contain a JSON object with the details of the deck:

<pre>
GET /deck/abc123

Response:
{
  "id": "abc123",
  "cards": [
    {
      "value": "2",
      "suit": "hearts"
    },
    {
      "value": "5",
      "suit": "diamonds"
    },
    ...
  ]
}
</pre>
