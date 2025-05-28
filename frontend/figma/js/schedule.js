/**
 * JavaScript для страницы расписания студента
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем авторизацию
    if (!window.authApi.isAuthenticated() || localStorage.getItem('userRole') !== 'student') {
        window.location.href = '../login-page.html';
        return;
    }

    // Переключение между видами расписания (день/неделя/месяц)
    const tabs = document.querySelectorAll('.tab');
    tabs.forEach(tab => {
        tab.addEventListener('click', function() {
            // Убираем активный класс у всех табов
            tabs.forEach(t => t.classList.remove('active'));
            
            // Добавляем активный класс текущему табу
            this.classList.add('active');
            
            // В реальном приложении здесь был бы запрос к API для загрузки соответствующего вида расписания
            showNotification(`Переключено на вид: ${this.textContent}`);
        });
    });

    // Обработка кнопки "Сегодня"
    const todayButton = document.querySelector('.today-button');
    if (todayButton) {
        todayButton.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                showNotification('Расписание на сегодня загружено');
            }, 800);
        });
    }

    // Обработка кнопки "Завтра"
    const tomorrowButton = document.querySelector('.tomorrow-button');
    if (tomorrowButton) {
        tomorrowButton.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                showNotification('Расписание на завтра загружено');
            }, 800);
        });
    }

    // Обработка кнопок деталей занятия
    const detailButtons = document.querySelectorAll('.details-button');
    detailButtons.forEach(button => {
        button.addEventListener('click', function() {
            // Получаем название предмета
            const subjectTitle = this.closest('.schedule-item').querySelector('.subject-title').textContent;
            showSubjectDetails(subjectTitle);
        });
    });

    /**
     * Показать детали предмета
     */
    function showSubjectDetails(subject) {
        // В реальном приложении здесь был бы запрос к API для получения деталей предмета
        // Для демонстрации просто показываем модальное окно с информацией
        
        // Создаем модальное окно
        const modal = document.createElement('div');
        modal.className = 'modal';
        
        // Стили для модального окна
        modal.style.position = 'fixed';
        modal.style.top = '0';
        modal.style.left = '0';
        modal.style.width = '100%';
        modal.style.height = '100%';
        modal.style.backgroundColor = 'rgba(0, 0, 0, 0.6)';
        modal.style.display = 'flex';
        modal.style.justifyContent = 'center';
        modal.style.alignItems = 'center';
        modal.style.zIndex = '9999';
        
        // Создаем контент модального окна
        const modalContent = document.createElement('div');
        modalContent.className = 'modal-content';
        modalContent.style.backgroundColor = '#FFFFFF';
        modalContent.style.padding = '24px';
        modalContent.style.borderRadius = '12px';
        modalContent.style.width = '480px';
        modalContent.style.maxWidth = '90%';
        modalContent.style.maxHeight = '80vh';
        modalContent.style.overflowY = 'auto';
        modalContent.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.15)';
        
        // Заголовок модального окна
        const modalTitle = document.createElement('h3');
        modalTitle.textContent = subject;
        modalTitle.style.fontFamily = 'Inter, sans-serif';
        modalTitle.style.fontWeight = '700';
        modalTitle.style.fontSize = '24px';
        modalTitle.style.marginTop = '0';
        modalTitle.style.marginBottom = '16px';
        
        // Информация о предмете
        const subjectInfo = document.createElement('div');
        subjectInfo.innerHTML = `
            <p><strong>Преподаватель:</strong> Иванов Иван Иванович</p>
            <p><strong>Аудитория:</strong> 301</p>
            <p><strong>Тип занятия:</strong> Лекция</p>
            <p><strong>Материалы:</strong> <a href="#" style="color: #3B19E6;">Презентация</a>, <a href="#" style="color: #3B19E6;">Методичка</a></p>
            <p><strong>Задания:</strong> <a href="#" style="color: #3B19E6;">Практическое задание №3</a></p>
        `;
        subjectInfo.style.fontFamily = 'Inter, sans-serif';
        subjectInfo.style.fontSize = '16px';
        subjectInfo.style.lineHeight = '1.5';
        
        // Кнопка закрытия
        const closeButton = document.createElement('button');
        closeButton.textContent = 'Закрыть';
        closeButton.style.backgroundColor = '#3B19E6';
        closeButton.style.color = '#FFFFFF';
        closeButton.style.border = 'none';
        closeButton.style.padding = '8px 16px';
        closeButton.style.borderRadius = '8px';
        closeButton.style.marginTop = '24px';
        closeButton.style.cursor = 'pointer';
        closeButton.style.fontFamily = 'Inter, sans-serif';
        closeButton.style.fontWeight = '500';
        
        // Добавляем элементы в модальное окно
        modalContent.appendChild(modalTitle);
        modalContent.appendChild(subjectInfo);
        modalContent.appendChild(closeButton);
        modal.appendChild(modalContent);
        
        // Добавляем модальное окно на страницу
        document.body.appendChild(modal);
        
        // Обработчик для кнопки закрытия
        closeButton.addEventListener('click', function() {
            modal.remove();
        });
        
        // Закрытие по клику вне модального окна
        modal.addEventListener('click', function(event) {
            if (event.target === modal) {
                modal.remove();
            }
        });
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
        
        const spinner = document.createElement('div');
        spinner.className = 'loading-spinner';
        spinner.style.width = '40px';
        spinner.style.height = '40px';
        spinner.style.border = '4px solid #f3f3f3';
        spinner.style.borderTop = '4px solid #3B19E6';
        spinner.style.borderRadius = '50%';
        spinner.style.animation = 'spin 1s linear infinite';
        
        loadingOverlay.appendChild(spinner);
        
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

    // Инициализация датапикера для выбора даты (заглушка)
    document.querySelectorAll('.current-date').forEach(dateElement => {
        dateElement.addEventListener('click', function() {
            showNotification('Выберите дату (функциональность в разработке)');
        });
    });
}); 