# Message Queue

Small Client/Server utilizing RabbitMQ for the transport layer


## Getting started

Make sure you have RabbitMQ. You may find instructions here:
https://www.rabbitmq.com/download.html

When you have RabbitMQ install start an instance of it. ex:	`brew services start rabbitmq`

Build both the client/server with `make build`

Other Makefile command include: `make clientDefault`, `make serverDefault`


##Client Configurations

You get usage help with the -h --help flags

###Request File

you must define a file to read the list of requests from.
Flag: -f <filename>
EnvVar: REQUEST_FILE

###Queue Name

you may define a name for the queue you are using, allowing for multiple client/server pairings.
Flag: -q <name>
DefaultValue: "requests"
EnvVar: QUEUE

###Queue Connection

you may configure the RabbitMQ host you would like to use.
Flag: -c <address>
DefaultValue: "amqp://guest:guest@localhost:5672/"
EnvVar: CONNECTION


##Server Configurations

You get usage help with the -h --help flags

###Queue Name

you may define a name for the queue you are using, allowing for multiple client/server pairings.
Flag: -q <name>
DefaultValue: "requests"
EnvVar: QUEUE

###Queue Connection

you may configure the RabbitMQ host you would like to use.
Flag: -c <address>
DefaultValue: "amqp://guest:guest@localhost:5672/"
EnvVar: CONNECTION

###Log File

you may define an output file for the results of the requests.
Flag: -l <filename>
DefaultValue: "./logfile-<Timestamp>.log"
EnvVar: LOG_FILE
