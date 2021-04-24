const query = `
CREATE TABLE product (
    product_id VARCHAR(10),
    productName VARCHAR(10),
    productDescription VARCHAR(10),
    productPrice INT);
`;

module.exports = query;