[Original Tutorial](https://go.dev/doc/tutorial/web-service-gin)

`:id` - Associate the /albums/:id path with the getAlbumByID function. In Gin, the colon preceding an item in the path signifies that the item is a path parameter.

Active the service with the following command:
```bash
go run .
```
- where the '.' signifies to run the go application in the current directory

## Commands
### Get ALL Albums ------------------
```bash
curl http://localhost:8080/albums
```
OR
```bash
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```

### Retrieve a Specific Album ------------------
```bash
curl http://localhost:8080/albums/2
```

### Add One Album ------------------
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```