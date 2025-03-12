// src/pages/CartPage.js

import React, { useEffect, useState } from 'react';
import { fetchCart, increaseQuantity, decreaseQuantity } from '../api/api';

const CartPage = () => {
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {
    const loadCart = async () => {
      try {
        const data = await fetchCart();
        setCartItems(data);
      } catch (err) {
        console.error(err.message);
      }
    };

    loadCart();
  }, []);

  // Функция для увеличения количества товара
  const handleIncrease = async (cartItemId, cartId) => {
    try {
      // Выполнение запроса на увеличение количества
      await increaseQuantity(cartItemId, cartId);

      // После изменения количества, обновляем корзину
      const updatedCart = await fetchCart();
      setCartItems(updatedCart);
    } catch (err) {
      console.error(err.message);
    }
  };

  // Функция для уменьшения количества товара
  const handleDecrease = async (cartItemId, cartId) => {
    try {
      // Выполнение запроса на уменьшение количества
      await decreaseQuantity(cartItemId, cartId);

      // После изменения количества, обновляем корзину
      const updatedCart = await fetchCart();
      setCartItems(updatedCart);
    } catch (err) {
      console.error(err.message);
    }
  };

  return (
    <div>
      <h2>Корзина</h2>
      {cartItems.length === 0 ? (
        <p>Корзина пуста</p>
      ) : (
        cartItems.map((item) => (
          <div key={item.ID} style={{ border: '1px solid #ccc', padding: '10px', margin: '10px' }}>
            <p><strong>Товар ID:</strong> {item.ProductID}</p>
            <p><strong>Количество:</strong> {item.Quantity}</p>
            <button onClick={() => handleIncrease(item.ID, item.OrderCartID)}>+</button>
            <button onClick={() => handleDecrease(item.ID, item.OrderCartID)}>-</button>
          </div>
        ))
      )}
    </div>
  );
};

export default CartPage;
