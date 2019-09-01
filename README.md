# devChallenge

devChallenge is  a simple browser-based chat application using Go.
It's based in  the gorilla library and it's composed of 3 main parts

- Api will spawn a little service on port 8080 and will accept tcp connections
on endpoint `/ws`

- Frontend is a little React app that will be used for rendering the chat
It should start on `http://localhost:3000/`.

- Bot is just a little client that will be connected to the Api 
tpc port by default and would act as another client listening the chat
when a message in the format of `AAPL.US`  It should read that comand
request and endpoint and return a parsed message to the client

### Prerequisites

- Install npm
- Install npx with install -g npx
- Install Go versions 1.11+

### Running the challenge
- Api, this should be the first one to run
```sh
$ cd api
$ go build main.go
$ go run main.go
```

or just after building
```
$ cd api
$ ./main
```

- FrontEnd
```sh
$ cd frontend
$ npm start
```

- Bot
```sh
$ cd bot
$ go build main.go
$ export ADDRESS=localhost:8080 && go run main.go
```
or just after building
```
$ cd bot
$ export ADDRESS=localhost:8080 && ./main
```



### Improvements/Tradeoffs
- Google api Key was stored in the react app for managing the login
with Firebase, this can be stored as and env variable but I didn't 
for the sake of simplicity, the key if for a fake chat app that
is never going to be developed.

- Messages are not stored in any database for simplicity

- Dockerize would iprove deployment

- Have the user information as part of the message

- list of latest

### Bonus
- Handle messages that are not understood or any exceptions raised within the bot.