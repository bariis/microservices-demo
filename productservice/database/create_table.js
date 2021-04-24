const client = require('./dbClient');
const query = require('./query');

(async () => {
    await client.connect();
    try {
        await client.query(query);
        console.log('Product table is created');
    } catch(err) {
        console.error('Error occured creatring Product table: ', err);
    }
    client.end();
})();


