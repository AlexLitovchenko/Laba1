import React, { useState } from 'react';

function RegisterPage() {
  // Состояние для формы
  const [name, setName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  // Функция для обработки отправки формы
  const handleSubmit = async (e) => {
    e.preventDefault();

    // Включаем лоадер
    setLoading(true);
    setError('');

    // Данные для отправки
    const data = {
      name,
      last_name: lastName,
      email,
      password,
    };

    try {
      // Отправка запроса на локальный сервер
      const response = await fetch('http://localhost:8080/register', { // Убедитесь, что это правильный адрес вашего локального сервера
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      // Если сервер вернул ошибку, выбрасываем исключение
      if (!response.ok) {
        throw new Error('Ошибка при отправке данных');
      }

      // Получаем ответ от сервера
      const result = await response.json();
      console.log('Регистрация успешна:', result);

      // Очистка формы
      setName('');
      setLastName('');
      setEmail('');
      setPassword('');
    } catch (err) {
      console.error('Ошибка:', err);
      setError('Ошибка при отправке данных');
    } finally {
      setLoading(false); // Отключаем лоадер
    }
  };

  return (
    <div>
      <h2>Форма Регистрации</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Имя:
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </label>
        <br />
        <label>
          Фамилия:
          <input
            type="text"
            value={lastName}
            onChange={(e) => setLastName(e.target.value)}
            required
          />
        </label>
        <br />
        <label>
          Email:
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </label>
        <br />
        <label>
          Пароль:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </label>
        <br />
        <button type="submit" disabled={loading}>
          {loading ? 'Отправка...' : 'Зарегистрироваться'}
        </button>
      </form>

      {/* Показываем ошибку, если она есть */}
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </div>
  );
}

export default RegisterPage;
