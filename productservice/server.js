const express = require('express');
const morgan = require('express');

const app = express();
const PORT = process.env.PORT || 5000;

app.use(morgan('dev'));
app.use(express.json());
app.use(express.urlencoded({extended : true}));

app.get('/', (req, res) => res.send('hello world'));

app.listen(PORT, () => {
    console.log(`Server is up at http://localhost:${PORT}`)
});
