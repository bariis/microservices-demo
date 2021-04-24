const { Client } = require('pg');

const client = new Client ({
    user: 'postgres',
    password: '123',
    host: 'db',
    database: 'microservice'
});

client.on('connect', () => {
    console.log('Connection established');
});

client.on('end', () => {
    console.log('Connection closed');
});

module.exports = client;


