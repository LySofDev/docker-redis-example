const db = require( 'redis' ).createClient( '6379', 'redis' );

db.set( 'foo', JSON.stringify( { message: "Set by Node." } ) );

db.get( 'foo', ( error, value ) => {
  console.log( "Data:", JSON.parse( value ) );
  process.exit( 0 );
} );
