<h1>Just practicing my golang skills.</h1> The base for this project is youtube video. But i dont like this, so Im trying to create my own service based on this idea and some configurations

<h2>To Sum up</h2>
This application is deployed on the server using the Selectel service. But deploy.yaml is correct(Check actions)<br />
It’s impossible to test it in the Internet, because cloud storage is beyond my means.<br />
Therefore, to test this application, you will need to install it locally, add your .env file with all the necessary data (see postgres.go file).<br />
Then, by launching the application using go run /cmd/main.go, check requests through services such as Postman

<h2>Examples of requests</h2>

> **For modificating requests you need to pass BasicAuth params user=igor; password=12345**

> **Add new url with alias request**

URL
localhost:8082/url

Method
`POST`

Request

```json
{
  "url": "https://yandex.ru",
  "alias": "ya"
}

```
Response
```json
{
  "status": "OK",
  "alias": "ya"
}
```

> **Redirect to existing url**

URL
localhost:8082/{alias}

Method
`GET`

Response

```
Opens requested page if exists
May return an error messages
```

> **Delete url by alias request**

URL
localhost:8082/url/{alias}

Method
`DELETE`

Response
```json
{
  "status": "OK",
}
```

<h2>Examples of curl's</h2>

> **Add new url with alias request**

```
curl --location 'localhost:8082/url' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic aWdvcjoxMjM0NQ==' \
--data '{
    "url": "https://yandex.ru",
    "alias": "ya"
}'
```

> **Redirect to existing url**

```
curl --location 'localhost:8082/ya'
```

> **Delete url by alias request**

```
curl --location --request DELETE 'localhost:8082/url/ya' \
--header 'Authorization: Basic aWdvcjoxMjM0NQ=='
```
