// Функции для работы с посещаемостью
document.addEventListener('DOMContentLoaded', function() {
    // Проверка авторизации
    function checkAuth() {
        const token = localStorage.getItem('accessToken');
        if (!token) {
            window.location.href = '../login-page.html';
            return false;
        }
        return true;
    }
    
    // Загрузка профиля преподавателя
    function loadTeacherProfile() {
        try {
            if (!checkAuth()) return;
            
            // Получаем имя пользователя из localStorage
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
    
    // Настройка переключения вкладок
    function setupTabs() {
        const tabs = document.querySelectorAll('.tab');
        tabs.forEach(tab => {
            tab.addEventListener('click', function() {
                tabs.forEach(t => t.classList.remove('active'));
                this.classList.add('active');
                
                const tabType = this.getAttribute('data-tab');
                
                // Показать/скрыть соответствующие секции в зависимости от вкладки
                if (tabType === 'today') {
                    document.querySelector('.attendance-stats').style.display = 'none';
                    document.querySelector('.attendance-table').style.display = 'table';
                    document.querySelector('.attendance-actions').style.display = 'flex';
                    document.querySelector('.attendance-filter').style.display = 'flex';
                    document.querySelector('.attendance-subject').style.display = 'flex';
                } else if (tabType === 'history') {
                    document.querySelector('.attendance-stats').style.display = 'none';
                    document.querySelector('.attendance-table').style.display = 'table';
                    document.querySelector('.attendance-actions').style.display = 'flex';
                    document.querySelector('.attendance-filter').style.display = 'flex';
                    document.querySelector('.attendance-subject').style.display = 'flex';
                    // В реальной системе здесь бы загружались исторические данные
                    loadHistoricalData();
                } else if (tabType === 'stats') {
                    document.querySelector('.attendance-stats').style.display = 'flex';
                    document.querySelector('.attendance-table').style.display = 'none';
                    document.querySelector('.attendance-actions').style.display = 'none';
                    document.querySelector('.attendance-filter').style.display = 'none';
                    document.querySelector('.attendance-subject').style.display = 'none';
                }
            });
        });
    }
    
    // Имитация загрузки исторических данных
    function loadHistoricalData() {
        // Эмуляция загрузки с задержкой для имитации сетевого запроса
        showLoading(true);
        
        setTimeout(() => {
            // Обновляем заголовок с информацией о предмете
            const subjectInfo = document.querySelector('.attendance-date');
            subjectInfo.textContent = 'Прикладная математика (занятие от 19.05.2025)';
            
            // Здесь могла бы быть реальная загрузка данных по API
            showLoading(false);
        }, 700);
    }
    
    // Показать/скрыть индикатор загрузки
    function showLoading(isLoading) {
        // Простая реализация без UI
        if (isLoading) {
            console.log('Загрузка данных...');
            // В реальном проекте здесь было бы отображение UI-элемента загрузки
        } else {
            console.log('Загрузка завершена');
        }
    }
    
    // Настройка фильтров
    function setupFilters() {
        // Обработчик выбора группы
        const groupFilter = document.getElementById('groupFilter');
        if (groupFilter) {
            groupFilter.addEventListener('change', function() {
                const selectedGroup = this.value;
                console.log('Выбрана группа:', selectedGroup);
                
                // Эмуляция загрузки данных выбранной группы
                if (selectedGroup !== 'all') {
                    showLoading(true);
                    
                    setTimeout(() => {
                        // Обновляем информацию о предмете
                        const subjectInfo = document.querySelector('.attendance-subject div:last-child div:last-child');
                        subjectInfo.textContent = `10:00 - 11:00 | Группа: ${selectedGroup.toUpperCase()}`;
                        
                        // Можно обновить весь список студентов для выбранной группы
                        showLoading(false);
                    }, 500);
                }
            });
        }
        
        // Обработчики для фильтров статуса
        const filters = document.querySelectorAll('.filter-item');
        filters.forEach(filter => {
            filter.addEventListener('click', function() {
                filters.forEach(f => f.classList.remove('active'));
                this.classList.add('active');
                
                const filterType = this.textContent.trim();
                console.log('Применен фильтр:', filterType);
                
                // Фильтрация студентов по статусу посещаемости
                const rows = document.querySelectorAll('.attendance-table tbody tr');
                
                if (filterType === 'Все') {
                    rows.forEach(row => row.style.display = '');
                } else if (filterType === 'Присутствует') {
                    rows.forEach(row => {
                        const isPresentChecked = row.querySelector('input[type="radio"]:nth-of-type(1)').checked;
                        row.style.display = isPresentChecked ? '' : 'none';
                    });
                } else if (filterType === 'Отсутствует') {
                    rows.forEach(row => {
                        const isAbsentChecked = row.querySelector('input[type="radio"]:nth-of-type(2)').checked;
                        row.style.display = isAbsentChecked ? '' : 'none';
                    });
                } else if (filterType === 'По уважительной') {
                    rows.forEach(row => {
                        const isExcusedChecked = row.querySelector('input[type="radio"]:nth-of-type(3)').checked;
                        row.style.display = isExcusedChecked ? '' : 'none';
                    });
                }
            });
        });
    }
    
    // Обработка кнопок сохранения
    function setupActionButtons() {
        const saveButton = document.querySelector('.save-button');
        if (saveButton) {
            saveButton.addEventListener('click', function() {
                showLoading(true);
                
                // Имитируем сохранение данных
                setTimeout(() => {
                    showLoading(false);
                    
                    // Сбор данных о посещаемости
                    const attendanceData = [];
                    const rows = document.querySelectorAll('.attendance-table tbody tr');
                    
                    rows.forEach((row, index) => {
                        const studentName = row.querySelector('.student-name span').textContent;
                        const isPresentChecked = row.querySelector('input[type="radio"]:nth-of-type(1)').checked;
                        const isAbsentChecked = row.querySelector('input[type="radio"]:nth-of-type(2)').checked;
                        const isExcusedChecked = row.querySelector('input[type="radio"]:nth-of-type(3)').checked;
                        const comment = row.querySelector('input[type="text"]').value;
                        
                        let status = 'unknown';
                        if (isPresentChecked) status = 'present';
                        if (isAbsentChecked) status = 'absent';
                        if (isExcusedChecked) status = 'excused';
                        
                        attendanceData.push({
                            studentName,
                            status,
                            comment
                        });
                    });
                    
                    console.log('Данные посещаемости сохранены:', attendanceData);
                    alert('Данные о посещаемости успешно сохранены!');
                    
                    // В реальном проекте здесь был бы API-запрос для сохранения данных
                }, 800);
            });
        }
        
        const cancelButton = document.querySelector('.cancel-button');
        if (cancelButton) {
            cancelButton.addEventListener('click', function() {
                if (confirm('Вы уверены, что хотите отменить все изменения?')) {
                    window.location.reload();
                }
            });
        }
    }
    
    // Обработка изменений в таблице
    function setupTableInteractions() {
        // Проверяем изменение радио-кнопок
        const radioButtons = document.querySelectorAll('input[type="radio"]');
        radioButtons.forEach(radio => {
            radio.addEventListener('change', function() {
                // Можно добавить визуальную обратную связь
                const row = this.closest('tr');
                
                if (this.parentNode.classList.contains('check-column') && 
                    this.parentNode === row.querySelector('td:nth-child(2)')) {
                    // Если отмечено "Присутствует"
                    row.style.backgroundColor = '';
                } else if (this.parentNode.classList.contains('check-column') && 
                          this.parentNode === row.querySelector('td:nth-child(3)')) {
                    // Если отмечено "Отсутствует"
                    row.style.backgroundColor = 'rgba(255, 200, 200, 0.1)';
                } else if (this.parentNode.classList.contains('check-column') && 
                          this.parentNode === row.querySelector('td:nth-child(4)')) {
                    // Если отмечено "По уважительной"
                    row.style.backgroundColor = 'rgba(200, 255, 200, 0.1)';
                }
            });
        });
    }
    
    // Инициализация всех функций
    loadTeacherProfile();
    setupTabs();
    setupFilters();
    setupActionButtons();
    setupTableInteractions();
}); 