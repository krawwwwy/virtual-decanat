// Функции для работы с посещаемостью
document.addEventListener('DOMContentLoaded', function() {
    // Данные студентов для разных групп
    const studentsData = {
        'mt101': [
            { name: 'Иванов Иван Иванович', status: 'present', comment: '' },
            { name: 'Петров Петр Петрович', status: 'present', comment: '' },
            { name: 'Сидорова Мария Александровна', status: 'absent', comment: 'Сообщила о болезни' },
            { name: 'Кузнецов Алексей Дмитриевич', status: 'excused', comment: 'Справка из медпункта' },
            { name: 'Новиков Дмитрий Сергеевич', status: 'present', comment: '' },
            { name: 'Морозова Елена Игоревна', status: 'absent', comment: '' },
            { name: 'Волков Артем Витальевич', status: 'present', comment: '' },
            { name: 'Соколова Анна Михайловна', status: 'present', comment: '' }
        ],
        'mt102': [
            { name: 'Смирнов Кирилл Андреевич', status: 'present', comment: '' },
            { name: 'Козлова Юлия Павловна', status: 'absent', comment: 'Отсутствует по семейным обстоятельствам' },
            { name: 'Никитин Максим Викторович', status: 'present', comment: '' },
            { name: 'Васильева Ольга Сергеевна', status: 'present', comment: '' },
            { name: 'Павлов Александр Дмитриевич', status: 'excused', comment: 'Справка' },
            { name: 'Семенова Валерия Алексеевна', status: 'present', comment: '' }
        ],
        'mt201': [
            { name: 'Орлова Светлана Игоревна', status: 'present', comment: '' },
            { name: 'Федоров Николай Владимирович', status: 'absent', comment: '' },
            { name: 'Михайлова Татьяна Степановна', status: 'present', comment: '' },
            { name: 'Зайцев Виктор Антонович', status: 'present', comment: '' },
            { name: 'Соловьева Екатерина Ильинична', status: 'excused', comment: 'Участие в олимпиаде' },
            { name: 'Лебедева Дарья Михайловна', status: 'absent', comment: 'Болеет' },
            { name: 'Титов Роман Никитич', status: 'present', comment: '' }
        ]
    };
    
    // Текущая выбранная группа
    let currentGroup = 'mt101';
    
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

    // Создание HTML для строки студента
    function createStudentRow(student, index) {
        // Определяем статус для установки правильного выбора радио-кнопок
        const isPresentChecked = student.status === 'present' ? 'checked' : '';
        const isAbsentChecked = student.status === 'absent' ? 'checked' : '';
        const isExcusedChecked = student.status === 'excused' ? 'checked' : '';
        
        // Определяем CSS-класс для строки в зависимости от статуса
        const rowStyle = student.status === 'absent' ? 'style="background-color: rgba(255, 200, 200, 0.1);"' : 
                        student.status === 'excused' ? 'style="background-color: rgba(200, 255, 200, 0.1);"' : '';
        
        return `
            <tr ${rowStyle}>
                <td>
                    <div class="student-name">
                        <div class="student-avatar"></div>
                        <span>${student.name}</span>
                    </div>
                </td>
                <td class="check-column">
                    <input type="radio" name="student_${index}" class="attendance-checkbox" ${isPresentChecked}>
                </td>
                <td class="check-column">
                    <input type="radio" name="student_${index}" class="attendance-checkbox" ${isAbsentChecked}>
                </td>
                <td class="check-column">
                    <input type="radio" name="student_${index}" class="attendance-checkbox" ${isExcusedChecked}>
                </td>
                <td>
                    <input type="text" placeholder="Комментарий" value="${student.comment}">
                </td>
            </tr>
        `;
    }

    // Обновление таблицы студентов для выбранной группы
    function renderStudentList(groupCode) {
        const tableBody = document.querySelector('.attendance-table tbody');
        if (!tableBody) return;
        
        if (groupCode === 'all') {
            groupCode = 'mt101'; // По умолчанию показываем первую группу при выборе "Все группы"
        }
        
        // Проверяем наличие данных для группы
        const students = studentsData[groupCode];
        if (!students || students.length === 0) {
            tableBody.innerHTML = '<tr><td colspan="5">Нет данных для выбранной группы</td></tr>';
            return;
        }
        
        // Формируем строки таблицы с данными студентов
        let tableHTML = '';
        students.forEach((student, index) => {
            tableHTML += createStudentRow(student, index + 1);
        });
        
        // Обновляем содержимое таблицы
        tableBody.innerHTML = tableHTML;
        
        // Повторно настраиваем обработчики для новых элементов
        setupTableInteractions();
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
                    
                    // Сбросить заголовок на текущий при возврате на вкладку "Сегодня"
                    document.querySelector('.attendance-date').textContent = 'Прикладная математика (кабинет 102)';
                } else if (tabType === 'stats') {
                    document.querySelector('.attendance-stats').style.display = 'flex';
                    document.querySelector('.attendance-table').style.display = 'none';
                    document.querySelector('.attendance-actions').style.display = 'none';
                    document.querySelector('.attendance-filter').style.display = 'none';
                    document.querySelector('.attendance-subject').style.display = 'none';
                }
            });
        });
        
        // Удаляем вкладку "История", если она есть
        const historyTab = document.querySelector('.tab[data-tab="history"]');
        if (historyTab) {
            historyTab.style.display = 'none';
        }
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
                
                // Сохраняем текущую выбранную группу
                currentGroup = selectedGroup === 'all' ? 'mt101' : selectedGroup;
                
                // Эмуляция загрузки данных выбранной группы
                showLoading(true);
                
                setTimeout(() => {
                    // Обновляем информацию о предмете
                    const groupInfoText = document.querySelector('.attendance-subject > div:last-child > div:nth-child(2)');
                    if (groupInfoText) {
                        groupInfoText.textContent = `10:00 - 11:00 | Группа: ${selectedGroup === 'all' ? 'МТ-101' : selectedGroup.toUpperCase()}`;
                    }
                    
                    // Обновляем заголовок предмета в зависимости от группы
                    const subjectTitle = document.querySelector('.attendance-date');
                    if (subjectTitle) {
                        if (selectedGroup === 'all' || selectedGroup === 'mt101') {
                            subjectTitle.textContent = 'Прикладная математика (кабинет 102)';
                        } else if (selectedGroup === 'mt102') {
                            subjectTitle.textContent = 'Начертательная геометрия (кабинет 103)';
                        } else if (selectedGroup === 'mt201') {
                            subjectTitle.textContent = 'Линейная алгебра (кабинет 104)';
                        }
                    }
                    
                    // Обновляем список студентов
                    renderStudentList(selectedGroup);
                    
                    showLoading(false);
                }, 500);
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
                        // Проверяем первую радио-кнопку (Присутствует) в строке
                        const cell = row.querySelector('td:nth-child(2)');
                        const isPresentChecked = cell && cell.querySelector('input[type="radio"]').checked;
                        row.style.display = isPresentChecked ? '' : 'none';
                    });
                } else if (filterType === 'Отсутствует') {
                    rows.forEach(row => {
                        // Проверяем вторую радио-кнопку (Отсутствует) в строке
                        const cell = row.querySelector('td:nth-child(3)');
                        const isAbsentChecked = cell && cell.querySelector('input[type="radio"]').checked;
                        row.style.display = isAbsentChecked ? '' : 'none';
                    });
                } else if (filterType === 'По уважительной') {
                    rows.forEach(row => {
                        // Проверяем третью радио-кнопку (По уважительной) в строке
                        const cell = row.querySelector('td:nth-child(4)');
                        const isExcusedChecked = cell && cell.querySelector('input[type="radio"]').checked;
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
                        const isPresentChecked = row.querySelector('td:nth-child(2) input[type="radio"]').checked;
                        const isAbsentChecked = row.querySelector('td:nth-child(3) input[type="radio"]').checked;
                        const isExcusedChecked = row.querySelector('td:nth-child(4) input[type="radio"]').checked;
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
                    // Восстанавливаем исходное состояние для текущей группы
                    renderStudentList(currentGroup);
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
                
                if (this.parentNode.classList.contains('check-column')) {
                    const td = this.parentNode;
                    // Определяем столбец по индексу ячейки
                    const cellIndex = Array.from(row.cells).indexOf(td);
                    
                    if (cellIndex === 1) { // Присутствует - первая ячейка с радио
                        row.style.backgroundColor = '';
                    } else if (cellIndex === 2) { // Отсутствует - вторая ячейка с радио
                        row.style.backgroundColor = 'rgba(255, 200, 200, 0.1)';
                    } else if (cellIndex === 3) { // По уважительной - третья ячейка с радио
                        row.style.backgroundColor = 'rgba(200, 255, 200, 0.1)';
                    }
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
    
    // Инициализация с данными первой группы по умолчанию
    renderStudentList('mt101');
}); 