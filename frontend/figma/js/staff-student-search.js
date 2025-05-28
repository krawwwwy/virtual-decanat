// JavaScript для страницы поиска студента сотрудником деканата
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
    
    // Загрузка профиля сотрудника
    function loadStaffProfile() {
        try {
            if (!checkAuth()) return;
            
            // Получаем имя пользователя из localStorage
            const userInfo = JSON.parse(localStorage.getItem('userInfo') || '{}');
            document.getElementById('userName').textContent = userInfo.name || 'Александрова А. А.';
            
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
    
    // Функция поиска студента (имитация API-запроса)
    async function searchStudent(studentId) {
        // В реальном приложении здесь был бы запрос к API
        return new Promise(resolve => {
            setTimeout(() => {
                // Демонстрационные данные
                if (studentId === '12345' || studentId === '') {
                    resolve({
                        name: 'Иванов Иван Иванович',
                        group: 'ИДБ-22-11',
                        rating: 42.3,
                        attendance: {
                            percent: 73.2,
                            status: 'Хорошая посещаемость'
                        },
                        debts: [
                            {
                                subject: 'Дискретная математика',
                                type: 'Домашнее задание',
                                deadline: '15.02.2025'
                            },
                            {
                                subject: 'Физика',
                                type: 'Лабораторная работа',
                                deadline: '21.03.2025'
                            },
                            {
                                subject: 'Машинное обучение',
                                type: 'Реферат',
                                deadline: '01.04.2025'
                            }
                        ],
                        subjects: [
                            {
                                name: 'Дискретная математика',
                                rating: 45.00,
                                progress: 100,
                                attendance: 100,
                                onTime: true
                            },
                            {
                                name: 'Физика',
                                rating: 45.00,
                                progress: 80,
                                attendance: 95,
                                onTime: true
                            },
                            {
                                name: 'Начертательная геометрия',
                                rating: 35.00,
                                progress: 60,
                                attendance: 90,
                                onTime: false
                            },
                            {
                                name: 'Прикладная физкультура',
                                rating: 54.00,
                                progress: 100,
                                attendance: 100,
                                onTime: true
                            },
                            {
                                name: 'Машинное обучение',
                                rating: 45.00,
                                progress: 80,
                                attendance: 95,
                                onTime: true
                            }
                        ]
                    });
                } else if (studentId === '67890') {
                    resolve({
                        name: 'Петрова Мария Сергеевна',
                        group: 'ИДБ-22-09',
                        rating: 47.8,
                        attendance: {
                            percent: 85.6,
                            status: 'Высокая посещаемость'
                        },
                        debts: [
                            {
                                subject: 'Программирование',
                                type: 'Курсовая работа',
                                deadline: '10.03.2025'
                            }
                        ],
                        subjects: [
                            {
                                name: 'Дискретная математика',
                                rating: 50.00,
                                progress: 100,
                                attendance: 100,
                                onTime: true
                            },
                            {
                                name: 'Физика',
                                rating: 48.00,
                                progress: 90,
                                attendance: 95,
                                onTime: true
                            },
                            {
                                name: 'Программирование',
                                rating: 42.00,
                                progress: 75,
                                attendance: 90,
                                onTime: false
                            },
                            {
                                name: 'Прикладная физкультура',
                                rating: 50.00,
                                progress: 100,
                                attendance: 100,
                                onTime: true
                            }
                        ]
                    });
                } else if (studentId === '123123') {
                    resolve({
                        name: 'Смирнов Алексей Владимирович',
                        group: 'ИДБ-21-05',
                        rating: 38.7,
                        attendance: {
                            percent: 65.4,
                            status: 'Средняя посещаемость'
                        },
                        debts: [
                            {
                                subject: 'Высшая математика',
                                type: 'Контрольная работа',
                                deadline: '20.02.2025'
                            },
                            {
                                subject: 'Иностранный язык',
                                type: 'Эссе',
                                deadline: '05.03.2025'
                            },
                            {
                                subject: 'Информационная безопасность',
                                type: 'Лабораторная работа',
                                deadline: '12.03.2025'
                            },
                            {
                                subject: 'Базы данных',
                                type: 'Курсовой проект',
                                deadline: '25.03.2025'
                            }
                        ],
                        subjects: [
                            {
                                name: 'Высшая математика',
                                rating: 32.00,
                                progress: 45,
                                attendance: 60,
                                onTime: false
                            },
                            {
                                name: 'Информационная безопасность',
                                rating: 40.00,
                                progress: 70,
                                attendance: 75,
                                onTime: false
                            },
                            {
                                name: 'Базы данных',
                                rating: 38.00,
                                progress: 65,
                                attendance: 70,
                                onTime: false
                            },
                            {
                                name: 'Иностранный язык',
                                rating: 36.00,
                                progress: 60,
                                attendance: 65,
                                onTime: false
                            },
                            {
                                name: 'Архитектура компьютеров',
                                rating: 43.00,
                                progress: 80,
                                attendance: 70,
                                onTime: true
                            }
                        ]
                    });
                } else if (studentId === '987654') {
                    resolve({
                        name: 'Кузнецова Елена Дмитриевна',
                        group: 'ИДБ-20-03',
                        rating: 56.2,
                        attendance: {
                            percent: 94.8,
                            status: 'Отличная посещаемость'
                        },
                        debts: [],
                        subjects: [
                            {
                                name: 'Теория алгоритмов',
                                rating: 55.00,
                                progress: 96,
                                attendance: 100,
                                onTime: true
                            },
                            {
                                name: 'Компьютерная графика',
                                rating: 58.00,
                                progress: 100,
                                attendance: 100,
                                onTime: true
                            },
                            {
                                name: 'Искусственный интеллект',
                                rating: 57.00,
                                progress: 98,
                                attendance: 95,
                                onTime: true
                            },
                            {
                                name: 'Методы оптимизации',
                                rating: 54.00,
                                progress: 95,
                                attendance: 90,
                                onTime: true
                            },
                            {
                                name: 'Проектирование ПО',
                                rating: 57.00,
                                progress: 98,
                                attendance: 95,
                                onTime: true
                            }
                        ]
                    });
                } else if (studentId === '555555') {
                    resolve({
                        name: 'Соколов Дмитрий Андреевич',
                        group: 'ИДБ-21-02',
                        rating: 44.5,
                        attendance: {
                            percent: 80.2,
                            status: 'Хорошая посещаемость'
                        },
                        debts: [
                            {
                                subject: 'Операционные системы',
                                type: 'Практическое задание',
                                deadline: '05.02.2025'
                            }
                        ],
                        subjects: [
                            {
                                name: 'Операционные системы',
                                rating: 42.00,
                                progress: 75,
                                attendance: 85,
                                onTime: false
                            },
                            {
                                name: 'Вычислительная техника',
                                rating: 45.00,
                                progress: 82,
                                attendance: 90,
                                onTime: true
                            },
                            {
                                name: 'Сети и телекоммуникации',
                                rating: 47.00,
                                progress: 85,
                                attendance: 80,
                                onTime: true
                            },
                            {
                                name: 'Веб-программирование',
                                rating: 48.00,
                                progress: 90,
                                attendance: 85,
                                onTime: true
                            }
                        ]
                    });
                } else {
                    resolve(null); // Студент не найден
                }
            }, 500); // Имитация задержки сети
        });
    }
    
    // Отображение данных студента
    function displayStudentInfo(student) {
        const studentInfoContainer = document.getElementById('studentInfo');
        
        if (!student) {
            // Если студент не найден
            studentInfoContainer.innerHTML = '<div class="student-not-found">Студент с указанным номером студенческого билета не найден.</div>';
            return;
        }
        
        // Обновление карточки студента
        const studentCard = `
            <div class="student-card">
                <div class="student-header">
                    <div>
                        <div class="student-name">${student.name}</div>
                        <div class="student-group">Группа ${student.group}</div>
                    </div>
                    <div class="student-rating">${student.rating}</div>
                </div>
            </div>
            
            <div class="attendance-card">
                <div class="attendance-title">Процент посещаемости</div>
                <div class="attendance-value">${student.attendance.percent}%</div>
                <div class="attendance-status">${student.attendance.status}</div>
            </div>
        `;
        
        // Таблица задолженностей
        let debtsTable = '';
        
        if (student.debts.length > 0) {
            debtsTable = `
                <section class="debts-section">
                    <h2 class="section-title">Задолженности студента</h2>
                    <table class="debts-table">
                        <thead>
                            <tr>
                                <th>Предмет</th>
                                <th>Вид задолженности</th>
                                <th>Крайний срок</th>
                            </tr>
                        </thead>
                        <tbody>
                            ${student.debts.map(debt => `
                                <tr>
                                    <td>${debt.subject}</td>
                                    <td>${debt.type}</td>
                                    <td>${debt.deadline}</td>
                                </tr>
                            `).join('')}
                        </tbody>
                    </table>
                    <button class="btn-edit">Редактировать задолженности</button>
                </section>
            `;
        } else {
            debtsTable = `
                <section class="debts-section">
                    <h2 class="section-title">Задолженности студента</h2>
                    <div style="background-color: var(--card-bg); border: 1px solid var(--border-color); border-radius: 12px; padding: 20px; text-align: center;">
                        <p>У студента нет задолженностей</p>
                    </div>
                </section>
            `;
        }
        
        // Таблица рейтинга по предметам
        const ratingsTable = `
            <section class="ratings-section">
                <h2 class="section-title">Рейтинг по предметам студента</h2>
                <table class="ratings-table">
                    <thead>
                        <tr>
                            <th>Предмет</th>
                            <th>Рейтинг</th>
                            <th>Прогресс</th>
                            <th>Посещаемость</th>
                            <th>Сделано вовремя</th>
                        </tr>
                    </thead>
                    <tbody>
                        ${student.subjects.map(subject => `
                            <tr>
                                <td>${subject.name}</td>
                                <td>${subject.rating.toFixed(2)}</td>
                                <td class="progress-cell">
                                    <div class="progress-container">
                                        <div class="progress-bar" style="width: ${subject.progress}%;"></div>
                                    </div>
                                    <span class="progress-value">${subject.progress}</span>
                                </td>
                                <td class="progress-cell">
                                    <div class="progress-container">
                                        <div class="progress-bar" style="width: ${subject.attendance}%;"></div>
                                    </div>
                                    <span class="progress-value">${subject.attendance}</span>
                                </td>
                                <td><button class="btn-${subject.onTime ? 'yes' : 'no'}">${subject.onTime ? 'Да' : 'Нет'}</button></td>
                            </tr>
                        `).join('')}
                    </tbody>
                </table>
                <button class="btn-edit">Редактировать рейтинг</button>
            </section>
        `;
        
        // Обновляем содержимое контейнера с информацией о студенте
        studentInfoContainer.innerHTML = studentCard + debtsTable + ratingsTable;
        
        // Добавляем обработчики для кнопок редактирования
        setupEditButtons();
    }
    
    // Настройка обработчиков для кнопок редактирования
    function setupEditButtons() {
        const editButtons = document.querySelectorAll('.btn-edit');
        editButtons.forEach(button => {
            button.addEventListener('click', function() {
                // В реальном приложении здесь открывалась бы форма редактирования
                alert('Функциональность редактирования будет добавлена в следующей версии.');
            });
        });
    }
    
    // Инициализация поиска
    function initSearch() {
        const searchForm = document.querySelector('.search-container');
        const searchInput = document.querySelector('.search-input');
        const searchButton = document.querySelector('.search-button');
        
        if (searchForm && searchInput && searchButton) {
            // Обработчик клика на кнопку поиска
            searchButton.addEventListener('click', async function() {
                const studentId = searchInput.value.trim();
                const student = await searchStudent(studentId);
                displayStudentInfo(student);
            });
            
            // Обработчик нажатия Enter в поле ввода
            searchInput.addEventListener('keypress', async function(event) {
                if (event.key === 'Enter') {
                    event.preventDefault();
                    const studentId = searchInput.value.trim();
                    const student = await searchStudent(studentId);
                    displayStudentInfo(student);
                }
            });
        }
    }
    
    // Инициализация страницы
    function init() {
        loadStaffProfile();
        initSearch();
        
        // Загружаем данные студента по умолчанию для демонстрации
        searchStudent('').then(displayStudentInfo);
        
        // Добавляем подсказку с доступными номерами студенческих билетов
        const searchInput = document.querySelector('.search-input');
        if (searchInput) {
            searchInput.title = 'Доступные номера для демонстрации: 12345, 67890, 123123, 987654, 555555';
        }
    }
    
    // Запуск инициализации
    init();
}); 