#/bin/sh

echo "Going to import data"

ls -la /listingsAndReviews.json
pwd

mongoimport --drop --host "localhost" --port "27017" --db "airbnb" --collection "listingsAndReviews" --file "listingsAndReviews.json" --authenticationDatabase admin -u mongorootuser -p mongorootpw