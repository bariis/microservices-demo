const pool = require('./database/dbClient');

async function getProducts(req, res) {
    let products;
    try {
        const client = await pool.connect();
        products = await client.query('SELECT * FROM product');
        client.release();
        res.status(200).json(products.rows);
    } catch(err) {
        console.error(`Error fetching product data: ${err}`);
        res.status(500).send();
    } 
}

async function getProductById(req, res) {
    let productId = parseInt(req.params.id); 
    try {
        const client = await pool.connect();
        const product = await client.query(`SELECT * FROM product WHERE product_id = ${productId}`);
        client.release();
        res.status(200).json(product.rows);
    } catch(err) {
        console.error(`Error fetching single product: ${err}`);
        res.status(500).send();
    }
}

module.exports = {
    getProducts,
    getProductById
}