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
            
            // В будущем здесь будет отправка данных на сервер
            console.log('Форма успешно отправлена:', formDataObj);
            
            // Показ сообщения об успешной регистрации
            alert('Регистрация успешно завершена!');
            
            // Редирект на главную страницу (в будущем можно заменить на страницу личного кабинета)
            window.location.href = '../main-page.html';
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