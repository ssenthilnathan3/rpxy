## Running the Application

To run the application, use the following command:

Make sure you have docker daemon running:

```sh
make run-containers
```

```sh
make run-proxy-server
```

Ensure that you have the necessary configurations set in `internal/config.yaml`.

## Health Check

To perform a health check, use the following command:

```sh
make health_check
```

This will verify that the application is running correctly.