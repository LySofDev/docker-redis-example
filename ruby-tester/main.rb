require 'redis'
require 'json'

db = Redis.new host: "redis", port: 6379

db.set "foo", JSON.generate( message: "Set by Ruby" )

puts JSON.parse db.get "foo"
