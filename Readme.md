# Displaying Swagger Documentation with go-swagger

This repository provides instructions on how to generate and display Swagger documentation for a Go project using the go-swagger tool. Swagger documentation is a useful way to document your API endpoints and make it easier for other developers to understand how to interact with your project.

## Prerequisites

Before you can generate and display Swagger documentation, you will need to install the following tools:

- [Go](https://golang.org/doc/install)
- [go-swagger](https://github.com/go-swagger/go-swagger)

## Generating Swagger Documentation

To generate Swagger documentation for your Go project, you can use the following command:

```
swagger generate spec -o ./swagger.json --scan-models
```

This command will generate Swagger documentation in JSON format and save it to a file called `swagger.json`.

## Displaying Swagger Documentation

To display the Swagger documentation, you can use the following command:

```
swagger serve ./swagger.json
```

This command will start a local server and display the Swagger documentation in a web browser. The Swagger UI will be available at `http://localhost:8080/swagger-ui/`. If you have specified a different port number while starting the server, you will need to replace `8080` with the appropriate port number in the URL.

Once you have accessed the Swagger documentation, you can explore the API endpoints and make requests using the UI.

## Conclusion

Swagger documentation is a useful tool for documenting your API endpoints and making it easier for other developers to understand how to interact with your project. With go-swagger, generating and displaying Swagger documentation for your Go project is a straightforward process. By following the steps outlined in this README, you can easily generate and display Swagger documentation for your project.