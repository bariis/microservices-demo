const { Client } = require('pg');
const { Pool } = require('pg');

const pool = new Pool ({
    user: 'postgres',
    password: '123',
    host: 'db',
    database: 'microservice'
});

module.exports = pool;