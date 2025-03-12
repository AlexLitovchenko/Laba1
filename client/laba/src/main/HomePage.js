// src/pages/HomePage.js

import React from 'react';
import { useNavigate } from 'react-router-dom';

const HomePage = () => {
  const navigate = useNavigate();

  return (
    <div>
      <h2>Добро пожаловать!</h2>
      <button onClick={() => navigate('/products')}>Продукты</button>
      <button onClick={() => navigate('/cart')}>Корзина</button>
    </div>
  );
};

export default HomePage;
