const knex = require('knex');

module.exports = knex({
    client: 'postgres',
    connection: {
        host: 'db',
        user: 'japon',
        password: '123', 
        database: 'japon'
    },
});