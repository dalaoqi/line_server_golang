# line_server_golang

This is a simple line server written in Golang.

## Usage

* save all messages sent by users
* get all messages sent by users
* broadcast


## Getting Started


### development environment

![Golang](https://img.shields.io/badge/Golang-1.19.1-blue)
![Docker](https://img.shields.io/badge/Docker-20.10.12-blue)
![docker-compose](https://img.shields.io/badge/docker_compose-1.29.2-blue)


You should replace the url of webhook in line developer console with a url generated by `ngrok http port`.

Example: `https://d685-61-65-116-30.jp.ngrok.io` -> `http://localhost:port`


## Running the System

```shell=
git clone https://github.com/dalaoqi/line_server_golang.git

cd line_server_golang

# start
make
```
`make help` to see more.

## Configuration

```yaml=
app:
  lineAccessToken: <your Channel access token>
  lineSecret: <your Channel secret>
  port: <listening port of server, default: 1234>
  
db:
  mongo:
    #name of database
    name: cinnox 
    url: mongodb://localhost:27017 
    
    lineEvent: 
      #name of collection
      name: lineEvent
```

## Example

![](assets/demo.gif)

[original size](https://imgur.com/a/rlKaBFi)