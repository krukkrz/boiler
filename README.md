# Boiler
It is a boiler plate generator for Spring MVC classes.

## Setup
1. Clone this repository.
1. Fire up application by running a command:
    ```bash
    ./boiler <your domain name> <your base package name>
    ```
1. Copy a content of a folder `target` and paste it in your `src/java/` directory

## Result
A result classes will have the following structure:
``` bash
.
└── example
    ├── controllers
    │   ├── ExampleController.java
    │   └── ExceptionController.java
    ├── models
    │   ├── dtos
    │   │   └── ExampleDto.java
    │   ├── entities
    │   │   └── Example.java
    │   └── mapping
    │       └── ExampleMapper.java
    ├── repositories
    │   └── ExampleRepository.java
    └── services
        └── ExampleService.java
```

Example of generated classes can be found in the folder `./exampleResult`