import React from 'react';
import { Link } from 'react-router-dom';

function HomePage() {
  return (
    <div>
      <h2>Добро пожаловать на главную страницу!</h2>
      <div>
        <Link to="/products">
          <button>Продукты</button>
        </Link>
        <Link to="/cart">
          <button>Корзина</button>
        </Link>
      </div>
    </div>
  );
}

export default HomePage;
