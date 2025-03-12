import React from 'react';
import { Routes, Route, useNavigate } from 'react-router-dom';
import LoginPage from './auth/LoginPage';
import RegisterPage from './auth/RegisterPage';
import HomePage from './main/HomePage';
import ProductsPage from './main/ProductsPage';
import CartPage from './main/CartPage';

function App() {
  const navigate = useNavigate(); // Хук для навигации

  return (
    <div>
      <Routes>
        {/* Главная страница с кнопками */}
        <Route 
          path="/" 
          element={
            <>
              <h1>Главная страница</h1>
              <button onClick={() => navigate('/login')}>Логин</button>
              <button onClick={() => navigate('/register')}>Регистрация</button>
            </>
          } 
        />

        {/* Страница логина */}
        <Route path="/login" element={<LoginPage />} />

        {/* Страница регистрации */}
        <Route path="/register" element={<RegisterPage />} />

        <Route path="/home" element={<HomePage />} />
        <Route path="/products" element={<ProductsPage />} />
        <Route path="/cart" element={<CartPage />} />
      </Routes>
    </div>
  );
}

export default App;