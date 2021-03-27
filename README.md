Service for renting a car
-------------------------

#### Requirements:
- [Golgan](https://golang.org/dl/)
- [Docker](https://www.docker.com/)

#### Usage:
To create a new migration use the following command:
```bash
make migrate-create name=create_user_table
```

To apply migrations use the following command:
```bash
make migrate-up
```

To run web server use the following command:
```bash
make web
```