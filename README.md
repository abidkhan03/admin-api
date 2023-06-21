# admin-api

API for admin website

## Pre-requisites
- [FreeLing](https://freeling-user-manual.readthedocs.io/en/latest/installation/installation-packages/) (`analyze` command)


## Setup
- Use `.env.sample` to create `.env`
- Run the following command to create postgres database inside a docker container
    ```
    docker-compose up -d
    ```

## Run
- Run FreeLing server
    ```
    /usr/bin/analyze -f es.cfg --flush --output json --server --port 50005
    ```

- To build before running, execute the following
    ```
    ./scripts/build.sh
    ```

- Use the following command to run the API on the local machine
    ```
    ./scripts/run.sh
    ```
     This command will automatically run any necessary database migrations before starting the API.
