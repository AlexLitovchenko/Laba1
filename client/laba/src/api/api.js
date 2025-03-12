// src/api/api.js

export const fetchProducts = async () => {
    const token = localStorage.getItem('Token');
    if (!token) throw new Error("Нет токена");
  
    const response = await fetch('http://localhost:8080/product/list', {
      method: 'GET',
      headers: {
        'token': token,
        'Content-Type': 'application/json',
      },
    });
  
    if (!response.ok) throw new Error('Ошибка загрузки продуктов');
    return response.json();
  };
  
  export const addToCart = async (productId) => {
    const token = localStorage.getItem('Token');
    const customerId = localStorage.getItem('CustomerID');
    if (!token || !customerId) throw new Error("Нет токена или customer_id");
  
    try {
      const response = await fetch('http://localhost:8080/product/add', {
        method: 'POST',
        headers: {
          'token': token,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ customer_id: customerId, product_id: productId }),
      });
  
      if (!response.ok) throw new Error('Ошибка добавления в корзину');
      const data = await response.json();
      localStorage.setItem('CartID', data); // Сохраняем ID корзины
      console.log(`Товар ${productId} добавлен в корзину. ID корзины: ${data}`);
    } catch (err) {
      console.error(err.message);
    }
  };
  
  export const fetchCart = async () => {
    const token = localStorage.getItem('Token');
    const customerId = localStorage.getItem('CustomerID');
    if (!token || !customerId) throw new Error("Нет токена или customer_id");
  
    const response = await fetch(`http://localhost:8080/cart?customer_id=${customerId}`, {
      method: 'GET',
      headers: {
        'token': token,
        'Content-Type': 'application/json',
      },
    });
  
    if (!response.ok) throw new Error('Ошибка загрузки корзины');
    return response.json();
  };

  // src/api/api.js

// Функция для увеличения количества товара в корзине
export const increaseQuantity = async (cartItemId, cartId) => {
    const token = localStorage.getItem('Token');
    const customerId = localStorage.getItem('CustomerID');
    if (!token || !customerId) throw new Error("Нет токена или customer_id");
  
    const response = await fetch('http://localhost:8080/cart/product-add', {
      method: 'POST',
      headers: {
        'token': token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        cart_item_id: cartItemId,
        cart_id: cartId,
      }),
    });
  
    if (!response.ok) throw new Error('Ошибка увеличения количества товара в корзине');
    return response.json();
  };
  
  // Функция для уменьшения количества товара в корзине
  export const decreaseQuantity = async (cartItemId, cartId) => {
    const token = localStorage.getItem('Token');
    const customerId = localStorage.getItem('CustomerID');
    if (!token || !customerId) throw new Error("Нет токена или customer_id");
  
    const response = await fetch('http://localhost:8080/cart/product-del', {
      method: 'POST',
      headers: {
        'token': token,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        cart_item_id: cartItemId,
        cart_id: cartId,
      }),
    });
  
    if (!response.ok) throw new Error('Ошибка уменьшения количества товара в корзине');
    return response.json();
  };
  