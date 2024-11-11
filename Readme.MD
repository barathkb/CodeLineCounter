
# Go Code Line Counter Loader

This project is a Go application that returns overall number of lines along with number lines in each file.

## Project Structure

- `main.go` : Contains server start logic.
- `utils.go` : Utility functions to support the main operations.
- `http_handler.go` : Contains the http request handler function.
- `constants.go` : Contains constant
- `LineCounter/` : Directory containing file analyser logic
- `DTO/` : Contains Object Definitions
- `Counter/` : Directory containing factory method for counting lines based on language.
- `Sample Test Files/` : Directory contains sample files for testing.

### Requirements
- Go 1.23 or higher

### Running the Application
1. Install Dependencies
```azure
    go mod tidy
```
2. Run Project
```azure
    go run main.go
```
3. Sample Request
```
curl --location 'http://localhost:8080/count?folderPath=folder-path'
```
4. Sample Response

```json
{
    "body": {
        "overall": {
            "file_name": "",
            "total_lines": 33,
            "blank_lines": 4,
            "comment_lines": 7,
            "code_lines": 17,
            "import_lines": 2,
            "variable_lines": 3
        },
        "files": [
            {
                "file_name": "sample.c",
                "total_lines": 10,
                "blank_lines": 1,
                "comment_lines": 2,
                "code_lines": 5,
                "import_lines": 1,
                "variable_lines": 1
            },
            {
                "file_name": "sample.go",
                "total_lines": 13,
                "blank_lines": 2,
                "comment_lines": 2,
                "code_lines": 7,
                "import_lines": 1,
                "variable_lines": 1
            },
            {
                "file_name": "sample.py",
                "total_lines": 10,
                "blank_lines": 1,
                "comment_lines": 3,
                "code_lines": 5,
                "import_lines": 0,
                "variable_lines": 1
            }
        ]
    },
    "message": "Successfully counted lines for provided folder"
}
```

### Testing
```azure
    cd Counter
    go test
```


### License

This project is licensed under the MIT License.