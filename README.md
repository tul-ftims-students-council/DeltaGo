# Simple API for Delta website in Go with Fiber and Docker

## Docker
Clone this repository and run:
```
docker-compose up
```

You can then hit the following endpoints:

Method: `POST` \
Route: `/register` \
Body:
```
{
  "name": "patryk", 
  "surname": "makowski", 
  "email": "p.makowski@samorzad.p.lodz.pl", 
  "phoneNumber": "123456789" 
}
```

Method: `GET` \
Route: `/:id/` \
Body: 
```
{
}
```
