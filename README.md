# Wobot Assignment.

## Tasks
- [ ] Task 1: Create a Restful API in golang.
- [ ] Task 2: Create a Register and Login API.
- [ ] Task 3: Create a Authentication Middleware, using JWT.
- [ ] Task 4: Create a endpoint to upload file by using multipart/form-data.
- [ ] Task 5: Create a endpoint to show available files.
- [ ] Task 6: Create a endpoint to show space available for the user.

## Instructions
- Fork the repository and clone it to your local machine.
- To run the application cd into project folder and do `go run *.go`

## Test Apis
- Use swagger `localhost:8080/swagger/index.html`
- Create User.
    ```
    curl -X POST -H "Content-Type: application/json" -d '{"username": "Sahaj", "password": "123"}' localhost:8080/api/v1/users/create
    ```
- Login User.
    ```
    curl -X POST -H "Content-Type: application/json" -d '{"username": "Sahaj", "password": "123"}' localhost:8080/api/v1/users/login
    ```
    Response.
    ```
    {"token":"<token>"}
    ```

- Upload File.
    ```
    curl -X POST -H "Authorization: Bearer <token>" -F "file=@<book.pdf>" localhost:8080/api/v1/upload
    ```

- Remaining Storage.
    ```
    curl -X GET -H "Authorization: Bearer <token>" -F "file=@<book.pdf>" localhost:8080/api/v1/storage/remaining
    ```

- Get files.
    ```
    curl -X GET -H "Authorization: Bearer <token>" localhost:8080/api/v1/files
    ```
    Response.
    ```
    {
        "files_info":
        [
            {"file_name":"book.pdf","file_size":"7.89 MB"},
            {"file_name":"pic.jpeg","file_size":"196.41 KB"}
        ]
    }
    ```