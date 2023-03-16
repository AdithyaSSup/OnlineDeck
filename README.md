# Toggle- Online Deck
This repository contains a REST API written in Golang that simulates a deck of cards.

Details of the Product requirement/PRD can be found [here](https://toggl.notion.site/Toggl-Backend-Unattended-Programming-Test-015a95428b044b4398ba62ccc72a007e)

* The API is built using the [Gin](https://github.com/gin-gonic/gin) router and follows the [Clean Architecture principles](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html), as described by Uncle Bob. This ensures that the code is modular and easy to maintain.

* The dependencies are organized using a controller-service-dao structure, where the dependencies are concentric. 
  The controller layer depends on the service layer, which in turn depends on the dao layer.  
  All dependencies are injected using the Google Wire library at the controller level.

* The package structure follows a similar pattern, where the service, model, and dao layers are separated into their own package folders. 
  Any client, such as a controller or worker, can depend on the implementations provided by these packages.

* Instead of using a database, the API utilizes an in-memory map that is created and injected into the DAO layer during server creation and startup.

## Getting Started
To get started with this API, you will need to have Golang installed on your machine. You can then clone this repository and run the following command from your termainal by navigating to entry directory:

<pre>
# cd to entry directory to run the project
cd entry
go mod tidy
#this will install all the depencies
go run main.go
#This will start the API on http://localhost:3000/api/v1.
</pre>

## Usage
### 1. Create a new deck <br>
To create a new full deck, send a POST request to /decks. The response will contain a JSON object with the id of the new deck:
To create a partial deck, send a POST request to /decks?cards=(card_code).
<pre>
#Request:
POST http://localhost:3000/api/v1/decks?cards=(card_code)
{
   shuffle :(true/false)
}

#Response:
{
    "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
    "shuffled": false,
    "remaining": 52
}

</pre>

### 2. Draw cards from an existing deck <br>
To draw cards from an existing deck, send a GET request to /deck/:id/draw?count=<number>, where :id is the ID of the deck you want to draw from and count is the count of cards that you want to draw. The response will contain a JSON object with the drawn cards:

<pre>
#Request:
GET http://localhost:3000/api/v1/decks/a251071b-662f-44b6-ba11-e24863039c59/draw?count=2

#Response:
{
    "cards": [
        {
            "value": "QUEEN",
            "suit": "HEARTS",
            "code": "QH"
        },
        {
            "value": "4",
            "suit": "DIAMONDS",
            "code": "4D"
        }
    ]
}
</pre>
### 3. Open an existing deck <br>
To open an existing deck, send a GET request to /deck/:id, where :id is the ID of the deck you want to open. The response will contain a JSON object with the details of the deck:

<pre>
#Request:
GET http://localhost:3000/api/v1/decks/a251071b-662f-44b6-ba11-e24863039c59

#Response:
{
    "deck_id": "a251071b-662f-44b6-ba11-e24863039c59",
    "shuffled": false,
    "remaining": 3,
    "cards": [
        {
            "value": "ACE",
            "suit": "SPADES",
            "code": "AS"
        },
				{
            "value": "KING",
            "suit": "HEARTS",
            "code": "KH"
        },
        {
            "value": "8",
            "suit": "CLUBS",
            "code": "8C"
        }
    ]
}
</pre>

### Tests
Unit tests for the service layer can be found in the "service" directory. To run these tests, navigate to the "pkg" folder in the terminal and run the following command:
<pre> go test -v ./... </pre>
It should be noted that even though the tests are written in the service layer, the DAO layer is mocked for all test cases, making them unit tests.
