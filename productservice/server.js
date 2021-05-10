const express = require('express');
const morgan = require('morgan');
const cors = require('cors');
const api = require('./api');

const app = express();
const PORT = process.env.PORT || 5000;

app.use(morgan('tiny'));
app.use(express.json());
app.use(express.urlencoded({extended : true}));
app.use(cors());

app.get('/api/products', api.getProducts);
app.get('/api/products/:id', api.getProductById);

app.listen(PORT, () => {
    console.log(`Server is up at http://localhost:${PORT}`);
});
