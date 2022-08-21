# Running this app

Firstly, run `docker-compose up`. This will bring up `mongo` on port 27017
and `mongo-express` on 8081. If you see any errors, make sure these ports don't have anything else running on them.

You should now be able to run `cmd/main.go`. 
You should see the following log line in your terminal:
```shell
2022/08/23 00:59:34 purchase was successful
```
This means a purchase has been created and stored in the database.

If you navigate to `localhost:8081` in the browser, you will se a UI that 
will allow you to see the purchase created in Mongo.
