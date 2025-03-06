# reverse proxy server

this module is part of my challenge **scale-hustle** to clear all the tasks in noob level!!

for context refer: https://github.com/ssenthilnathan3/scale-hustle.git

## running the app

to run the application, use the following command:

make sure you have docker daemon running:

```sh
make run-containers
```

```sh
make run-proxy-server
```

ensure that you have the necessary configurations set in `internal/config.yaml`.

## health check

to perform a health check, use the following command:

```sh
make health_check
```

this will verify that the application is running correctly.