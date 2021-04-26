const pool = require('./database/dbClient');

async function getProducts(req, res) {
    let result;
    try {
        const client = await pool.connect();
        result = await client.query('SELECT * FROM product');
        client.release();
        res.status(200).json(result.rows);
    } catch(err) {
        console.error(`Error fetching product data: ${err}`);
        res.status(500).send();
    } 
}

module.exports = getProducts;