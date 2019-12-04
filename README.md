# GoBerry

## Usage
```bash
$ ./goberry -pin 1010
```
* pin: the controller token

## Understanding the logs
```
18:20:49 JSON path not set # expected since the flag telling where the json config is was not defined, it defaults to the current directory
18:20:49 Starting 4 consumers # starting 4 threads for consuming what the server is sending
18:20:49 Starting GoBerry # go berry running
18:20:49 Consumer up # message for each consumer/thread that is up
18:20:49 Consumer up
18:20:49 Consumer up
18:20:49 Consumer up
18:20:49 Controller already exists # if the controller already exists in the API
18:20:49 Server accepting connections # server is up
18:20:49 Client trying connection with: 127.0.0.1:9001 # tcp client trying to connect with the server
18:20:49 Everything started # last print after starting all threads
18:20:49 Polling controller # http polling controller stuff
18:20:49 Client could not connect:
	 dial tcp 127.0.0.1:9001: connect: connection refused # if the client could not connect to the server, will try again soon
```

## Requirements
* Go Installed

## Installing GoLang
[Link](https://golang.org/doc/install)

## Building
### Local
For building the project to run locally run this command inside the project folder
```bash
$ go build
```
This generates an executable binary called `goberry` which you can execute.

### RaspBerry PI
Building the project to RaspBerry PI requires setting environments variables that will make it build to Arm architecture. The following command will be enough:
```bash
$ env GOARCH=arm GOOS=linux go build
```
If ran succesfully a binary will be created called `goberry` which you must run inside the RaspBerry PI.

## Sending to the RaspBerry PI
Sending to the RaspBerry PI is simple, the only requirements is being able to `ssh` into it. Pretend the RaspBerry PI has IP 192.168.100.100.
```bash
$ env GOARCH=arm GOOS=linux go build # builds the goberry
$ scp ./goberry pi@192.168.100.100:/home/pi/GoBerry/ # sends the binary
$ scp ./cfg.json pi@192.168.100.100:/home/pi/GoBerry/ # sends the config file
```

## Configurations
A JSON is available and is a requirement for running GoBerry, it must be inside the same directory of the binary running.

```json
{
  "api": {
    "protocol": "http",
    "host": "127.0.0.1",
    "port": 4001,
    "pollInterval": 2
  },
  "socketServer": {
    "host": "",
    "port": "9000"
  },
  "socketClient": {
    "host": "127.0.0.1",
    "port": "9001"
  }
}
```

There are three parts: api, socketServer and socketClient.

### API
* protocol: keep it HTTP, it is the protocol used to communicate with the API
* host: the host (ip or domain) running the API
* port: tcp port for connecting
* pollInterval: interval in seconds for each poll made to the api looking for new commands

### Socket Server
* host: blank means the localhost, since GoBerry is running inside RaspBerry PI it should be kept blank
* port: tcp port that will be opened

### Socket Client
* host: the host of the socket client opened by the electronics team
* port: tcp port to connect
