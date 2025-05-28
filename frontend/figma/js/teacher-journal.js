// Функции для работы с журналом успеваемости
document.addEventListener('DOMContentLoaded', function() {
    // Данные студентов для разных групп и предметов
    const studentsData = {
        math: {
            mt101: [
                { name: 'Иванов Иван Иванович', grades: [5, 4, 5, 4, 5, 4, 5], exam: 5, final: 5 },
                { name: 'Петров Петр Петрович', grades: [4, 4, 3, 4, 5, 4, 4], exam: 4, final: 4 },
                { name: 'Сидорова Мария Александровна', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Кузнецов Алексей Дмитриевич', grades: [3, 3, 2, 4, 3, 3, 4], exam: 3, final: 3 },
                { name: 'Новиков Дмитрий Сергеевич', grades: [4, 4, 4, 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Морозова Елена Игоревна', grades: [5, 4, 5, 5, 5, 4, 5], exam: 5, final: 5 },
                { name: 'Волков Артем Витальевич', grades: [4, 3, 4, 3, 4, 3, 4], exam: 4, final: 4 },
                { name: 'Соколова Анна Михайловна', grades: [5, 5, 4, 5, 5, 5, 5], exam: 5, final: 5 }
            ],
            mt102: [
                { name: 'Смирнов Кирилл Андреевич', grades: [4, 4, 4, 3, 4, 3, 4], exam: 4, final: 4 },
                { name: 'Козлова Юлия Павловна', grades: [5, 4, 5, 5, 4, 5, 4], exam: 4, final: 4 },
                { name: 'Никитин Максим Викторович', grades: [3, 3, 3, 3, 2, 3, 3], exam: 3, final: 3 },
                { name: 'Васильева Ольга Сергеевна', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Павлов Александр Дмитриевич', grades: [4, 'н', 4, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Семенова Валерия Алексеевна', grades: [4, 3, 4, 4, 3, 4, 4], exam: 4, final: 4 }
            ],
            mt201: [
                { name: 'Орлова Светлана Игоревна', grades: [5, 5, 5, 5, 5, 5, 4], exam: 5, final: 5 },
                { name: 'Федоров Николай Владимирович', grades: [3, 'н', 3, 'н', 3, 3, 3], exam: 3, final: 3 },
                { name: 'Михайлова Татьяна Степановна', grades: [4, 4, 4, 5, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Зайцев Виктор Антонович', grades: [4, 4, 4, 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Соловьева Екатерина Ильинична', grades: [5, 5, 5, 5, 'н', 5, 5], exam: 5, final: 5 },
                { name: 'Лебедева Дарья Михайловна', grades: [4, 3, 3, 4, 4, 4, 3], exam: 4, final: 4 },
                { name: 'Титов Роман Никитич', grades: [3, 3, 2, 3, 2, 3, 3], exam: 3, final: 3 }
            ]
        },
        geometry: {
            mt101: [
                { name: 'Иванов Иван Иванович', grades: [4, 3, 4, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Петров Петр Петрович', grades: [3, 3, 2, 3, 3, 3, 3], exam: 3, final: 3 },
                { name: 'Сидорова Мария Александровна', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Кузнецов Алексей Дмитриевич', grades: [4, 3, 4, 3, 3, 4, 4], exam: 4, final: 4 },
                { name: 'Новиков Дмитрий Сергеевич', grades: [3, 4, 3, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Морозова Елена Игоревна', grades: [5, 4, 5, 5, 4, 5, 5], exam: 5, final: 5 },
                { name: 'Волков Артем Витальевич', grades: [3, 3, 3, 3, 3, 3, 3], exam: 3, final: 3 },
                { name: 'Соколова Анна Михайловна', grades: [4, 5, 4, 4, 5, 5, 5], exam: 5, final: 5 }
            ],
            mt102: [
                { name: 'Смирнов Кирилл Андреевич', grades: [5, 4, 4, 5, 4, 5, 4], exam: 5, final: 5 },
                { name: 'Козлова Юлия Павловна', grades: [4, 'н', 'н', 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Никитин Максим Викторович', grades: [4, 4, 4, 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Васильева Ольга Сергеевна', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Павлов Александр Дмитриевич', grades: [3, 3, 3, 3, 4, 3, 3], exam: 3, final: 3 },
                { name: 'Семенова Валерия Алексеевна', grades: [4, 4, 3, 4, 4, 4, 4], exam: 4, final: 4 }
            ],
            mt201: []
        },
        algebra: {
            mt101: [
                { name: 'Иванов Иван Иванович', grades: [3, 3, 4, 3, 3, 3, 3], exam: 3, final: 3 },
                { name: 'Петров Петр Петрович', grades: [4, 4, 4, 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Сидорова Мария Александровна', grades: [5, 5, 5, 4, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Кузнецов Алексей Дмитриевич', grades: [3, 2, 3, 3, 2, 3, 3], exam: 3, final: 3 },
                { name: 'Новиков Дмитрий Сергеевич', grades: [4, 4, 3, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Морозова Елена Игоревна', grades: [5, 5, 4, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Волков Артем Витальевич', grades: [3, 3, 3, 3, 3, 3, 3], exam: 3, final: 3 },
                { name: 'Соколова Анна Михайловна', grades: [4, 4, 4, 5, 4, 4, 4], exam: 4, final: 4 }
            ],
            mt102: [],
            mt201: [
                { name: 'Орлова Светлана Игоревна', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Федоров Николай Владимирович', grades: [3, 3, 'н', 3, 3, 'н', 3], exam: 3, final: 3 },
                { name: 'Михайлова Татьяна Степановна', grades: [4, 4, 4, 4, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Зайцев Виктор Антонович', grades: [3, 4, 3, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Соловьева Екатерина Ильинична', grades: [5, 5, 5, 5, 5, 5, 5], exam: 5, final: 5 },
                { name: 'Лебедева Дарья Михайловна', grades: [4, 3, 3, 3, 4, 4, 4], exam: 4, final: 4 },
                { name: 'Титов Роман Никитич', grades: [2, 3, 2, 3, 3, 2, 3], exam: 3, final: 3 }
            ]
        }
    };
    
    // Текущий выбранный предмет и группа
    let currentSubject = 'math';
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
        let gradesHTML = '';
        let gradeSum = 0;
        let gradeCount = 0;
        
        // Генерируем HTML для ячеек с оценками
        student.grades.forEach(grade => {
            const isAbsent = grade === 'н';
            const gradeClass = isAbsent ? 'grade-absent' : '';
            
            gradesHTML += `
                <td>
                    <input type="text" class="grade-input ${gradeClass}" value="${grade}" maxlength="1">
                </td>
            `;
            
            // Подсчет среднего балла (исключая пропуски)
            if (!isAbsent) {
                gradeSum += parseInt(grade);
                gradeCount++;
            }
        });
        
        // Вычисляем средний балл
        const avgGrade = gradeCount > 0 ? (gradeSum / gradeCount).toFixed(1) : '-';
        
        return `
            <tr data-student-id="${index}">
                <td class="student-name-column">
                    <div class="student-name">
                        <div class="student-avatar"></div>
                        <span>${student.name}</span>
                    </div>
                </td>
                ${gradesHTML}
                <td class="avg-grade">${avgGrade}</td>
                <td>
                    <input type="text" class="grade-input" value="${student.exam}" maxlength="1">
                </td>
                <td>
                    <input type="text" class="grade-input" value="${student.final}" maxlength="1">
                </td>
            </tr>
        `;
    }

    // Обновление таблицы студентов для выбранных предмета и группы
    function renderStudentList() {
        const tableBody = document.querySelector('.journal-table tbody');
        if (!tableBody) return;
        
        // Проверяем наличие данных для группы и предмета
        const students = studentsData[currentSubject]?.[currentGroup];
        if (!students || students.length === 0) {
            tableBody.innerHTML = '<tr><td colspan="11" class="no-data">Нет данных для выбранной группы и предмета</td></tr>';
            return;
        }
        
        // Формируем строки таблицы с данными студентов
        let tableHTML = '';
        students.forEach((student, index) => {
            tableHTML += createStudentRow(student, index);
        });
        
        // Обновляем содержимое таблицы
        tableBody.innerHTML = tableHTML;
        
        // Обновляем информацию о предмете и группе
        updateJournalInfo();
        
        // Повторно настраиваем обработчики для новых элементов
        setupGradeInputs();
    }
    
    // Обновление информации о предмете и группе в заголовке
    function updateJournalInfo() {
        const journalTitle = document.querySelector('.journal-info .journal-title');
        const journalGroupInfo = document.querySelector('.journal-info > div:last-child > div');
        const subjectIcon = document.querySelector('.journal-info .subject-icon img');
        
        if (journalTitle && journalGroupInfo) {
            // Устанавливаем название предмета
            let subjectName = 'Прикладная математика';
            if (currentSubject === 'geometry') {
                subjectName = 'Начертательная геометрия';
            } else if (currentSubject === 'algebra') {
                subjectName = 'Линейная алгебра';
            }
            journalTitle.textContent = subjectName;
            
            // Устанавливаем группу
            journalGroupInfo.textContent = `Группа: ${currentGroup.toUpperCase()} | Семестр: Весенний 2025`;
            
            // Устанавливаем иконку предмета
            if (subjectIcon) {
                if (currentSubject === 'math' || currentSubject === 'algebra') {
                    subjectIcon.src = '../../assets/images/teacher/math_icon.svg';
                } else if (currentSubject === 'geometry') {
                    subjectIcon.src = '../../assets/images/teacher/geometry_icon.svg';
                }
            }
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
                
                // В реальной системе здесь бы переключались виды представления
                console.log('Выбрана вкладка:', tabType);
                
                // Демонстрационная логика отображения разных вкладок
                if (tabType === 'grades') {
                    document.querySelector('.journal-table').style.display = 'table';
                    document.querySelector('.journal-actions').style.display = 'flex';
                    showCommentsView(false);
                    showAnalyticsView(false);
                } else if (tabType === 'comments') {
                    document.querySelector('.journal-table').style.display = 'none';
                    document.querySelector('.journal-actions').style.display = 'flex';
                    showCommentsView(true);
                    showAnalyticsView(false);
                } else if (tabType === 'analytics') {
                    document.querySelector('.journal-table').style.display = 'none';
                    document.querySelector('.journal-actions').style.display = 'none';
                    showCommentsView(false);
                    showAnalyticsView(true);
                }
            });
        });
    }
    
    // Показать представление с комментариями
    function showCommentsView(show) {
        // Проверяем, существует ли элемент
        let commentsView = document.querySelector('.comments-view');
        
        if (show) {
            // Если элемента нет, создаем его
            if (!commentsView) {
                commentsView = document.createElement('div');
                commentsView.className = 'comments-view';
                commentsView.innerHTML = `
                    <h3>Комментарии к успеваемости</h3>
                    <div class="comment-list">
                        <div class="comment-item">
                            <div class="comment-header">
                                <div class="student-name">
                                    <div class="student-avatar"></div>
                                    <span>Иванов Иван Иванович</span>
                                </div>
                                <div class="comment-date">Добавлено: 20.05.2025</div>
                            </div>
                            <div class="comment-text">
                                <textarea rows="3" class="comment-input">Хорошо справляется с задачами, но нужно обратить внимание на выполнение домашних заданий.</textarea>
                            </div>
                        </div>
                        <div class="comment-item">
                            <div class="comment-header">
                                <div class="student-name">
                                    <div class="student-avatar"></div>
                                    <span>Петров Петр Петрович</span>
                                </div>
                                <div class="comment-date">Добавлено: 22.05.2025</div>
                            </div>
                            <div class="comment-text">
                                <textarea rows="3" class="comment-input">Требуется дополнительная работа над темой "Интегралы". Рекомендованы консультации.</textarea>
                            </div>
                        </div>
                        <button class="action-button save-button" style="margin-top: 20px;">Сохранить комментарии</button>
                    </div>
                    <style>
                        .comment-item {
                            background-color: #fff;
                            border-radius: 8px;
                            padding: 16px;
                            margin-bottom: 16px;
                            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
                        }
                        .comment-header {
                            display: flex;
                            justify-content: space-between;
                            margin-bottom: 12px;
                        }
                        .comment-date {
                            color: var(--text-secondary);
                            font-size: 14px;
                        }
                        .comment-text {
                            margin-top: 8px;
                        }
                        .comment-input {
                            width: 100%;
                            padding: 8px;
                            border: 1px solid var(--border-color);
                            border-radius: 4px;
                            resize: vertical;
                        }
                    </style>
                `;
                
                // Добавляем в DOM после заголовка (перед таблицей)
                const contentDiv = document.querySelector('.journal-content');
                const journalInfo = document.querySelector('.journal-info');
                contentDiv.insertBefore(commentsView, journalInfo.nextSibling);
                
                // Добавляем обработчик событий для кнопки сохранения
                const saveButton = commentsView.querySelector('.save-button');
                saveButton.addEventListener('click', function() {
                    showLoading(true);
                    setTimeout(() => {
                        showLoading(false);
                        alert('Комментарии успешно сохранены!');
                    }, 700);
                });
            }
            commentsView.style.display = 'block';
        } else if (commentsView) {
            commentsView.style.display = 'none';
        }
    }
    
    // Показать представление с аналитикой
    function showAnalyticsView(show) {
        // Проверяем, существует ли элемент
        let analyticsView = document.querySelector('.analytics-view');
        
        if (show) {
            // Если элемента нет, создаем его
            if (!analyticsView) {
                analyticsView = document.createElement('div');
                analyticsView.className = 'analytics-view';
                analyticsView.innerHTML = `
                    <h3>Аналитика успеваемости</h3>
                    <div class="analytics-cards">
                        <div class="analytics-card">
                            <div class="analytics-value">4.1</div>
                            <div class="analytics-label">Средний балл по группе</div>
                            <div class="analytics-trend positive">+0.2 к прошлому семестру</div>
                        </div>
                        <div class="analytics-card">
                            <div class="analytics-value">87%</div>
                            <div class="analytics-label">Успеваемость</div>
                            <div class="analytics-trend positive">+5% к прошлому семестру</div>
                        </div>
                        <div class="analytics-card">
                            <div class="analytics-value">25%</div>
                            <div class="analytics-label">Отличники</div>
                            <div class="analytics-trend positive">+10% к прошлому семестру</div>
                        </div>
                    </div>
                    <div class="analytics-distribution">
                        <h4>Распределение оценок</h4>
                        <div class="grade-bars">
                            <div class="grade-bar">
                                <div class="grade-label">5</div>
                                <div class="grade-progress">
                                    <div class="grade-progress-fill" style="width: 25%;"></div>
                                </div>
                                <div class="grade-percent">25%</div>
                            </div>
                            <div class="grade-bar">
                                <div class="grade-label">4</div>
                                <div class="grade-progress">
                                    <div class="grade-progress-fill" style="width: 40%;"></div>
                                </div>
                                <div class="grade-percent">40%</div>
                            </div>
                            <div class="grade-bar">
                                <div class="grade-label">3</div>
                                <div class="grade-progress">
                                    <div class="grade-progress-fill" style="width: 30%;"></div>
                                </div>
                                <div class="grade-percent">30%</div>
                            </div>
                            <div class="grade-bar">
                                <div class="grade-label">2</div>
                                <div class="grade-progress">
                                    <div class="grade-progress-fill" style="width: 5%;"></div>
                                </div>
                                <div class="grade-percent">5%</div>
                            </div>
                        </div>
                    </div>
                    <style>
                        .analytics-cards {
                            display: flex;
                            gap: 16px;
                            margin-bottom: 24px;
                        }
                        .analytics-card {
                            flex: 1;
                            background-color: #fff;
                            border-radius: 8px;
                            padding: 16px;
                            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
                        }
                        .analytics-value {
                            font-size: 28px;
                            font-weight: 700;
                            margin-bottom: 8px;
                        }
                        .analytics-label {
                            font-size: 14px;
                            color: var(--text-secondary);
                            margin-bottom: 8px;
                        }
                        .analytics-trend {
                            font-size: 12px;
                            font-weight: 500;
                        }
                        .analytics-trend.positive {
                            color: green;
                        }
                        .analytics-trend.negative {
                            color: red;
                        }
                        .analytics-distribution {
                            background-color: #fff;
                            border-radius: 8px;
                            padding: 16px;
                            box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
                        }
                        .grade-bars {
                            display: flex;
                            flex-direction: column;
                            gap: 12px;
                        }
                        .grade-bar {
                            display: flex;
                            align-items: center;
                            gap: 12px;
                        }
                        .grade-label {
                            width: 24px;
                            text-align: center;
                            font-weight: 600;
                        }
                        .grade-progress {
                            flex: 1;
                            height: 20px;
                            background-color: #f0f0f0;
                            border-radius: 10px;
                            overflow: hidden;
                        }
                        .grade-progress-fill {
                            height: 100%;
                            background-color: var(--primary-color);
                        }
                        .grade-percent {
                            width: 40px;
                            text-align: right;
                        }
                    </style>
                `;
                
                // Добавляем в DOM после заголовка (перед таблицей)
                const contentDiv = document.querySelector('.journal-content');
                const journalInfo = document.querySelector('.journal-info');
                contentDiv.insertBefore(analyticsView, journalInfo.nextSibling);
            }
            analyticsView.style.display = 'block';
        } else if (analyticsView) {
            analyticsView.style.display = 'none';
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
        // Обработчики выбора предмета и группы
        const subjectFilter = document.getElementById('subjectFilter');
        const groupFilter = document.getElementById('groupFilter');
        
        if (subjectFilter) {
            subjectFilter.addEventListener('change', function() {
                currentSubject = this.value;
                console.log('Выбран предмет:', currentSubject);
                
                showLoading(true);
                setTimeout(() => {
                    renderStudentList();
                    showLoading(false);
                }, 300);
            });
        }
        
        if (groupFilter) {
            groupFilter.addEventListener('change', function() {
                currentGroup = this.value;
                console.log('Выбрана группа:', currentGroup);
                
                showLoading(true);
                setTimeout(() => {
                    renderStudentList();
                    showLoading(false);
                }, 300);
            });
        }
    }
    
    // Настройка ячеек с оценками
    function setupGradeInputs() {
        const gradeInputs = document.querySelectorAll('.grade-input');
        if (!gradeInputs.length) return;
        
        gradeInputs.forEach(input => {
            // При фокусе сохраняем предыдущее значение
            input.addEventListener('focus', function() {
                this.setAttribute('data-previous-value', this.value);
            });
            
            // При изменении проводим валидацию
            input.addEventListener('input', function() {
                const value = this.value.trim();
                // Проверяем, что значение - это число от 2 до 5 или "н" (отсутствие)
                if (value && value !== 'н' && (isNaN(value) || parseInt(value) < 2 || parseInt(value) > 5)) {
                    const prevValue = this.getAttribute('data-previous-value') || '';
                    this.value = prevValue;
                } else {
                    // Если значение является отсутствием (н), добавляем класс
                    if (value === 'н') {
                        this.classList.add('grade-absent');
                    } else {
                        this.classList.remove('grade-absent');
                    }
                    
                    // Обновляем средний балл для строки
                    const row = this.closest('tr');
                    if (row) {
                        updateAverageGrade(row);
                    }
                }
            });
        });
    }
    
    // Обновление среднего балла для строки
    function updateAverageGrade(row) {
        // Собираем все оценки (кроме экзамена и итоговой)
        const gradeInputs = row.querySelectorAll('td:not(.avg-grade):not(:last-child):not(:last-child):not(:first-child) .grade-input');
        
        let sum = 0;
        let count = 0;
        
        gradeInputs.forEach(input => {
            const value = input.value.trim();
            // Учитываем только числовые оценки
            if (value !== 'н' && !isNaN(value) && value !== '') {
                sum += parseInt(value);
                count++;
            }
        });
        
        // Обновляем ячейку со средним баллом
        const avgCell = row.querySelector('.avg-grade');
        if (avgCell) {
            avgCell.textContent = count > 0 ? (sum / count).toFixed(1) : '-';
        }
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
                    
                    // Сбор данных из таблицы
                    const gradesData = [];
                    const rows = document.querySelectorAll('.journal-table tbody tr');
                    
                    rows.forEach(row => {
                        // Пропускаем строки с сообщением об отсутствии данных
                        if (row.classList.contains('no-data')) return;
                        
                        const studentId = row.getAttribute('data-student-id');
                        const studentName = row.querySelector('.student-name span').textContent;
                        
                        // Собираем оценки
                        const grades = [];
                        const gradeInputs = row.querySelectorAll('td:not(.avg-grade):not(:last-child):not(:last-child):not(:first-child) .grade-input');
                        gradeInputs.forEach(input => {
                            grades.push(input.value.trim());
                        });
                        
                        // Получаем экзаменационную и итоговую оценки
                        const examGrade = row.querySelector('td:nth-last-child(2) .grade-input').value;
                        const finalGrade = row.querySelector('td:last-child .grade-input').value;
                        
                        // Добавляем в массив
                        gradesData.push({
                            studentId,
                            studentName,
                            grades,
                            examGrade,
                            finalGrade
                        });
                    });
                    
                    console.log('Данные журнала сохранены:', gradesData);
                    alert('Данные успешно сохранены!');
                }, 800);
            });
        }
        
        const cancelButton = document.querySelector('.cancel-button');
        if (cancelButton) {
            cancelButton.addEventListener('click', function() {
                if (confirm('Вы уверены, что хотите отменить все изменения?')) {
                    // Восстановление данных с сервера (имитация)
                    showLoading(true);
                    setTimeout(() => {
                        // Восстанавливаем исходное состояние для текущей группы и предмета
                        renderStudentList();
                        showLoading(false);
                    }, 500);
                }
            });
        }
    }

    // Настройка функции экспорта в CSV
    function setupExport() {
        const exportButton = document.querySelector('.export-button');
        if (exportButton) {
            exportButton.addEventListener('click', function() {
                exportToCSV();
            });
        }
    }

    // Функция экспорта данных журнала в CSV-файл
    function exportToCSV() {
        // Получаем заголовки таблицы
        const headerRow = document.querySelector('.journal-table thead tr:first-child');
        const subHeaderRow = document.querySelector('.journal-table thead tr:last-child');
        
        if (!headerRow || !subHeaderRow) return;
        
        // Массив для заголовков
        const headers = ['Студент'];
        
        // Добавляем заголовки для оценок
        const lessonCells = subHeaderRow.querySelectorAll('th');
        lessonCells.forEach(cell => {
            if (cell.textContent.trim()) {  // Пропускаем пустые ячейки
                const date = cell.querySelector('.journal-date')?.textContent || '';
                const type = cell.querySelector('.journal-notes')?.textContent || '';
                headers.push(`${date} ${type}`.trim());
            }
        });
        
        // Получаем данные студентов
        const rows = document.querySelectorAll('.journal-table tbody tr');
        const data = [];
        
        rows.forEach(row => {
            // Пропускаем строки с сообщением об отсутствии данных
            if (row.classList.contains('no-data')) return;
            
            const rowData = [];
            
            // Имя студента
            const studentName = row.querySelector('.student-name span').textContent;
            rowData.push(studentName);
            
            // Оценки
            const cells = row.querySelectorAll('td');
            cells.forEach((cell, index) => {
                if (index === 0) return; // Пропускаем первую ячейку с именем
                
                const input = cell.querySelector('.grade-input');
                if (input) {
                    rowData.push(input.value);
                } else {
                    rowData.push(cell.textContent.trim());
                }
            });
            
            data.push(rowData);
        });
        
        // Создаем содержимое CSV
        let csvContent = "\uFEFF"; // BOM для поддержки кириллицы
        
        // Добавляем заголовки
        csvContent += headers.join(";") + "\r\n";
        
        // Добавляем данные
        data.forEach(row => {
            csvContent += row.join(";") + "\r\n";
        });
        
        // Создаем Blob и ссылку для скачивания
        const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8" });
        const url = URL.createObjectURL(blob);
        
        // Определяем имя предмета для имени файла
        let subjectName = 'Математика';
        if (currentSubject === 'geometry') {
            subjectName = 'Геометрия';
        } else if (currentSubject === 'algebra') {
            subjectName = 'Алгебра';
        }
        
        // Создаем имя файла с датой
        const date = new Date();
        const fileName = `Журнал_${subjectName}_${currentGroup.toUpperCase()}_${date.toLocaleDateString()}.csv`;
        
        // Создаем ссылку и эмулируем клик для скачивания
        const link = document.createElement("a");
        link.setAttribute("href", url);
        link.setAttribute("download", fileName);
        link.style.display = "none";
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    }

    // Инициализация
    loadTeacherProfile();
    setupTabs();
    setupFilters();
    setupActionButtons();
    setupExport();
    
    // Инициализация с данными первой группы и предмета по умолчанию
    renderStudentList();
}); 