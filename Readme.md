## MTA Hosting Optimizer

https://drive.google.com/file/d/1E4OYMLtOvEgi-LjhFu3nGIJhPLIQkjQF/view?usp=drive_link

- Installation
  - Clone the Repo.
  - Run `go mod tidy` to install dependencies.
  - Run `go run main.go` to start the server with desired threshold, default is 1.
    - Set the threshold env variable named as `THRESHOLD`.
  - Run `go test -cover ./...` to run the tests (unit + integration).
  
- About
  - Hardcoded ip config mock service with the given data.
  - Basic structure involves controller layer communicating with service layer, I did not include repo layer for sake of simplicity.
    - Ideally in real world scenerio, contoller layer calls service layer and service layer calls repo layer to get the data.
  - Test coverage is 100% for the unit test of hosting and ip config mock service.
  - Contains integration tests as well, under folder named integration.

- Attached Screenshots
  - Code Ran with output
    
    <img width="852" alt="Screenshot 2023-09-08 at 3 06 37 AM" src="https://github.com/iamask22/mta-hosting-optimizer/assets/144318958/227b4414-5a6e-4a16-8dd8-4793383e7689">
  - Test Coverage
    
    <img width="953" alt="Screenshot 2023-09-08 at 3 08 25 AM" src="https://github.com/iamask22/mta-hosting-optimizer/assets/144318958/847d7b11-7296-407f-89b7-fa4ff25c2870">

