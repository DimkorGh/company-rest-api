Company Rest API

##### Makefile Commands

| Command                               | Usage                                                      |
|---------------------------------------|------------------------------------------------------------|
| docker.app.start                      | `Start all services`                                       |
| docker.format                         | `Reformats go source code via docker`                      |
| docker.linter.run                     | `Run linter`                                               |
| docker.test.unit                      | `Run unit tests via docker`                                |
| docker.test.all                       | `Run both unit and integration tests via docker`           |
| docker.test.all.coverage.withView     | `Run both unit and integration tests via docker with view` |
| docker.mock.generate FILE={FILE_PATH} | `Generate mock for a specific file via docker`             |

<br>

* In order to execute makefile commands type **make** plus a command from the table above

  make {command}

---

### Testing app endpoints via Postman

* Import *postmanData/collection/4dcdf0e0-52bc-40de-ad10-34bda41a2ade.json* and
  *postmanData/environment/c14aa92a-3df3-40c5-8937-da4e9e6b0164.json* files into postman as you can see below

![import.png](postmanData%2Fscreenshots%2Fimport.png)

* After import postman collection should look like this

![collections.png](postmanData%2Fscreenshots%2Fcollections.png)

* API uses jwt authentication. You can add authorization bearer token to postman requests as you can see below 

![auth.png](postmanData%2Fscreenshots%2Fauth.png)

* **GET /token** endpoint returns an authorization token which lasts 5 minutes. When you execute this call, the specific postman collection parses the token and **adds it automatically** to the headers of all the other requests.

### Notes

* *mongo-init.js* is used in order to add unique constraint to company's collection name field
* .env is pushed to the repository for testing purposes ONLY. In a production environment it should never be tracked
