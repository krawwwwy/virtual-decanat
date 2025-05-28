// Функция для проверки авторизации
function checkAuth() {
    // Получение информации о пользователе из localStorage
    const userInfo = JSON.parse(localStorage.getItem('userInfo'));
    
    // Если пользователь не авторизован или не является преподавателем, перенаправляем на страницу входа
    if (!userInfo || userInfo.role !== 'teacher') {
        window.location.href = '../../screens/login-page.html';
    }
    
    // Отображение имени пользователя
    if (userInfo) {
        updateProfileInfo(userInfo);
    }
}

// Функция обновления информации профиля
function updateProfileInfo(userInfo) {
    const profileName = document.querySelector('.profile-name');
    const profileAvatar = document.querySelector('.profile-avatar');
    
    if (profileName) {
        profileName.textContent = `${userInfo.firstName} ${userInfo.lastName}`;
    }
    
    if (profileAvatar) {
        // Если у пользователя есть аватар, используем его, иначе используем инициалы
        if (userInfo.avatarUrl) {
            profileAvatar.innerHTML = `<img src="${userInfo.avatarUrl}" alt="${userInfo.firstName}">`;
        } else {
            const initials = `${userInfo.firstName.charAt(0)}${userInfo.lastName.charAt(0)}`;
            profileAvatar.textContent = initials;
        }
    }
}

// Обработка переключения вкладок
function setupTabs() {
    const tabs = document.querySelectorAll('.tab');
    
    tabs.forEach(tab => {
        tab.addEventListener('click', () => {
            // Удаляем активный класс со всех вкладок
            tabs.forEach(t => t.classList.remove('active'));
            
            // Добавляем активный класс к выбранной вкладке
            tab.classList.add('active');
            
            // Обновляем отображение контента в зависимости от выбранной вкладки
            updateContentByTab(tab.textContent.trim());
        });
    });
}

// Обновление отображаемого контента в зависимости от выбранной вкладки
function updateContentByTab(tabName) {
    console.log(`Переключение на вкладку: ${tabName}`);
    
    // Здесь будет логика для обновления содержимого страницы
    // в зависимости от выбранной вкладки (Группы, Студенты, Оценки)
    
    // Например, можно загружать соответствующие данные с сервера
    // fetchDataForTab(tabName);
}

// Настройка поиска по дисциплинам
function setupSearch() {
    const searchInput = document.querySelector('.search-input');
    
    if (searchInput) {
        searchInput.addEventListener('input', (e) => {
            const searchText = e.target.value.toLowerCase();
            filterSubjects(searchText);
        });
    }
}

// Фильтрация дисциплин по поисковому запросу
function filterSubjects(searchText) {
    const subjectCards = document.querySelectorAll('.subject-card');
    
    subjectCards.forEach(card => {
        const title = card.querySelector('.subject-title').textContent.toLowerCase();
        const lecturer = card.querySelector('.subject-lecturer').textContent.toLowerCase();
        
        if (title.includes(searchText) || lecturer.includes(searchText)) {
            card.style.display = 'flex';
        } else {
            card.style.display = 'none';
        }
    });
}

// Настройка кнопки добавления дисциплины
function setupAddButton() {
    const addButton = document.querySelector('.btn-primary');
    
    if (addButton) {
        addButton.addEventListener('click', () => {
            showAddSubjectModal();
        });
    }
}

// Отображение модального окна для добавления новой дисциплины
function showAddSubjectModal() {
    // Создаем модальное окно
    const modal = document.createElement('div');
    modal.className = 'modal';
    
    modal.innerHTML = `
        <div class="modal-content">
            <span class="close-modal">&times;</span>
            <h2>Добавить новую дисциплину</h2>
            <form id="add-subject-form">
                <div class="form-group">
                    <label for="subject-name">Название дисциплины</label>
                    <input type="text" id="subject-name" required>
                </div>
                <div class="form-group">
                    <label for="subject-lecturer">Преподаватель</label>
                    <input type="text" id="subject-lecturer" required>
                </div>
                <button type="submit" class="btn btn-primary">Добавить</button>
            </form>
        </div>
    `;
    
    // Добавляем модальное окно в DOM
    document.body.appendChild(modal);
    
    // Получаем элементы
    const closeButton = modal.querySelector('.close-modal');
    const form = modal.querySelector('#add-subject-form');
    
    // Настраиваем закрытие модального окна
    closeButton.addEventListener('click', () => {
        document.body.removeChild(modal);
    });
    
    // Закрытие при клике вне содержимого модального окна
    modal.addEventListener('click', (e) => {
        if (e.target === modal) {
            document.body.removeChild(modal);
        }
    });
    
    // Обрабатываем отправку формы
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        
        const subjectName = document.getElementById('subject-name').value;
        const subjectLecturer = document.getElementById('subject-lecturer').value;
        
        // Здесь будет вызов API для добавления дисциплины
        console.log(`Добавление новой дисциплины: ${subjectName}, ${subjectLecturer}`);
        
        // Добавляем новую карточку дисциплины
        addSubjectCard(subjectName, subjectLecturer);
        
        // Закрываем модальное окно
        document.body.removeChild(modal);
    });
}

// Добавление новой карточки дисциплины
function addSubjectCard(name, lecturer) {
    const subjectsGrid = document.querySelector('.subjects-grid');
    
    const newCard = document.createElement('div');
    newCard.className = 'subject-card';
    newCard.innerHTML = `
        <div class="subject-icon">
            <img src="../../assets/images/teacher/subject-icon.svg" alt="${name}">
        </div>
        <div class="subject-info">
            <h3 class="subject-title">${name}</h3>
            <p class="subject-lecturer">Преподаватель: ${lecturer}</p>
        </div>
        <div class="subject-actions">
            <button class="btn-icon">
                <img src="../../assets/images/teacher/dashboard-icon.svg" alt="Опции">
            </button>
        </div>
    `;
    
    subjectsGrid.appendChild(newCard);
}

// Добавляем стили для модального окна
function addModalStyles() {
    const styleElement = document.createElement('style');
    styleElement.textContent = `
        .modal {
            display: flex;
            position: fixed;
            z-index: 1000;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            background-color: rgba(0, 0, 0, 0.5);
            align-items: center;
            justify-content: center;
        }
        
        .modal-content {
            background-color: #fff;
            padding: 24px;
            border-radius: 8px;
            width: 90%;
            max-width: 500px;
            position: relative;
        }
        
        .close-modal {
            position: absolute;
            right: 16px;
            top: 16px;
            font-size: 24px;
            cursor: pointer;
        }
        
        .form-group {
            margin-bottom: 16px;
        }
        
        .form-group label {
            display: block;
            margin-bottom: 8px;
            font-weight: 500;
        }
        
        .form-group input {
            width: 100%;
            padding: 8px;
            border: 1px solid #DBDBE5;
            border-radius: 4px;
        }
        
        #add-subject-form .btn-primary {
            margin-top: 16px;
        }
    `;
    
    document.head.appendChild(styleElement);
}

// Инициализация страницы
function init() {
    checkAuth();
    setupTabs();
    setupSearch();
    setupAddButton();
    addModalStyles();
    
    // Настраиваем обработчики для кнопок действий дисциплин
    setupSubjectActions();
}

// Обработчики для кнопок действий дисциплин
function setupSubjectActions() {
    document.addEventListener('click', (e) => {
        // Находим ближайший родительский элемент кнопки с классом btn-icon
        const button = e.target.closest('.btn-icon');
        
        if (button) {
            // Находим ближайшую карточку дисциплины
            const card = button.closest('.subject-card');
            
            if (card) {
                const title = card.querySelector('.subject-title').textContent;
                showSubjectActionsMenu(button, title);
            }
            
            // Находим ближайшую карточку студента
            const studentCard = button.closest('.student-card');
            
            if (studentCard) {
                const name = studentCard.querySelector('.student-name').textContent;
                showStudentActionsMenu(button, name);
            }
        }
    });
}

// Отображение меню действий для дисциплины
function showSubjectActionsMenu(button, subjectName) {
    // Создаем контекстное меню
    const menu = document.createElement('div');
    menu.className = 'context-menu';
    menu.innerHTML = `
        <ul>
            <li data-action="edit">Редактировать</li>
            <li data-action="delete">Удалить</li>
        </ul>
    `;
    
    // Позиционируем меню возле кнопки
    const rect = button.getBoundingClientRect();
    menu.style.position = 'absolute';
    menu.style.top = `${rect.bottom + window.scrollY}px`;
    menu.style.left = `${rect.left + window.scrollX}px`;
    menu.style.backgroundColor = '#fff';
    menu.style.boxShadow = '0 2px 10px rgba(0, 0, 0, 0.1)';
    menu.style.borderRadius = '4px';
    menu.style.zIndex = '100';
    
    menu.querySelector('ul').style.listStyle = 'none';
    menu.querySelector('ul').style.padding = '8px 0';
    menu.querySelector('ul').style.margin = '0';
    
    const items = menu.querySelectorAll('li');
    items.forEach(item => {
        item.style.padding = '8px 16px';
        item.style.cursor = 'pointer';
        item.addEventListener('mouseenter', () => {
            item.style.backgroundColor = '#F0F0F5';
        });
        item.addEventListener('mouseleave', () => {
            item.style.backgroundColor = 'transparent';
        });
    });
    
    // Добавляем обработчики действий
    menu.querySelector('[data-action="edit"]').addEventListener('click', () => {
        console.log(`Редактирование дисциплины: ${subjectName}`);
        document.body.removeChild(menu);
        // Здесь будет функция редактирования дисциплины
    });
    
    menu.querySelector('[data-action="delete"]').addEventListener('click', () => {
        console.log(`Удаление дисциплины: ${subjectName}`);
        document.body.removeChild(menu);
        // Здесь будет функция удаления дисциплины
    });
    
    // Добавляем меню в DOM
    document.body.appendChild(menu);
    
    // Закрываем меню при клике вне его
    document.addEventListener('click', function closeMenu(e) {
        if (!menu.contains(e.target) && e.target !== button) {
            document.body.removeChild(menu);
            document.removeEventListener('click', closeMenu);
        }
    });
}

// Отображение меню действий для студента
function showStudentActionsMenu(button, studentName) {
    // Создаем контекстное меню
    const menu = document.createElement('div');
    menu.className = 'context-menu';
    menu.innerHTML = `
        <ul>
            <li data-action="view">Просмотр профиля</li>
            <li data-action="grade">Выставить оценку</li>
        </ul>
    `;
    
    // Позиционируем меню возле кнопки
    const rect = button.getBoundingClientRect();
    menu.style.position = 'absolute';
    menu.style.top = `${rect.bottom + window.scrollY}px`;
    menu.style.left = `${rect.left + window.scrollX}px`;
    menu.style.backgroundColor = '#fff';
    menu.style.boxShadow = '0 2px 10px rgba(0, 0, 0, 0.1)';
    menu.style.borderRadius = '4px';
    menu.style.zIndex = '100';
    
    menu.querySelector('ul').style.listStyle = 'none';
    menu.querySelector('ul').style.padding = '8px 0';
    menu.querySelector('ul').style.margin = '0';
    
    const items = menu.querySelectorAll('li');
    items.forEach(item => {
        item.style.padding = '8px 16px';
        item.style.cursor = 'pointer';
        item.addEventListener('mouseenter', () => {
            item.style.backgroundColor = '#F0F0F5';
        });
        item.addEventListener('mouseleave', () => {
            item.style.backgroundColor = 'transparent';
        });
    });
    
    // Добавляем обработчики действий
    menu.querySelector('[data-action="view"]').addEventListener('click', () => {
        console.log(`Просмотр профиля студента: ${studentName}`);
        document.body.removeChild(menu);
        // Здесь будет функция просмотра профиля студента
    });
    
    menu.querySelector('[data-action="grade"]').addEventListener('click', () => {
        console.log(`Выставление оценки для: ${studentName}`);
        document.body.removeChild(menu);
        showGradeModal(studentName);
    });
    
    // Добавляем меню в DOM
    document.body.appendChild(menu);
    
    // Закрываем меню при клике вне его
    document.addEventListener('click', function closeMenu(e) {
        if (!menu.contains(e.target) && e.target !== button) {
            document.body.removeChild(menu);
            document.removeEventListener('click', closeMenu);
        }
    });
}

// Отображение модального окна для выставления оценки
function showGradeModal(studentName) {
    // Создаем модальное окно
    const modal = document.createElement('div');
    modal.className = 'modal';
    
    modal.innerHTML = `
        <div class="modal-content">
            <span class="close-modal">&times;</span>
            <h2>Выставить оценку для ${studentName}</h2>
            <form id="grade-form">
                <div class="form-group">
                    <label for="subject-select">Дисциплина</label>
                    <select id="subject-select" required>
                        <option value="Математика">Математика</option>
                        <option value="Физика">Физика</option>
                        <option value="Литература">Литература</option>
                        <option value="Биология">Биология</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="grade-value">Оценка</label>
                    <select id="grade-value" required>
                        <option value="A">A</option>
                        <option value="B">B</option>
                        <option value="C">C</option>
                        <option value="D">D</option>
                        <option value="F">F</option>
                    </select>
                </div>
                <div class="form-group">
                    <label for="grade-comment">Комментарий (необязательно)</label>
                    <textarea id="grade-comment" rows="3"></textarea>
                </div>
                <button type="submit" class="btn btn-primary">Сохранить</button>
            </form>
        </div>
    `;
    
    // Добавляем стили для селектов и textarea
    const additionalStyle = document.createElement('style');
    additionalStyle.textContent = `
        select {
            width: 100%;
            padding: 8px;
            border: 1px solid #DBDBE5;
            border-radius: 4px;
        }
        
        textarea {
            width: 100%;
            padding: 8px;
            border: 1px solid #DBDBE5;
            border-radius: 4px;
            resize: vertical;
        }
    `;
    document.head.appendChild(additionalStyle);
    
    // Добавляем модальное окно в DOM
    document.body.appendChild(modal);
    
    // Получаем элементы
    const closeButton = modal.querySelector('.close-modal');
    const form = modal.querySelector('#grade-form');
    
    // Настраиваем закрытие модального окна
    closeButton.addEventListener('click', () => {
        document.body.removeChild(modal);
        document.head.removeChild(additionalStyle);
    });
    
    // Закрытие при клике вне содержимого модального окна
    modal.addEventListener('click', (e) => {
        if (e.target === modal) {
            document.body.removeChild(modal);
            document.head.removeChild(additionalStyle);
        }
    });
    
    // Обрабатываем отправку формы
    form.addEventListener('submit', (e) => {
        e.preventDefault();
        
        const subject = document.getElementById('subject-select').value;
        const grade = document.getElementById('grade-value').value;
        const comment = document.getElementById('grade-comment').value;
        
        // Здесь будет вызов API для сохранения оценки
        console.log(`Выставление оценки: ${studentName}, ${subject}, ${grade}, ${comment}`);
        
        // Обновляем отображение оценки в карточке студента (если это текущая дисциплина)
        updateStudentGrade(studentName, grade);
        
        // Закрываем модальное окно
        document.body.removeChild(modal);
        document.head.removeChild(additionalStyle);
        
        // Показываем уведомление об успехе
        showNotification(`Оценка ${grade} выставлена для ${studentName}`);
    });
}

// Обновление отображения оценки студента
function updateStudentGrade(studentName, grade) {
    const studentCards = document.querySelectorAll('.student-card');
    
    studentCards.forEach(card => {
        const name = card.querySelector('.student-name').textContent;
        
        if (name === studentName) {
            const gradeElement = card.querySelector('.student-grade');
            gradeElement.textContent = `Оценка: ${grade}`;
        }
    });
}

// Отображение уведомления
function showNotification(message) {
    const notification = document.createElement('div');
    notification.className = 'notification';
    notification.textContent = message;
    
    // Стили для уведомления
    notification.style.position = 'fixed';
    notification.style.bottom = '20px';
    notification.style.right = '20px';
    notification.style.backgroundColor = '#4CAF50';
    notification.style.color = 'white';
    notification.style.padding = '12px 20px';
    notification.style.borderRadius = '4px';
    notification.style.boxShadow = '0 2px 10px rgba(0, 0, 0, 0.1)';
    notification.style.zIndex = '1000';
    
    // Добавляем уведомление в DOM
    document.body.appendChild(notification);
    
    // Удаляем уведомление через 3 секунды
    setTimeout(() => {
        document.body.removeChild(notification);
    }, 3000);
}

// Запускаем инициализацию при загрузке документа
document.addEventListener('DOMContentLoaded', init); 