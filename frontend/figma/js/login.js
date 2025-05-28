/**
 * JavaScript для страницы входа
 */
document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.querySelector('.auth-form');

    if (loginForm) {
        loginForm.addEventListener('submit', async function(e) {
            e.preventDefault();
            
            // Сбор данных формы
            const email = document.getElementById('email').value;
            const password = document.getElementById('password').value;
            const rememberMe = document.getElementById('remember').checked;
            
            // Проверка валидности Email
            if (!validateEmail(email)) {
                showError('Введите корректный email');
                return;
            }
            
            // Показываем индикатор загрузки
            showLoading();
            
            try {
                // Выполняем запрос на вход
                const response = await window.authApi.login(email, password);
                
                // Сохраняем токены
                window.authApi.saveTokens(response);
                
                // Если пользователь выбрал "Запомнить меня", устанавливаем флаг
                if (rememberMe) {
                    localStorage.setItem('rememberMe', 'true');
                }
                
                // Получаем информацию о пользователе
                const userInfo = await window.authApi.getProfile(response.access_token);
                
                // Сохраняем роль пользователя
                localStorage.setItem('userRole', userInfo.role);
                
                // Скрываем индикатор загрузки
                hideLoading();
                
                // Показываем сообщение об успехе
                showSuccess('Вход выполнен успешно!');
                
                // Перенаправляем пользователя на соответствующую страницу в зависимости от роли
                setTimeout(() => {
                    switch (userInfo.role) {
                        case 'student':
                            window.location.href = 'student/dashboard.html';
                            break;
                        case 'teacher':
                            window.location.href = 'teacher/dashboard.html';
                            break;
                        case 'dean_office':
                            window.location.href = 'staff/dashboard.html';
                            break;
                        default:
                            window.location.href = 'main-page.html';
                    }
                }, 1000);
            } catch (error) {
                hideLoading();
                showError(error.message || 'Неверный email или пароль');
                console.error('Login error:', error);
            }
        });
    }
    
    // Проверка валидности email
    function validateEmail(email) {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(email);
    }
    
    // Показать индикатор загрузки
    function showLoading() {
        let loadingOverlay = document.querySelector('.loading-overlay');
        
        if (!loadingOverlay) {
            loadingOverlay = document.createElement('div');
            loadingOverlay.className = 'loading-overlay';
            loadingOverlay.innerHTML = '<div class="loading-spinner"></div>';
            document.body.appendChild(loadingOverlay);
            
            // Добавляем стили для индикатора загрузки
            const style = document.createElement('style');
            style.textContent = `
                .loading-overlay {
                    position: fixed;
                    top: 0;
                    left: 0;
                    width: 100%;
                    height: 100%;
                    background-color: rgba(0, 0, 0, 0.5);
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    z-index: 9999;
                }
                .loading-spinner {
                    width: 40px;
                    height: 40px;
                    border: 4px solid #f3f3f3;
                    border-top: 4px solid #3498db;
                    border-radius: 50%;
                    animation: spin 1s linear infinite;
                }
                @keyframes spin {
                    0% { transform: rotate(0deg); }
                    100% { transform: rotate(360deg); }
                }
            `;
            document.head.appendChild(style);
        }
        
        loadingOverlay.style.display = 'flex';
    }
    
    // Скрыть индикатор загрузки
    function hideLoading() {
        const loadingOverlay = document.querySelector('.loading-overlay');
        if (loadingOverlay) {
            loadingOverlay.style.display = 'none';
        }
    }
    
    // Показать сообщение об ошибке
    function showError(message) {
        let errorContainer = document.querySelector('.error-container');
        
        if (!errorContainer) {
            errorContainer = document.createElement('div');
            errorContainer.className = 'error-container';
            loginForm.prepend(errorContainer);
            
            // Добавляем стили для контейнера ошибок
            const style = document.createElement('style');
            style.textContent = `
                .error-container {
                    background-color: #f8d7da;
                    color: #721c24;
                    padding: 10px;
                    border-radius: 5px;
                    margin-bottom: 15px;
                    text-align: center;
                }
            `;
            document.head.appendChild(style);
        }
        
        errorContainer.textContent = message;
        errorContainer.style.display = 'block';
        
        // Скрываем сообщение через 5 секунд
        setTimeout(() => {
            errorContainer.style.display = 'none';
        }, 5000);
    }
    
    // Показать сообщение об успехе
    function showSuccess(message) {
        let successContainer = document.querySelector('.success-container');
        
        if (!successContainer) {
            successContainer = document.createElement('div');
            successContainer.className = 'success-container';
            loginForm.prepend(successContainer);
            
            // Добавляем стили для контейнера успеха
            const style = document.createElement('style');
            style.textContent = `
                .success-container {
                    background-color: #d4edda;
                    color: #155724;
                    padding: 10px;
                    border-radius: 5px;
                    margin-bottom: 15px;
                    text-align: center;
                }
            `;
            document.head.appendChild(style);
        }
        
        successContainer.textContent = message;
        successContainer.style.display = 'block';
        
        // Скрываем сообщение через 5 секунд
        setTimeout(() => {
            successContainer.style.display = 'none';
        }, 5000);
    }
}); 