const express = require('express');
const morgan = require('express');

const db = require('./db');

const PORT = process.env.PORT || 5000;
const app = express();

app.use(morgan('dev'));
app.use(express.json());
app.use(express.urlencoded({extended : true}));

app.get('/', (req, res) => res.send('hello world'));

app.get('/users', async (req, res) => {
    const users = await db.select().from('users');
    res.json(users);
});

app.listen(PORT, () => {
    console.log(`Server is up at http://localhost:${PORT}`)
});