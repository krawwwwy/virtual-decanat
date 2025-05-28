/**
 * JavaScript для страницы успеваемости студента
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем авторизацию
    if (!window.authApi.isAuthenticated() || localStorage.getItem('userRole') !== 'student') {
        window.location.href = '../login-page.html';
        return;
    }

    // Кнопка расчета рейтинга
    const calculateRatingBtn = document.querySelector('.rating-position .primary-button');
    if (calculateRatingBtn) {
        calculateRatingBtn.addEventListener('click', async function() {
            try {
                showLoading();
                // Имитация запроса к API
                await new Promise(resolve => setTimeout(resolve, 1500));
                
                // Показываем результат
                hideLoading();
                showNotification('Ваш рейтинг в группе: 5 из 25');
            } catch (error) {
                hideLoading();
                showError('Не удалось рассчитать рейтинг');
            }
        });
    }

    // Кнопка посещаемости
    const attendanceBtn = document.querySelector('button:nth-of-type(1)');
    if (attendanceBtn) {
        attendanceBtn.addEventListener('click', async function() {
            try {
                showLoading();
                // Имитация запроса к API
                await new Promise(resolve => setTimeout(resolve, 1500));
                
                // Показываем результат
                hideLoading();
                showNotification('Ваша посещаемость: 88%');
            } catch (error) {
                hideLoading();
                showError('Не удалось получить данные о посещаемости');
            }
        });
    }

    // Кнопка просмотра оценок
    const gradesBtn = document.querySelector('button:nth-of-type(2)');
    if (gradesBtn) {
        gradesBtn.addEventListener('click', function() {
            // Можно добавить переход на детальную страницу с оценками или показать модальное окно
            showNotification('Отображены все ваши оценки за текущий семестр');
        });
    }

    // Кнопка просмотра задолженностей
    const debtsBtn = document.querySelector('button:nth-of-type(3)');
    if (debtsBtn) {
        debtsBtn.addEventListener('click', function() {
            // Можно добавить переход на детальную страницу с задолженностями или показать модальное окно
            showNotification('Отображены все ваши текущие задолженности');
        });
    }

    // Подсветка строк с задолженностями
    highlightDebtRows();
});

/**
 * Подсветить строки с задолженностями
 */
function highlightDebtRows() {
    const debtTable = document.querySelector('.dashboard-content h2:nth-of-type(4) + .card .data-table');
    if (debtTable) {
        const rows = debtTable.querySelectorAll('tbody tr');
        const currentDate = new Date();
        
        rows.forEach(row => {
            const deadlineCell = row.cells[2];
            if (deadlineCell) {
                const deadline = new Date(deadlineCell.textContent);
                // Если дедлайн близко (меньше 7 дней)
                const daysLeft = Math.ceil((deadline - currentDate) / (1000 * 60 * 60 * 24));
                
                if (daysLeft < 7) {
                    row.style.backgroundColor = '#fff0f0';
                }
            }
        });
    }
}

/**
 * Показать уведомление
 */
function showNotification(message) {
    // Создаем элемент уведомления
    const notification = document.createElement('div');
    notification.className = 'notification';
    notification.textContent = message;
    
    // Стили для уведомления
    notification.style.position = 'fixed';
    notification.style.top = '20px';
    notification.style.right = '20px';
    notification.style.backgroundColor = '#3B19E6';
    notification.style.color = 'white';
    notification.style.padding = '10px 20px';
    notification.style.borderRadius = '5px';
    notification.style.boxShadow = '0 2px 10px rgba(0, 0, 0, 0.2)';
    notification.style.zIndex = '1000';
    
    document.body.appendChild(notification);
    
    // Удаляем уведомление через 3 секунды
    setTimeout(() => {
        notification.remove();
    }, 3000);
}

/**
 * Показать ошибку
 */
function showError(message) {
    // Создаем элемент уведомления
    const notification = document.createElement('div');
    notification.className = 'error';
    notification.textContent = message;
    
    // Стили для уведомления об ошибке
    notification.style.position = 'fixed';
    notification.style.top = '20px';
    notification.style.right = '20px';
    notification.style.backgroundColor = '#e63946';
    notification.style.color = 'white';
    notification.style.padding = '10px 20px';
    notification.style.borderRadius = '5px';
    notification.style.boxShadow = '0 2px 10px rgba(0, 0, 0, 0.2)';
    notification.style.zIndex = '1000';
    
    document.body.appendChild(notification);
    
    // Удаляем уведомление через 3 секунды
    setTimeout(() => {
        notification.remove();
    }, 3000);
}

/**
 * Показать индикатор загрузки
 */
function showLoading() {
    // Создаем элемент индикатора загрузки
    let loadingOverlay = document.createElement('div');
    loadingOverlay.className = 'loading-overlay';
    loadingOverlay.innerHTML = '<div class="loading-spinner"></div>';
    
    // Стили для индикатора загрузки
    loadingOverlay.style.position = 'fixed';
    loadingOverlay.style.top = '0';
    loadingOverlay.style.left = '0';
    loadingOverlay.style.width = '100%';
    loadingOverlay.style.height = '100%';
    loadingOverlay.style.backgroundColor = 'rgba(0, 0, 0, 0.5)';
    loadingOverlay.style.display = 'flex';
    loadingOverlay.style.justifyContent = 'center';
    loadingOverlay.style.alignItems = 'center';
    loadingOverlay.style.zIndex = '9999';
    
    const spinner = loadingOverlay.querySelector('.loading-spinner');
    spinner.style.width = '40px';
    spinner.style.height = '40px';
    spinner.style.border = '4px solid #f3f3f3';
    spinner.style.borderTop = '4px solid #3B19E6';
    spinner.style.borderRadius = '50%';
    spinner.style.animation = 'spin 1s linear infinite';
    
    // Добавляем стили анимации
    const style = document.createElement('style');
    style.textContent = `
        @keyframes spin {
            0% { transform: rotate(0deg); }
            100% { transform: rotate(360deg); }
        }
    `;
    document.head.appendChild(style);
    
    document.body.appendChild(loadingOverlay);
}

/**
 * Скрыть индикатор загрузки
 */
function hideLoading() {
    const loadingOverlay = document.querySelector('.loading-overlay');
    if (loadingOverlay) {
        loadingOverlay.remove();
    }
} 