import React, { useState, useEffect } from 'react';
import { getProducts } from '../api/list-products';

const Catalog = () => {

    const [products, setProducts] = useState([]);

    useEffect(() => {
        let mounted = true;
        getProducts()
            .then(list => {
                if(mounted) {
                    setProducts(list);
                }
            })
            return () => mounted = false;
    }, []);

    return (
        <h1>Catalog</h1> 
    );
};

export default Catalog;