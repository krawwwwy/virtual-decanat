/**
 * Общий JavaScript для страниц завершения регистрации
 */
document.addEventListener('DOMContentLoaded', function() {
    // Обработка отправки формы
    const form = document.querySelector('.registration-form');
    
    if (form) {
        form.addEventListener('submit', function(e) {
            e.preventDefault();
            
            // Проверка валидации формы
            if (!validateForm(form)) {
                return;
            }
            
            // Сбор данных формы
            const formData = new FormData(form);
            const formDataObj = {};
            
            formData.forEach((value, key) => {
                formDataObj[key] = value;
            });
            
            // Получение user_id из localStorage или URL
            const userId = localStorage.getItem('userId') || getQueryParam('userId');
            const role = getQueryParam('role') || localStorage.getItem('userRole');
            
            if (!userId || !role) {
                showError(form, 'Не удалось получить данные пользователя. Пожалуйста, вернитесь на страницу регистрации.');
                return;
            }
            
            // Разбиваем ФИО на составляющие
            const fullNameParts = formDataObj.fio ? formDataObj.fio.split(' ') : ['', '', ''];
            const lastName = fullNameParts[0] || '';
            const firstName = fullNameParts[1] || '';
            const middleName = fullNameParts[2] || '';
            
            // Формирование данных для запроса
            const requestData = {
                user_id: parseInt(userId),
                first_name: firstName,
                last_name: lastName,
                middle_name: middleName,
                role: role,
                birth_date: formDataObj.birth_date || "",
                phone: formDataObj.phone || "",
            };
            
            // Добавление специфичных полей в зависимости от роли
            if (role === 'student') {
                requestData.group = formDataObj.group || "";
                requestData.student_id = formDataObj.student_id || "";
                
                // Лог для отладки
                console.log('Отправляемые данные формы:', formDataObj);
                console.log('Отправляемые данные запроса:', requestData);
            }
            
            // Отправка данных на сервер
            completeRegistration(requestData);
        });
    }
    
    // Добавление имен для полей ввода (для корректной отправки формы)
    const inputs = document.querySelectorAll('.form-group input');
    
    inputs.forEach(input => {
        if (!input.hasAttribute('name')) {
            // Устанавливаем имя поля на основе placeholder
            const placeholder = input.getAttribute('placeholder');
            if (placeholder) {
                const name = placeholder.toLowerCase().replace(/\s+/g, '_').replace(/[^a-z0-9_]/g, '');
                input.setAttribute('name', name);
            }
        }
    });
    
    // Отправка данных на сервер для завершения регистрации
    async function completeRegistration(userData) {
        try {
            showLoading();
            console.log('Отправляемые данные:', userData); // Для отладки
            
            // Вызов API для завершения регистрации
            const response = await window.authApi.completeRegistration(userData);
            
            // Показ сообщения об успешной регистрации
            hideLoading();
            showSuccess('Регистрация успешно завершена!');
            
            // Редирект на страницу входа
            setTimeout(() => {
                window.location.href = '../login-page.html';
            }, 2000);
        } catch (error) {
            hideLoading();
            showError(form, error.message || 'Произошла ошибка при завершении регистрации.');
            console.error('Registration error:', error);
        }
    }
    
    // Показать индикатор загрузки
    function showLoading() {
        // Находим или создаем элемент загрузки
        let loadingElement = document.querySelector('.loading-overlay');
        
        if (!loadingElement) {
            loadingElement = document.createElement('div');
            loadingElement.className = 'loading-overlay';
            loadingElement.innerHTML = '<div class="loading-spinner"></div>';
            document.body.appendChild(loadingElement);
            
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
        
        loadingElement.style.display = 'flex';
    }
    
    // Скрыть индикатор загрузки
    function hideLoading() {
        const loadingElement = document.querySelector('.loading-overlay');
        if (loadingElement) {
            loadingElement.style.display = 'none';
        }
    }
    
    // Показать сообщение об успехе
    function showSuccess(message) {
        // Находим или создаем элемент сообщения
        let messageElement = document.querySelector('.success-message');
        
        if (!messageElement) {
            messageElement = document.createElement('div');
            messageElement.className = 'success-message';
            document.body.appendChild(messageElement);
            
            // Добавляем стили для сообщения
            const style = document.createElement('style');
            style.textContent = `
                .success-message {
                    position: fixed;
                    top: 20px;
                    left: 50%;
                    transform: translateX(-50%);
                    background-color: #4CAF50;
                    color: white;
                    padding: 15px 20px;
                    border-radius: 5px;
                    z-index: 1000;
                    text-align: center;
                    font-weight: 500;
                }
            `;
            document.head.appendChild(style);
        }
        
        messageElement.textContent = message;
        messageElement.style.display = 'block';
        
        // Скрываем сообщение через 5 секунд
        setTimeout(() => {
            messageElement.style.display = 'none';
        }, 5000);
    }
    
    // Получение параметра из URL
    function getQueryParam(name) {
        const urlParams = new URLSearchParams(window.location.search);
        return urlParams.get(name);
    }
    
    // Функция для валидации формы
    function validateForm(form) {
        let isValid = true;
        const requiredInputs = form.querySelectorAll('input[required]');
        
        requiredInputs.forEach(input => {
            if (!input.value.trim()) {
                isValid = false;
                showError(input, 'Это поле обязательно для заполнения');
            } else {
                clearError(input);
                
                // Дополнительная валидация по типу поля
                if (input.type === 'email' && !validateEmail(input.value)) {
                    isValid = false;
                    showError(input, 'Введите корректный email');
                } else if (input.type === 'tel' && !validatePhone(input.value)) {
                    isValid = false;
                    showError(input, 'Введите корректный номер телефона');
                }
            }
        });
        
        // Проверка специфических полей по имени
        const passwordFields = form.querySelectorAll('input[name="password"], input[name="confirm_password"]');
        if (passwordFields.length === 2) {
            if (passwordFields[0].value !== passwordFields[1].value) {
                isValid = false;
                showError(passwordFields[1], 'Пароли не совпадают');
            }
        }
        
        return isValid;
    }
    
    // Показать ошибку под полем ввода
    function showError(input, message) {
        // Если input - это форма, показываем общую ошибку
        if (input.tagName === 'FORM') {
            // Создаем общее сообщение об ошибке
            let errorContainer = document.querySelector('.form-error-container');
            
            if (!errorContainer) {
                errorContainer = document.createElement('div');
                errorContainer.className = 'form-error-container';
                errorContainer.style.backgroundColor = '#f8d7da';
                errorContainer.style.color = '#721c24';
                errorContainer.style.padding = '10px';
                errorContainer.style.borderRadius = '5px';
                errorContainer.style.marginBottom = '20px';
                errorContainer.style.textAlign = 'center';
                
                input.prepend(errorContainer);
            }
            
            errorContainer.textContent = message;
            errorContainer.style.display = 'block';
            return;
        }
        
        // Удаление предыдущего сообщения об ошибке
        clearError(input);
        
        const errorElement = document.createElement('div');
        errorElement.className = 'error-message';
        errorElement.textContent = message;
        errorElement.style.color = '#e74c3c';
        errorElement.style.fontSize = '12px';
        errorElement.style.marginTop = '5px';
        
        // Добавление стиля ошибки для поля ввода
        input.style.borderColor = '#e74c3c';
        
        // Вставка сообщения об ошибке после поля ввода
        input.parentNode.appendChild(errorElement);
    }
    
    // Очистить ошибку
    function clearError(input) {
        // Если input - это форма, очищаем общую ошибку
        if (input.tagName === 'FORM') {
            const errorContainer = document.querySelector('.form-error-container');
            if (errorContainer) {
                errorContainer.style.display = 'none';
            }
            return;
        }
        
        const parent = input.parentNode;
        const errorElement = parent.querySelector('.error-message');
        
        if (errorElement) {
            parent.removeChild(errorElement);
        }
        
        input.style.borderColor = '';
    }
    
    // Валидация email
    function validateEmail(email) {
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return re.test(email);
    }
    
    // Валидация телефона
    function validatePhone(phone) {
        // Проверка на 11 цифр (российский формат)
        const re = /^\d{11}$/;
        return re.test(phone);
    }
}); 