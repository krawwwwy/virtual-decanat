/**
 * JavaScript для работы со страницей расписания учебных групп
 */

document.addEventListener('DOMContentLoaded', function() {
    // Элементы интерфейса
    const groupInput = document.getElementById('groupInput');
    const scheduleTitle = document.querySelector('.schedule-title');
    const todayButton = document.querySelector('.today-button');
    const editButtons = document.querySelectorAll('.action-button:not(.today-button):not(.manage-button)');
    const manageButton = document.querySelector('.manage-button');
    const showMoreButton = document.querySelector('.show-more');
    const scheduleItems = document.querySelectorAll('.schedule-item');
    const editLinks = document.querySelectorAll('.edit-link');
    
    // Демо-данные для групп
    const groups = [
        "ИДБ-22-10", 
        "ИДБ-22-11", 
        "ЭДБ-24-13"
    ];
    
    // Демо-данные расписания по группам
    const scheduleData = {
        "ИДБ-22-10": [
            { day: "Понедельник", date: "21 февраля", subjects: [
                { name: "Дискретная математика", time: "9:00 AM - 10:30 AM", icon: "math-icon.png" },
                { name: "Начертательная геометрия", time: "11:00 AM - 12:30 PM", icon: "geometry-icon.png" },
                { name: "Машинное обучение", time: "1:00 PM - 2:30 PM", icon: "ml-icon.png" },
                { name: "Лабораторная работа Физика", time: "3:00 PM - 4:30 PM", icon: "physics-icon.png" }
            ]},
            { day: "Вторник", date: "22 февраля", subjects: [
                { name: "Проектирование информационных систем", time: "9:00 AM - 10:30 AM", icon: "systems-icon.png" },
                { name: "Основы Web-разработки", time: "11:00 AM - 12:30 PM", icon: "web-icon.png" },
                { name: "Управление интеллектуальными активами", time: "1:00 PM - 2:30 PM", icon: "assets-icon.png" },
                { name: "Информационные системы и технологии", time: "3:00 PM - 4:30 PM", icon: "it-icon.png" }
            ]}
        ],
        "ИДБ-22-11": [
            { day: "Понедельник", date: "21 февраля", subjects: [
                { name: "Машинное обучение", time: "9:00 AM - 10:30 AM", icon: "ml-icon.png" },
                { name: "Дискретная математика", time: "11:00 AM - 12:30 PM", icon: "math-icon.png" },
                { name: "Лабораторная работа Физика", time: "1:00 PM - 2:30 PM", icon: "physics-icon.png" },
                { name: "Начертательная геометрия", time: "3:00 PM - 4:30 PM", icon: "geometry-icon.png" }
            ]},
            { day: "Вторник", date: "22 февраля", subjects: [
                { name: "Основы Web-разработки", time: "9:00 AM - 10:30 AM", icon: "web-icon.png" },
                { name: "Проектирование информационных систем", time: "11:00 AM - 12:30 PM", icon: "systems-icon.png" },
                { name: "Информационные системы и технологии", time: "1:00 PM - 2:30 PM", icon: "it-icon.png" },
                { name: "Управление интеллектуальными активами", time: "3:00 PM - 4:30 PM", icon: "assets-icon.png" }
            ]}
        ],
        "ЭДБ-24-13": [
            { day: "Понедельник", date: "21 февраля", subjects: [
                { name: "Информационные системы и технологии", time: "9:00 AM - 10:30 AM", icon: "it-icon.png" },
                { name: "Управление интеллектуальными активами", time: "11:00 AM - 12:30 PM", icon: "assets-icon.png" },
                { name: "Основы Web-разработки", time: "1:00 PM - 2:30 PM", icon: "web-icon.png" },
                { name: "Проектирование информационных систем", time: "3:00 PM - 4:30 PM", icon: "systems-icon.png" }
            ]},
            { day: "Вторник", date: "22 февраля", subjects: [
                { name: "Лабораторная работа Физика", time: "9:00 AM - 10:30 AM", icon: "physics-icon.png" },
                { name: "Машинное обучение", time: "11:00 AM - 12:30 PM", icon: "ml-icon.png" },
                { name: "Дискретная математика", time: "1:00 PM - 2:30 PM", icon: "math-icon.png" },
                { name: "Начертательная геометрия", time: "3:00 PM - 4:30 PM", icon: "geometry-icon.png" }
            ]}
        ]
    };
    
    // Функция для загрузки расписания группы
    function loadGroupSchedule(groupName) {
        // Проверяем, существует ли группа в нашем списке
        if (!groups.includes(groupName)) {
            alert(`Группа ${groupName} не найдена. Пожалуйста, выберите одну из существующих групп: ${groups.join(', ')}`);
            return;
        }
        
        // Отображаем заголовок с названием группы
        scheduleTitle.textContent = `Расписание учебной группы ${groupName}`;
        
        // В реальном приложении здесь был бы запрос к API
        console.log(`Загружаю расписание для группы: ${groupName}`);
        
        // Имитация загрузки данных (в реальном приложении данные загружались бы с сервера)
        const schedule = scheduleData[groupName];
        if (!schedule) {
            alert(`Для группы ${groupName} нет доступного расписания.`);
            return;
        }
        
        // Можно добавить логику для обновления расписания на странице
        // В данном демо просто сообщаем, что расписание загружено
        alert(`Расписание для группы ${groupName} загружено!`);
    }
    
    // Обработчик события нажатия Enter в поле ввода группы
    if (groupInput) {
        groupInput.addEventListener('keydown', function(e) {
            if (e.key === 'Enter') {
                loadGroupSchedule(this.value.trim());
            }
        });
    }
    
    // Обработчик кнопки "Сегодня"
    if (todayButton) {
        todayButton.addEventListener('click', function() {
            alert('Показываю расписание на сегодня');
            // Логика переключения на текущий день
        });
    }
    
    // Обработчики кнопок редактирования
    editButtons.forEach(button => {
        button.addEventListener('click', function() {
            alert('Открываю режим редактирования расписания');
            // Логика для перехода в режим редактирования
        });
    });
    
    // Обработчик кнопки управления расписанием
    if (manageButton) {
        manageButton.addEventListener('click', function() {
            alert('Открываю панель управления расписанием');
            // Логика для перехода в панель управления расписанием
        });
    }
    
    // Обработчик для "Смотреть больше"
    if (showMoreButton) {
        showMoreButton.addEventListener('click', function() {
            alert('Загружаю больше дней расписания');
            // Логика для загрузки дополнительных дней расписания
        });
    }
    
    // Обработчики для ссылок редактирования занятий
    editLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const subject = this.closest('.schedule-item').querySelector('.subject-title').textContent;
            alert(`Редактирую занятие: ${subject}`);
            // Логика для редактирования конкретного занятия
        });
    });
    
    // Проверка URL на наличие параметра группы
    function checkUrlForGroup() {
        if (window.location.search.includes('group=')) {
            const urlParams = new URLSearchParams(window.location.search);
            const group = urlParams.get('group');
            if (groups.includes(group)) {
                groupInput.value = group;
                loadGroupSchedule(group);
            }
        }
    }
    
    // Инициализация при загрузке страницы
    checkUrlForGroup();
}); 