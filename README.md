# Improvised 
A dynamic TCP load balancer for easily switch between multiple services.

## Goal
The goal of this project is to add and remove on the fly services.

Now it only supports Redis as services configuration 

## Usage
- Put a list of services (like `localhost:8080`) under the key `improvised:servers` on redis
- run improvised

```
Usage:
  improvised [OPTIONS]

Application Options:
      --redisAddress  Redis address
      --redisUsername Redis username
      --redisPassword Redis password
      --redisDB       Redis db number (default: 0)
  -p, --port          The port where improvised will run (default: 8888)

Help Options:
  -h, --help           Show this help message
```

Note: the default port is `8888` you can change it by putting a new port as the program argument.
