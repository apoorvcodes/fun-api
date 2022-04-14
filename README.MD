
# Backend

Basic Api setup with minima

## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT`

`NEWS_KEY`


## Installation

Clone project through git

```bash
  git clone https://github.com/TeamIndex/Backend.git
  cd Backend
  air 
```
    
## API Reference

#### Get random jokes

```http
  GET /jokes/:type
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `type` | `string` | **Required**. Type of joke |

#### Get lates news based on query

```http
  GET /news/:query
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `query`      | `string` | **Required**. Query of news to fetch |


#### Get meme from subreddit

```http
  GET /meme/:subreddit/:count
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `subreddit`      | `string` | **Required**. The subreddit to fetch meme from |
| `count`      | `string` | **Required**. Number of memes to fetch |


#### Get important historical event of that day

```http
  GET /onthisday/:type
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `type`      | `string` | **Required**. Type of event to fetxh ex. Birth |


#### Get random or all short stories

```http
  GET /story?q
```

| Query | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `q`      | `string` | **Optional**. Whether to fetch all stories or not |


#### Get random fact

```http
  GET /fact
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `_`      | `_` | _ |

## Running Tests

To run tests, run the following command

```bash
  go test 
```

