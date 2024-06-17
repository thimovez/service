## Installation

```bash
$ git clone git@github.com:thimovez/service.git
```

## Running Application
Make sure you have Docker, Docker Compose installed.
For start express server and setup db paste this command into your terminal:
This script throw error but its happened becouse initialization db not ready
```bash
make start
```

```bash
docker compose up service
```

| Method   | URL Pattern     | Action                       | Command                                              |
|:---------|:----------------|:-----------------------------|:-----------------------------------------------------|
| POST     | /snippet/view   | Display a specific snippet   | curl -i -X POST http://localhost:4000/snippet/create |
| POST     | /snippet/create | Display a specific snippet   | curl -i -X POST http://localhost:4000/snippet/create |
| POST     | /snippet/create | Create a new snippet         | curl -i -X POST http://localhost:4000/snippet/create |