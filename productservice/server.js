const express = require('express');
const morgan = require('express');
const cors = require('cors');
const getProducts = require('./api');

const app = express();
const PORT = process.env.PORT || 5000;

app.use(morgan('dev'));
app.use(express.json());
app.use(express.urlencoded({extended : true}));
app.use(cors());
app.get('/api/products', getProducts);

app.listen(PORT, () => {
    console.log(`Server is up at http://localhost:${PORT}`);
});
