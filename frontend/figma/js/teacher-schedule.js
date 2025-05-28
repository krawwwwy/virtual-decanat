// Функции для работы с расписанием преподавателя
document.addEventListener('DOMContentLoaded', function() {
    // Функция для получения данных профиля преподавателя из API
    async function loadTeacherProfile() {
        try {
            const token = localStorage.getItem('accessToken');
            if (!token) {
                window.location.href = '../login-page.html';
                return;
            }
            
            // Имитация загрузки данных преподавателя
            // В реальном приложении здесь будет запрос к API
            const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}');
            document.getElementById('userName').textContent = userInfo.name || 'Преподаватель';
            
            // Обработчик для кнопки выхода
            document.getElementById('logoutBtn').addEventListener('click', function() {
                localStorage.removeItem('accessToken');
                localStorage.removeItem('refreshToken');
                localStorage.removeItem('tokenExpires');
                localStorage.removeItem('userInfo');
                localStorage.removeItem('userRole');
                window.location.href = '../login-page.html';
            });
        } catch (error) {
            console.error('Ошибка при загрузке профиля:', error);
        }
    }
    
    // Функция для переключения вкладок расписания (день/неделя/календарь)
    function setupTabSwitcher() {
        const tabs = document.querySelectorAll('.tab');
        const views = document.querySelectorAll('.schedule-view');
        
        tabs.forEach(tab => {
            tab.addEventListener('click', function() {
                // Удаляем активный класс у всех вкладок
                tabs.forEach(t => t.classList.remove('active'));
                // Добавляем активный класс текущей вкладке
                this.classList.add('active');
                
                // Скрываем все представления
                views.forEach(view => view.classList.remove('active-view'));
                
                // Показываем выбранное представление
                const viewName = this.getAttribute('data-view');
                document.querySelector(`.${viewName}-view`).classList.add('active-view');
            });
        });
    }
    
    // Функция для обработки выбора дня недели
    function setupWeekDaySelector() {
        const weekDays = document.querySelectorAll('.week-day');
        
        weekDays.forEach(day => {
            day.addEventListener('click', function() {
                // Удаляем активный класс у всех дней
                weekDays.forEach(d => {
                    d.classList.remove('active');
                    d.querySelector('.day-indicator').classList.remove('active');
                });
                
                // Добавляем активный класс выбранному дню
                this.classList.add('active');
                this.querySelector('.day-indicator').classList.add('active');
                
                // Здесь должна быть логика загрузки расписания для выбранного дня
                const dateSelected = this.getAttribute('data-date');
                console.log(`Выбран день: ${dateSelected}`);
                
                // Показываем расписание для выбранного дня (в данном случае просто имитация)
                document.querySelectorAll('.schedule-list[data-date]').forEach(list => {
                    list.style.display = 'none';
                });
                
                const scheduleForDay = document.querySelector(`.schedule-list[data-date="${dateSelected}"]`);
                if (scheduleForDay) {
                    scheduleForDay.style.display = 'block';
                }
            });
        });
    }
    
    // Функция для кнопки "Сегодня"
    function setupTodayButton() {
        const todayButton = document.querySelector('.today-button');
        if (todayButton) {
            todayButton.addEventListener('click', function() {
                // Логика выбора сегодняшней даты
                // В реальном приложении здесь будет установка текущей даты
                alert('Показать расписание на сегодня');
            });
        }
        
        const thisWeekButton = document.querySelector('.this-week-button');
        if (thisWeekButton) {
            thisWeekButton.addEventListener('click', function() {
                // Логика выбора текущей недели
                alert('Показать расписание на текущую неделю');
            });
        }
    }
    
    // Функция для деталей занятия
    function setupLessonDetails() {
        const detailsButtons = document.querySelectorAll('.details-button');
        
        detailsButtons.forEach(button => {
            button.addEventListener('click', function() {
                const scheduleItem = this.closest('.schedule-item');
                const subjectTitle = scheduleItem.querySelector('.subject-title').textContent;
                const subjectTime = scheduleItem.querySelector('.subject-time').textContent;
                const subjectGroup = scheduleItem.querySelector('.subject-group').textContent;
                
                alert(`Детали занятия:\n${subjectTitle}\n${subjectTime}\n${subjectGroup}`);
            });
        });
    }
    
    // Инициализация всех функций
    loadTeacherProfile();
    setupTabSwitcher();
    setupWeekDaySelector();
    setupTodayButton();
    setupLessonDetails();
}); 