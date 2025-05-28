/**
 * JavaScript для страницы расписания студента
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем авторизацию
    if (!window.authApi.isAuthenticated() || localStorage.getItem('userRole') !== 'student') {
        window.location.href = '../login-page.html';
        return;
    }

    // Переключение между видами расписания (день/неделя/календарь)
    const tabs = document.querySelectorAll('.tab');
    tabs.forEach(tab => {
        tab.addEventListener('click', function() {
            // Убираем активный класс у всех табов
            tabs.forEach(t => t.classList.remove('active'));
            
            // Добавляем активный класс текущему табу
            this.classList.add('active');
            
            // Получаем вид расписания
            const view = this.dataset.view;
            
            // Скрываем все виды и показываем выбранный
            document.querySelectorAll('.schedule-view').forEach(v => {
                v.classList.remove('active-view');
            });
            document.querySelector(`.${view}-view`).classList.add('active-view');
            
            showNotification(`Вид расписания: ${this.textContent}`);
        });
    });

    // Обработка кнопки "Сегодня" в дневном виде
    const todayButton = document.querySelector('.today-button');
    if (todayButton) {
        todayButton.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                showNotification('Расписание на сегодня загружено');
                
                // Переключаем на сегодняшний день (28 мая)
                document.querySelector('.day-view .current-date').textContent = 'Вторник, 28 мая';
            }, 800);
        });
    }

    // Обработка кнопки "Текущая неделя" в недельном виде
    const thisWeekButton = document.querySelector('.this-week-button');
    if (thisWeekButton) {
        thisWeekButton.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                showNotification('Расписание на текущую неделю загружено');
                
                // Переключаем на текущую неделю
                document.querySelector('.week-title').textContent = '27 мая - 2 июня';
                
                // Подсвечиваем активный день (28 мая - вторник)
                document.querySelectorAll('.week-day').forEach(day => {
                    day.classList.remove('active');
                    day.querySelector('.day-indicator').classList.remove('active');
                });
                const activeDay = document.querySelector('.week-day[data-date="28.05"]');
                activeDay.classList.add('active');
                activeDay.querySelector('.day-indicator').classList.add('active');
                
                // Показываем расписание на активный день
                showDaySchedule('28.05');
            }, 800);
        });
    }

    // Обработка выбора дня в недельном виде
    const weekDays = document.querySelectorAll('.week-day');
    weekDays.forEach(day => {
        day.addEventListener('click', function() {
            // Убираем активный класс у всех дней
            weekDays.forEach(d => {
                d.classList.remove('active');
                d.querySelector('.day-indicator').classList.remove('active');
            });
            
            // Добавляем активный класс текущему дню
            this.classList.add('active');
            this.querySelector('.day-indicator').classList.add('active');
            
            // Показываем расписание на выбранный день
            const dateStr = this.dataset.date;
            showDaySchedule(dateStr);
        });
    });

    /**
     * Показать расписание для выбранного дня в недельном виде
     */
    function showDaySchedule(dateStr) {
        showLoading();
        
        // Имитация загрузки
        setTimeout(() => {
            hideLoading();
            
            // Скрываем все расписания
            document.querySelectorAll('.week-schedule .schedule-list').forEach(list => {
                list.style.display = 'none';
            });
            
            // Скрываем сообщение о пустом дне
            document.querySelector('.empty-day-message').style.display = 'none';
            
            // Если есть расписание на этот день, показываем его
            const scheduleList = document.querySelector(`.week-schedule .schedule-list[data-date="${dateStr}"]`);
            if (scheduleList) {
                scheduleList.style.display = 'flex';
                showNotification(`Расписание на ${dateStr} загружено`);
            } else {
                // Если нет расписания, показываем сообщение
                document.querySelector('.empty-day-message').style.display = 'flex';
                showNotification(`На ${dateStr} занятия отсутствуют`);
            }
        }, 400);
    }

    // Обработка кнопок в календарном виде
    const calendarToday = document.querySelector('.calendar-today');
    if (calendarToday) {
        calendarToday.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                
                // Переключаем на сегодняшний день (28 мая)
                showCalendarDate('28');
                showNotification('Календарь на текущий месяц загружен');
            }, 800);
        });
    }

    // Обработка кнопок навигации по календарю
    const calendarPrev = document.querySelector('.calendar-prev');
    const calendarNext = document.querySelector('.calendar-next');
    
    if (calendarPrev) {
        calendarPrev.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                document.querySelector('.calendar-title').textContent = 'Апрель 2025';
                showNotification('Загружен предыдущий месяц');
            }, 800);
        });
    }
    
    if (calendarNext) {
        calendarNext.addEventListener('click', function() {
            showLoading();
            
            // Имитация запроса к API
            setTimeout(() => {
                hideLoading();
                document.querySelector('.calendar-title').textContent = 'Июнь 2025';
                showNotification('Загружен следующий месяц');
            }, 800);
        });
    }

    // Обработка выбора дня в календаре
    const calendarDays = document.querySelectorAll('.calendar-day');
    calendarDays.forEach(day => {
        day.addEventListener('click', function() {
            // Проверяем, не принадлежит ли день к другому месяцу
            if (this.classList.contains('prev-month') || this.classList.contains('next-month')) {
                showNotification('Выберите день текущего месяца');
                return;
            }
            
            // Убираем активный класс у всех дней
            calendarDays.forEach(d => d.classList.remove('active-day'));
            
            // Добавляем активный класс текущему дню
            this.classList.add('active-day');
            
            // Показываем расписание на выбранный день
            const day = this.textContent;
            showCalendarDate(day);
        });
    });

    /**
     * Показать расписание на выбранную дату в календаре
     */
    function showCalendarDate(day) {
        showLoading();
        
        // Имитация загрузки
        setTimeout(() => {
            hideLoading();
            
            // Выделяем активный день
            calendarDays.forEach(d => d.classList.remove('active-day'));
            const selectedDay = Array.from(calendarDays).find(d => 
                !d.classList.contains('prev-month') && 
                !d.classList.contains('next-month') && 
                d.textContent === day
            );
            
            if (selectedDay) {
                selectedDay.classList.add('active-day');
            }
            
            // Если день содержит занятия (16, 24, 30 мая и 28 мая)
            if (['16', '24', '28', '30'].includes(day)) {
                // Обновляем заголовок для выбранного дня
                let dayName = 'день';
                
                if (day === '16') dayName = 'Четверг, 16 мая';
                else if (day === '24') dayName = 'Пятница, 24 мая';
                else if (day === '28') dayName = 'Вторник, 28 мая';
                else if (day === '30') dayName = 'Четверг, 30 мая';
                
                document.querySelector('.selected-day-title').textContent = dayName;
                
                // Показываем занятия (для демо всегда показываем одни и те же)
                document.querySelector('.selected-day .schedule-list').style.display = 'flex';
                
                if (day === '30') {
                    showNotification('На этот день есть только одно занятие: Физика');
                }
            } else {
                // Если нет занятий на этот день
                document.querySelector('.selected-day-title').textContent = `${day} мая - занятия отсутствуют`;
                document.querySelector('.selected-day .schedule-list').style.display = 'none';
            }
        }, 400);
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
}); 