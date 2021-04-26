const express = require('express');
const morgan = require('express');
const getProducts = require('./api');

const app = express();
const PORT = process.env.PORT || 5000;

app.use(morgan('dev'));
app.use(express.json());
app.use(express.urlencoded({extended : true}));

app.get('/api/products', getProducts);

app.listen(PORT, () => {
    console.log(`Server is up at http://localhost:${PORT}`);
});
