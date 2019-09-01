# devChallenge

devChallenge is  a simple browser-based chat application using Go.
It's based in  the gorilla library and it's composed for 3 main parts

- Backend will spawn a service on port 8080 and will accept tcp connections
on endpoint `/ws`

- Frontend is a little React app that will be used for rendering the chat
It should start on `http://localhost:3000/` or you can set a `PORT`env variable
and the app should run there.

- Bot is just a little client that will be connected to the Backend 
tpc port by default and would act as another client listening the chat
when a message in the format of `AAPL.US`  It should read that comand
request and endpoint and return a parsed message to the client

### Prerequisites

- Install npm
- Install npx with install -g npx
- Install Go versions 1.11+

###Running the test
For the sake of simplicity and adding the binary files to the repository
- Backend, this should be the first one to run
```sh
$ cd backend
$ go run main.go
```

or just
```
$ cd backend
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
$ go run main.go
```
or just
```
$ cd bot
$ ./main
```



### Improvements/Tradeoffs
- Google api Key was stored in the react app for managing the login
with Firebase, this can be stored as and env variable but I didn't 
for the sake of simplicity, the key if for a fake chat app that
is never going to be developed.

- Port for the bot to be able to connect to any tcp could be added 
as an env variable too.

- Messages are not stored in any database for simplicity