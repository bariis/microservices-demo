export async function getProducts() {
    const response = await fetch('http://localhost:5000/api/products');
    const products = await response.json();
    console.log(products);
    return products;
}