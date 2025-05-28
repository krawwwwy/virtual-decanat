/**
 * API для взаимодействия с бэкендом аутентификации
 */

const API_BASE_URL = '/api/v1';

// Функция для выполнения запросов с заданными параметрами
async function fetchApi(endpoint, method = 'GET', data = null, token = null) {
    const headers = {
        'Content-Type': 'application/json'
    };

    if (token) {
        headers['Authorization'] = `Bearer ${token}`;
    }

    const config = {
        method,
        headers,
        credentials: 'include'
    };

    if (data && (method === 'POST' || method === 'PUT' || method === 'PATCH')) {
        config.body = JSON.stringify(data);
    }

    try {
        const response = await fetch(`${API_BASE_URL}${endpoint}`, config);
        const responseData = await response.json();

        if (!response.ok) {
            throw new Error(responseData.error || 'Произошла ошибка при выполнении запроса');
        }

        return responseData;
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}

// Функция для регистрации пользователя
async function register(userData) {
    return fetchApi('/auth/register', 'POST', userData);
}

// Функция для завершения регистрации пользователя
async function completeRegistration(userData) {
    return fetchApi('/auth/complete-registration', 'POST', userData);
}

// Функция для входа пользователя
async function login(email, password) {
    return fetchApi('/auth/login', 'POST', { email, password });
}

// Функция для обновления токена
async function refreshToken(refreshToken) {
    return fetchApi('/auth/refresh', 'POST', { refresh_token: refreshToken });
}

// Функция для получения профиля пользователя
async function getProfile(token) {
    return fetchApi('/users/profile', 'GET', null, token);
}

// Функция для обновления профиля пользователя
async function updateProfile(userData, token) {
    return fetchApi('/users/profile', 'PUT', userData, token);
}

// Функция для изменения пароля пользователя
async function changePassword(oldPassword, newPassword, token) {
    return fetchApi('/users/change-password', 'POST', { old_password: oldPassword, new_password: newPassword }, token);
}

// Сохранение токенов в localStorage
function saveTokens(tokens) {
    localStorage.setItem('accessToken', tokens.access_token);
    localStorage.setItem('refreshToken', tokens.refresh_token);
    localStorage.setItem('tokenExpires', Date.now() + (tokens.expires_in * 1000));
}

// Получение токена из localStorage
function getToken() {
    return localStorage.getItem('accessToken');
}

// Получение токена обновления из localStorage
function getRefreshToken() {
    return localStorage.getItem('refreshToken');
}

// Очистка токенов из localStorage
function clearTokens() {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    localStorage.removeItem('tokenExpires');
}

// Проверка, авторизован ли пользователь
function isAuthenticated() {
    const token = getToken();
    const expires = localStorage.getItem('tokenExpires');
    
    if (!token || !expires) {
        return false;
    }
    
    return Date.now() < parseInt(expires);
}

// Экспорт функций API
window.authApi = {
    register,
    completeRegistration,
    login,
    refreshToken,
    getProfile,
    updateProfile,
    changePassword,
    saveTokens,
    getToken,
    getRefreshToken,
    clearTokens,
    isAuthenticated
}; 