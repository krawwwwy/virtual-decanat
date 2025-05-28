/**
 * JavaScript для работы со страницей студенческих объединений
 */

document.addEventListener('DOMContentLoaded', function() {
    // Элементы интерфейса
    const filterButtons = document.querySelectorAll('.filter-button');
    const searchInput = document.querySelector('.search-input');
    const searchButton = document.querySelector('.search-button');
    const clubCards = document.querySelectorAll('.club-card');
    const addClubButton = document.querySelector('.add-club-button');
    
    // Демо-данные клубов с категориями для фильтрации
    const clubsData = [
        { 
            id: 1, 
            name: 'IT Клуб "Кодеры"', 
            category: 'IT', 
            members: 42, 
            founded: 2022,
            description: 'Сообщество студентов, интересующихся разработкой программного обеспечения, искусственным интеллектом и Data Science.'
        },
        { 
            id: 2, 
            name: 'Спортивный клуб "Станкин-Атлет"', 
            category: 'Спорт', 
            members: 78, 
            founded: 2018,
            description: 'Спортивное сообщество для студентов, занимающихся различными видами спорта: футбол, баскетбол, волейбол и другие.'
        },
        { 
            id: 3, 
            name: 'Волонтерский центр "Доброе сердце"', 
            category: 'Волонтерство', 
            members: 63, 
            founded: 2019,
            description: 'Объединение студентов, помогающих в организации мероприятий университета и участвующих в социальных проектах города.'
        },
        { 
            id: 4, 
            name: 'Творческая студия "Муза"', 
            category: 'Творчество', 
            members: 35, 
            founded: 2021,
            description: 'Объединение для студентов, увлекающихся различными видами искусства: музыка, театр, рисование, фотография.'
        },
        { 
            id: 5, 
            name: 'Научное общество "Эврика"', 
            category: 'Наука', 
            members: 29, 
            founded: 2020,
            description: 'Сообщество для студентов, интересующихся научной деятельностью, проведением исследований и участием в конференциях.'
        },
        { 
            id: 6, 
            name: 'Клуб робототехники "Механизм"', 
            category: 'IT', 
            members: 22, 
            founded: 2022,
            description: 'Объединение для студентов, занимающихся проектированием и созданием роботов, участвующих в соревнованиях по робототехнике.'
        }
    ];
    
    // Функция для фильтрации объединений
    function filterClubs(category) {
        clubCards.forEach((card, index) => {
            if (category === 'Все объединения' || clubsData[index].category === category) {
                card.style.display = 'block';
            } else {
                card.style.display = 'none';
            }
        });
    }
    
    // Функция для поиска объединений
    function searchClubs(query) {
        const searchTerm = query.toLowerCase().trim();
        
        clubCards.forEach((card, index) => {
            const clubData = clubsData[index];
            const nameMatch = clubData.name.toLowerCase().includes(searchTerm);
            const descMatch = clubData.description.toLowerCase().includes(searchTerm);
            
            if (nameMatch || descMatch) {
                card.style.display = 'block';
            } else {
                card.style.display = 'none';
            }
        });
    }
    
    // Обработчики кнопок фильтрации
    filterButtons.forEach(button => {
        button.addEventListener('click', function() {
            // Убираем активный класс у всех кнопок
            filterButtons.forEach(btn => btn.classList.remove('active'));
            // Добавляем активный класс нажатой кнопке
            this.classList.add('active');
            // Фильтруем объединения
            filterClubs(this.textContent);
        });
    });
    
    // Обработчик поиска
    searchButton.addEventListener('click', function() {
        searchClubs(searchInput.value);
    });
    
    // Обработчик нажатия Enter в поле поиска
    searchInput.addEventListener('keydown', function(e) {
        if (e.key === 'Enter') {
            searchClubs(this.value);
        }
    });
    
    // Обработчик кнопки добавления нового объединения
    addClubButton.addEventListener('click', function() {
        showAddClubModal();
    });
    
    // Функция для отображения модального окна добавления объединения
    function showAddClubModal() {
        alert('Открытие формы для добавления нового студенческого объединения');
        // В реальном приложении здесь был бы код для отображения модального окна с формой
    }
    
    // Обработчики для кнопок "Редактировать" и "Участники"
    document.querySelectorAll('.club-button').forEach(button => {
        button.addEventListener('click', function(e) {
            e.preventDefault();
            const clubCard = this.closest('.club-card');
            const clubName = clubCard.querySelector('.club-name').textContent;
            
            if (this.textContent === 'Редактировать') {
                alert(`Редактирование объединения: ${clubName}`);
                // В реальном приложении здесь был бы код для перехода на страницу редактирования
            } else if (this.textContent === 'Участники') {
                alert(`Просмотр участников объединения: ${clubName}`);
                // В реальном приложении здесь был бы код для отображения списка участников
            }
        });
    });
}); 