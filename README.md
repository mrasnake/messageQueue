# Message Queue

Small Client/Server utilizing RabbitMQ for the transport layer

Message Queue is a small program designed to solve an interview take home assignment.

## Assignment Guidelines

You need to implement a **Client-Server application** with the following requirements:
* multiple-threaded server;
* clients;
* External queue between the clients and server;

Clients:
* Should be configured from a command line or from a file (you decide);
* Can read data from a file or from a command line (you decide);
* Can request server to AddItem(), RemoveItem(), GetItem(), GetAllItems()
* Data is in the form of strings;

* Clients can be added / removed while not intefering to the server or other clients ;

Server:
* Has a data structure that holds the data in the memory
* Server should be able to add an item, remove an item, get a single or all item from the data structure;

External queue:
* Can be Amazon Simple Queue Service (SQS) or RabbitMQ (you decide);


Clients send requests to the external queue - while the server reads those and execute them on its data structure. You define the structure of the messages (AddItem, RemoveItem, GetItem, GetAllItems)


The flow of the project:
1. Multiple clients are sending requests to the queue (and not waiting for the response).
2. Server is reading requests from the queue and processing them, the output of the server is written to a log file
3. Server should be able to process items in parallel
4. log messages (debug, error) are written to stdout


Definition of success:
* Working project that can be executed on your computer;
* Being able to explain how the project works and how to deploy the project (for the first time) on another computer;
* If you take something from the Internet or consult anyone for anything, you should be able to understand it perfectly;
* Code has no bugs, no dangling references / assets / resources, no resource leaks;
* Code is clean and readable;
* Code is reasonably efficient (server idle time will be measured).
* Working with channels when needed


You should develop the project using Golang.


## Getting started

Make sure you have RabbitMQ. You may find instructions here:
https://www.rabbitmq.com/download.html

When you have RabbitMQ install start an instance of it. ex:	`brew services start rabbitmq`

Build both the client/server with `make build`

Other Makefile command include: `make clientDefault`, `make serverDefault`


## Client Configurations

You get usage help with the -h --help flags

### Request File

you must define a file to read the list of requests from.
Flag: -f <filename>
EnvVar: REQUEST_FILE

### Queue Name

you may define a name for the queue you are using, allowing for multiple client/server pairings.
Flag: -q <name>
DefaultValue: "requests"
EnvVar: QUEUE

### Queue Connection

you may configure the RabbitMQ host you would like to use.
Flag: -c <address>
DefaultValue: "amqp://guest:guest@localhost:5672/"
EnvVar: CONNECTION


## Server Configurations

You get usage help with the -h --help flags

### Queue Name

you may define a name for the queue you are using, allowing for multiple client/server pairings.
Flag: -q <name>
DefaultValue: "requests"
EnvVar: QUEUE

### Queue Connection

you may configure the RabbitMQ host you would like to use.
Flag: -c <address>
DefaultValue: "amqp://guest:guest@localhost:5672/"
EnvVar: CONNECTION

### Log File

you may define an output file for the results of the requests.
Flag: -l <filename>
DefaultValue: "./logfile-<Timestamp>.log"
EnvVar: LOG_FILE

## Makefile Commands
### Build
`make build`

### Run Client In Default Location
`make clientDefault`

### Run Server In Default Location
`make serverDefault`