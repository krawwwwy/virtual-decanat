/**
 * Виртуальный Деканат - Основной JavaScript файл
 */

// Дожидаемся загрузки DOM
document.addEventListener('DOMContentLoaded', () => {
    console.log('Виртуальный Деканат - загрузка завершена');
    
    // Добавляем анимацию для декоративных элементов
    animateDecorations();
});

/**
 * Функция для анимации декоративных элементов на странице
 */
function animateDecorations() {
    const leftDecoration = document.querySelector('.decoration-element.left');
    const rightDecoration = document.querySelector('.decoration-element.right');
    
    if (leftDecoration && rightDecoration) {
        // Добавляем небольшую пульсацию с разной скоростью
        setInterval(() => {
            leftDecoration.style.transform = `scale(${1 + Math.sin(Date.now() * 0.001) * 0.1})`;
        }, 50);
        
        setInterval(() => {
            rightDecoration.style.transform = `scale(${1 + Math.sin(Date.now() * 0.0015) * 0.1})`;
        }, 50);
    }
}

/**
 * Функция для перехода к странице входа
 */
function navigateToLogin() {
    window.location.href = 'login.html';
}

/**
 * Функция для перехода к странице регистрации
 */
function navigateToRegistration() {
    window.location.href = 'registration.html';
}

/**
 * Функция для перехода к странице абитуриентов
 */
function navigateToApplicant() {
    window.location.href = 'applicant.html';
} 