// src/pages/ProductPage.js

import React, { useEffect, useState } from 'react';
import { fetchProducts, addToCart } from '../api/api';

const ProductPage = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const loadProducts = async () => {
      try {
        const data = await fetchProducts();
        setProducts(data);
      } catch (err) {
        console.error(err.message);
      }
    };

    loadProducts();
  }, []);

  return (
    <div>
      <h2>Список товаров</h2>
      {products.map((product) => (
        <div key={product.ID} style={{ border: '1px solid #ccc', padding: '10px', margin: '10px' }}>
          <h3>{product.Title}</h3>
          <p>{product.Description}</p>
          <p>Цена: {product.Price} руб.</p>
          <button onClick={() => addToCart(product.ID)}>Добавить в корзину</button>
        </div>
      ))}
    </div>
  );
};

export default ProductPage;
