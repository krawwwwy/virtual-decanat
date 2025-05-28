// Функции для работы с журналом успеваемости
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
        // Обработчик выбора предмета
        const subjectFilter = document.getElementById('subjectFilter');
        if (subjectFilter) {
            subjectFilter.addEventListener('change', function() {
                const selectedSubject = this.value;
                showLoading(true);
                
                setTimeout(() => {
                    // Обновляем информацию о предмете
                    const subjectTitle = document.querySelector('.journal-title');
                    
                    if (selectedSubject === 'math') {
                        subjectTitle.textContent = 'Прикладная математика';
                        document.querySelector('.subject-icon img').src = '../../assets/images/teacher/math_icon.svg';
                    } else if (selectedSubject === 'geometry') {
                        subjectTitle.textContent = 'Начертательная геометрия';
                        document.querySelector('.subject-icon img').src = '../../assets/images/teacher/geometry_icon.svg';
                    } else if (selectedSubject === 'algebra') {
                        subjectTitle.textContent = 'Линейная алгебра';
                        document.querySelector('.subject-icon img').src = '../../assets/images/teacher/math_icon.svg';
                    }
                    
                    // Здесь могла бы быть загрузка оценок для выбранного предмета
                    showLoading(false);
                }, 500);
            });
        }
        
        // Обработчик выбора группы
        const groupFilter = document.getElementById('groupFilter');
        if (groupFilter) {
            groupFilter.addEventListener('change', function() {
                const selectedGroup = this.value;
                showLoading(true);
                
                setTimeout(() => {
                    // Обновляем информацию о группе
                    const groupInfo = document.querySelector('.journal-info > div:last-child > div:last-child');
                    groupInfo.textContent = `Группа: ${selectedGroup.toUpperCase()} | Семестр: Весенний 2025`;
                    
                    // Здесь могла бы быть загрузка оценок для выбранной группы
                    showLoading(false);
                }, 500);
            });
        }
    }
    
    // Обработка ввода оценок
    function setupGradeInputs() {
        const inputs = document.querySelectorAll('.grade-input');
        inputs.forEach(input => {
            // Сохраняем исходное значение для определения изменений
            input.setAttribute('data-original', input.value);
            
            input.addEventListener('input', function() {
                // Если ввели "н" или "Н", сделать красным (пропуск)
                if (this.value.toLowerCase() === 'н') {
                    this.classList.add('grade-absent');
                } else {
                    this.classList.remove('grade-absent');
                }
                
                // Визуальная отметка измененных значений
                if (this.value !== this.getAttribute('data-original')) {
                    this.style.backgroundColor = 'rgba(255, 255, 0, 0.1)';
                } else {
                    this.style.backgroundColor = '';
                }
                
                // Автоматический расчет среднего балла, если это не ячейка среднего балла или экзамена
                if (!this.closest('td').classList.contains('avg-grade') && 
                    !this.closest('td').classList.contains('exam-grade')) {
                    updateAverageGrade(this.closest('tr'));
                }
            });
        });
    }
    
    // Расчет среднего балла для строки
    function updateAverageGrade(row) {
        // Находим все обычные ячейки с оценками (не средний балл и не экзамен)
        const gradeInputs = Array.from(row.querySelectorAll('td:not(.avg-grade):not(.exam-grade) .grade-input'));
        
        // Вычисляем средний балл
        let totalGrade = 0;
        let validGradesCount = 0;
        
        gradeInputs.forEach(input => {
            // Проверяем, является ли значение числом
            if (!isNaN(input.value) && input.value.trim() !== '') {
                totalGrade += parseInt(input.value);
                validGradesCount++;
            }
        });
        
        // Обновляем средний балл
        if (validGradesCount > 0) {
            const avgGrade = totalGrade / validGradesCount;
            const avgCell = row.querySelector('td.avg-grade:nth-of-type(9)');
            if (avgCell) {
                avgCell.textContent = avgGrade.toFixed(1);
                avgCell.style.backgroundColor = 'rgba(26, 26, 229, 0.1)';
            }
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
                    // Сбор данных об оценках
                    const gradesData = [];
                    const rows = document.querySelectorAll('.journal-table tbody tr');
                    
                    rows.forEach(row => {
                        const studentName = row.querySelector('.student-name span').textContent;
                        const grades = Array.from(row.querySelectorAll('.grade-input')).map(input => {
                            return {
                                value: input.value,
                                wasChanged: input.value !== input.getAttribute('data-original')
                            };
                        });
                        
                        gradesData.push({
                            studentName,
                            grades
                        });
                        
                        // Сбрасываем фон у измененных ячеек
                        row.querySelectorAll('.grade-input').forEach(input => {
                            input.style.backgroundColor = '';
                            input.setAttribute('data-original', input.value);
                        });
                    });
                    
                    console.log('Данные оценок сохранены:', gradesData);
                    showLoading(false);
                    alert('Оценки успешно сохранены!');
                    
                    // В реальном проекте здесь был бы API-запрос для сохранения данных
                }, 800);
            });
        }
        
        const cancelButton = document.querySelector('.cancel-button');
        if (cancelButton) {
            cancelButton.addEventListener('click', function() {
                if (confirm('Вы уверены, что хотите отменить все изменения оценок?')) {
                    // Восстанавливаем оригинальные значения
                    document.querySelectorAll('.grade-input').forEach(input => {
                        input.value = input.getAttribute('data-original');
                        input.style.backgroundColor = '';
                        
                        // Если восстановилось значение "н", добавляем класс
                        if (input.value.toLowerCase() === 'н') {
                            input.classList.add('grade-absent');
                        } else {
                            input.classList.remove('grade-absent');
                        }
                    });
                    
                    // Восстанавливаем средние баллы
                    document.querySelectorAll('.journal-table tbody tr').forEach(row => {
                        updateAverageGrade(row);
                    });
                }
            });
        }
    }
    
    // Инициализация всех функций
    loadTeacherProfile();
    setupTabs();
    setupFilters();
    setupGradeInputs();
    setupActionButtons();
}); 