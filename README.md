# BlogAggregatorGo
This is a guided project in Go from [boot.dev](https://www.boot.dev/) where you make a webserver rss posts scrapper.

## Setup
You will need to create a postgres database and create a `.env` file like this one:
```.env
PORT="8080"
CONN="postgres://user:password@host:port/database?sslmode=disable"
```
For executing migrations I used [goose](https://github.com/pressly/goose) while in `./sql/schema/`:

```bash
goose postgres postgres://user:password@host:port/database up
```

or in the root dir:
```bash
goose -dir ./sql/schema postgres postgres://user:password@host:port/database up
```

Finally `go build && ./BlogAggregatorGo`

## Endpoints
### `GET /v1/readiness`:
I made this endpoint to test `RespondWithJSON`.
Response example:
```json
{
  "status": "ok"
}
```

### `GET /v1/err`:
I made this endpoint to test `RespondWithError`.
Response example:
```json
{
  "error": "Internal Server Error"
}
```

### `POST /v1/users`:
Creates an user and returns it's information and the `ApiKey` used for future authorization.

Body example:
```json
{
  "name": "YourName"
}
```
Response example:
```json
{
  "id": "093cce59-51bb-41ad-aa2a-22b678c8baac",
  "created_at": "2024-02-13T16:17:57.203873Z",
  "updated_at": "2024-02-13T16:17:57.203873Z",
  "name": "YourName",
  "api_key": "8c3e80fcc8daca693290ca993ce8ded0a11317044918408f963d8bee6acdc9b1"
}
```

### `GET /v1/users`:
Returns information about user authenticated by `ApiKey`.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

Response example:
```json
{
  "id": "093cce59-51bb-41ad-aa2a-22b678c8baac",
  "created_at": "2024-02-13T16:17:57.203873Z",
  "updated_at": "2024-02-13T16:17:57.203873Z",
  "name": "YourName",
  "api_key": "8c3e80fcc8daca693290ca993ce8ded0a11317044918408f963d8bee6acdc9b1"
}
```

### `POST /v1/feed`:
Creates a new feed.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

Body example:
```json
{
  "name": "Feed name",
  "url": "urlofthefeed.com"
}
```
Response example:
```json
{
  "feed": {
    "id": "900c5010-7885-48a7-9ed0-f20466eb76ff",
    "name": "Feed name",
    "url": "urlofthefeed.com",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-13T16:25:03.656305Z",
    "updated_at": "2024-02-13T16:25:03.656305Z",
    "last_updated_at": {
      "Time": "0001-01-01T00:00:00Z",
      "Valid": false
    }
  },
  "feed_follow": {
    "id": "b3017f76-6725-4e05-a731-6c93cefd3f80",
    "feed_id": "900c5010-7885-48a7-9ed0-f20466eb76ff",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-13T16:25:03.660531Z",
    "updated_at": "2024-02-13T16:25:03.660531Z"
  }
}
```

### `GET /v1/feeds`:
Returns all feeds.

Response example:
```json
[
  {
    "id": "741d3546-51bc-42a9-9d1d-58f28b0a40cf",
    "name": "Amazing random feed",
    "url": "https://amazingrandomfeed.com/",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-11T17:06:53.130634Z",
    "updated_at": "2024-02-13T16:25:33.438107Z",
    "last_updated_at": {
      "Time": "2024-02-13T16:25:33.438107Z",
      "Valid": true
    }
  },
  {
    "id": "900c5010-7885-48a7-9ed0-f20466eb76ff",
    "name": "Feed name",
    "url": "urlofthefeed.com",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-13T16:25:03.656305Z",
    "updated_at": "2024-02-13T16:25:33.43813Z",
    "last_updated_at": {
      "Time": "2024-02-13T16:25:33.43813Z",
      "Valid": true
    }
  },
  {
    "id": "ab6b65f3-a66f-48fb-8852-6a25c014697f",
    "name": "Another random feed",
    "url": "https://anotherrandomfeed.com/index.xml",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-10T22:34:44.652131Z",
    "updated_at": "2024-02-13T16:25:33.441105Z",
    "last_updated_at": {
      "Time": "2024-02-13T16:25:33.441105Z",
      "Valid": true
    }
  }
]
```

### `POST /v1/feed_follows`:
Follows an existing feed.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

Body example:
```json
{
  "feed_id": "ab6b65f3-a66f-48fb-8852-6a25c014697f"
}
```
Response example:
```json
{
  "id": "20635fe9-c339-4c93-8df4-a798de671cc7",
  "feed_id": "ab6b65f3-a66f-48fb-8852-6a25c014697f",
  "user_id": "093cce59-51bb-41ad-aa2a-22b678c8baac",
  "created_at": "2024-02-13T16:34:49.747039Z",
  "updated_at": "2024-02-13T16:34:49.747039Z"
}
```

### `DELETE /v1/feed_follows/{feedFollowID}`:
Deletes a feed_follows by feedFollowID.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

### `GET /v1/feed_follows`:
Return every feed follow of the user.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

Response example:
```json
[
  {
    "id": "f84075a2-9325-4ef3-b4c2-739b34731658",
    "feed_id": "741d3546-51bc-42a9-9d1d-58f28b0a40cf",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-11T17:06:53.136533Z",
    "updated_at": "2024-02-11T17:06:53.136533Z"
  },
  {
    "id": "b3017f76-6725-4e05-a731-6c93cefd3f80",
    "feed_id": "900c5010-7885-48a7-9ed0-f20466eb76ff",
    "user_id": "5347d2c0-b6b9-4317-8119-b970066e6ead",
    "created_at": "2024-02-13T16:25:03.660531Z",
    "updated_at": "2024-02-13T16:25:03.660531Z"
  }
]
```

### `GET /v1/posts`:
Return posts of the user feed.

Authentication: Bearer Token
```
Authorization: ApiKey YourAPIKey
```

Body example:
```json
{
  "limit": 2
}
```
Response example:
```json
[
  {
    "id": "04373594-bd87-4ca8-8fb0-fc5479a4642d",
    "created_at": "2024-02-13T14:31:39.90568Z",
    "updated_at": "2024-02-13T14:31:39.90568Z",
    "title": "RandomTitle",
    "url": "https://randomfeed.com/randomTitle/",
    "description": {
      "String": "\n      \n      \n      \n      \n      \n      \n    ",
      "Valid": true
    },
    "published_at": {
      "Time": "2024-01-28T00:00:00Z",
      "Valid": true
    },
    "feed_id": "ab6b65f3-a66f-48fb-8852-6a25c014697f"
  },
  {
    "id": "952a783e-f980-454c-896a-b883f3ef8623",
    "created_at": "2024-02-13T14:31:39.900174Z",
    "updated_at": "2024-02-13T14:31:39.900174Z",
    "title": "second random title",
    "url": "https://secondrandomfeed.com/title",
    "description": {
      "String": "\n      \n      \n      \n      \n      \n      \n    ",
      "Valid": true
    },
    "published_at": {
      "Time": "2024-01-31T00:00:00Z",
      "Valid": true
    },
    "feed_id": "ab6b65f3-a66f-48fb-8852-6a25c014697f"
  }
]
```