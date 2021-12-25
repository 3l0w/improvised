# Improvised 
A dynamic TCP load balancer for easily switch between multiple services.

## Goal
The goal of this project is to add and remove on the fly services.

Now it only supports Redis as services configuration 

## Usage
- Connect Redis by modifying the source code (will be changed ASAP)
- Put a list of services (like `localhost:8080`) under the key `improvised:servers`
- start the program

Note: the default port is `8888` you can change it by putting a new port as the program argument.
