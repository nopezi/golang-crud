GET assets/_search
{
  "query": {
    "match_all": {}
  }
}

GET /_search
{
  "from": 5,
  "size": 20,
  "query": {
    "match": {
      "user.id": "kimchy"
    }
  }
}