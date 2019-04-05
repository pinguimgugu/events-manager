# Service provide a publisher/subscriber architecture exposing resources to create and consuption of events by http protocol

- Payload of events should be have follow example content
   `{"name":"sample", "metadata": "my_content"}`

- Launch service
    Make sure you have been installed docker-compose and docker engine

# build base image
    docker build . -t go-tools

# getting up external services and running application
    ./application.sh serve

# running application
    ./application.sh run


Use example:

Create event
```
    curl -X POST \
    http://localhost:7000/events/v1/events/ \
    -H 'Content-Type: application/json' \
    -d '{"name":"sample", "metadata": {"content": {"my_event_id": 1, "my_event_data":"test"}}}' 
```

Consumer one
```
    curl -i -X POST \
    http://localhost:7000/events/v1/streaming/ \
    -H "Content-Type: application/json" \
    -d '{"consumer_name": "app-consumer-one", "event_name": "sample"}'
```

Consumer two
```
    curl -i -X POST \
    http://localhost:7000/events/v1/streaming/ \
    -H "Content-Type: application/json" \
    -d '{"consumer_name": "app-consumer-two", "event_name": "sample"}'
```