## Flight Calculator
A simple microservice API that can help us understand and track how a particular person's flight path may be queried.

## Prerequisites
Docker: Ensure Docker is installed on your machine. If not, download and install it from the Docker official website.

## Running the Application
Open your terminal and navigate to the root directory of the project.
Run the command `make docker`. This will pull the necessary Docker images and boot up the server.
If the setup is successful, the terminal will display logs indicating the server's activity. In the end you should get a similar output as below

```
...                                                           
flight-calculator-1  | Server started at http://localhost:8080
```

Below is an example curl

```
curl --location 'http://localhost:8080/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
    "flights": [
        {"source": "IND", "destination": "EWR"},
        {"source": "SFO", "destination": "ATL"},
        {"source": "GSO", "destination": "IND"},
        {"source": "ATL", "destination": "GSO"}
    ]
}

'
```

## Testing

Open your terminal and navigate to the root directory of the project.
Run the command `make test`.
Upon successful test execution, the terminal will present a summary of the results, giving insights into the health and robustness of the code. You should be greeted with the following console output

```
➜  flight-calculator git:(main) ✗ make test
?       flights-calculator/cmd/flight-calculator        [no test files]
=== RUN   TestCalculateHandler
--- PASS: TestCalculateHandler (0.00s)
PASS
coverage: 46.5% of statements
ok      flights-calculator/pkg/handler  0.020s  coverage: 74.4% of statements
```

## External Libraries and Decisions

## Versions

Go: 1.21
Docker: 4.5.0
Docker-Compose: 3.4
