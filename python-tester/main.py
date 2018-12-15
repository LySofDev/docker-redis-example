from redis import Redis
from json import dumps, loads

db = Redis( host="redis", port="6379" )

db.set( 'foo', dumps( { "message": "Set by Python." } ) )

print( loads( db.get( 'foo' ) ) )
