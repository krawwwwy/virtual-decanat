/**
 * JavaScript для страниц дашборда
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем, авторизован ли пользователь
    if (!window.authApi.isAuthenticated()) {
        // Если не авторизован, перенаправляем на страницу входа
        window.location.href = '../login-page.html';
        return;
    }

    // Получаем элементы из DOM
    const userNameElement = document.getElementById('userName');
    const logoutBtn = document.getElementById('logoutBtn');
    
    // Обработка кнопки выхода
    if (logoutBtn) {
        logoutBtn.addEventListener('click', function(e) {
            e.preventDefault();
            
            // Очищаем токены и данные пользователя
            window.authApi.clearTokens();
            localStorage.removeItem('userRole');
            localStorage.removeItem('rememberMe');
            
            // Перенаправляем на страницу входа
            window.location.href = '../login-page.html';
        });
    }
    
    // Загружаем информацию о пользователе
    async function loadUserInfo() {
        try {
            // Получаем токен доступа и обновляем его при необходимости
            let token = window.authApi.getToken();
            
            // Если токен истёк, пытаемся обновить его
            if (!token || Date.now() >= parseInt(localStorage.getItem('tokenExpires'))) {
                const refreshToken = window.authApi.getRefreshToken();
                if (refreshToken) {
                    try {
                        const response = await window.authApi.refreshToken(refreshToken);
                        window.authApi.saveTokens(response);
                        token = response.access_token;
                    } catch (error) {
                        // Если не удалось обновить токен, перенаправляем на страницу входа
                        window.location.href = '../login-page.html';
                        return;
                    }
                } else {
                    // Если нет refresh token, перенаправляем на страницу входа
                    window.location.href = '../login-page.html';
                    return;
                }
            }
            
            // Получаем информацию о пользователе
            const userInfo = await window.authApi.getProfile(token);
            
            // Отображаем имя пользователя
            if (userNameElement) {
                const fullName = [userInfo.last_name, userInfo.first_name].filter(Boolean).join(' ');
                userNameElement.textContent = fullName || userInfo.email;
            }
        } catch (error) {
            console.error('Failed to load user info:', error);
            // В случае ошибки можно показать сообщение пользователю
            if (userNameElement) {
                userNameElement.textContent = 'Ошибка загрузки данных';
            }
        }
    }
    
    // Загружаем информацию о пользователе при загрузке страницы
    loadUserInfo();
}); 