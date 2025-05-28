/**
 * JavaScript для страницы студенческих клубов
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем авторизацию
    if (!window.authApi.isAuthenticated() || localStorage.getItem('userRole') !== 'student') {
        window.location.href = '../login-page.html';
        return;
    }

    // Инициализация фильтрации клубов
    const filterTags = document.querySelectorAll('.filter-tag');
    const clubCards = document.querySelectorAll('.club-card[data-category]');
    const searchInput = document.getElementById('club-search');
    
    // Переменные для хранения состояния фильтров
    let activeFilters = new Set();
    let searchQuery = '';

    // Обработка клика на тегах фильтрации
    filterTags.forEach(tag => {
        tag.addEventListener('click', function() {
            const filter = this.getAttribute('data-filter');
            
            // Переключение активного состояния фильтра
            if (this.classList.contains('active')) {
                this.classList.remove('active');
                activeFilters.delete(filter);
            } else {
                this.classList.add('active');
                activeFilters.add(filter);
            }
            
            // Применяем фильтрацию
            applyFilters();
        });
    });

    // Обработка ввода в поле поиска
    if (searchInput) {
        searchInput.addEventListener('input', function() {
            searchQuery = this.value.toLowerCase().trim();
            applyFilters();
        });
    }

    // Кнопки "Отмена" и "Применить"
    const cancelButton = document.querySelector('.button-secondary');
    const applyButton = document.querySelector('.button-primary');
    
    if (cancelButton) {
        cancelButton.addEventListener('click', function() {
            // Сбрасываем все фильтры
            filterTags.forEach(tag => tag.classList.remove('active'));
            activeFilters.clear();
            
            // Сбрасываем поле поиска
            if (searchInput) {
                searchInput.value = '';
                searchQuery = '';
            }
            
            // Показываем все клубы
            clubCards.forEach(card => {
                card.style.display = 'flex';
            });
            
            // Уведомление о сбросе фильтров
            showNotification('Фильтры сброшены');
        });
    }
    
    if (applyButton) {
        applyButton.addEventListener('click', function() {
            // Применяем фильтры
            applyFilters();
            showNotification('Фильтры применены');
        });
    }

    // Функция фильтрации карточек клубов
    function applyFilters() {
        // Если нет активных фильтров и поиск пуст, показываем все карточки
        if (activeFilters.size === 0 && searchQuery === '') {
            clubCards.forEach(card => {
                card.style.display = 'flex';
            });
            return;
        }

        // Фильтруем карточки
        clubCards.forEach(card => {
            const category = card.getAttribute('data-category');
            const title = card.querySelector('.club-title').textContent.toLowerCase();
            const description = card.querySelector('.club-schedule').textContent.toLowerCase();
            
            // Проверяем совпадение по фильтрам категорий
            const matchesCategory = activeFilters.size === 0 || activeFilters.has(category);
            
            // Проверяем совпадение по поисковому запросу
            const matchesSearch = searchQuery === '' || 
                title.includes(searchQuery) || 
                description.includes(searchQuery);
            
            // Если карточка соответствует всем критериям, показываем ее
            if (matchesCategory && matchesSearch) {
                card.style.display = 'flex';
            } else {
                card.style.display = 'none';
            }
        });
    }

    // Обработчики для кнопок "Вступить"
    const joinButtons = document.querySelectorAll('.club-join-button');
    joinButtons.forEach(button => {
        button.addEventListener('click', function() {
            const clubCard = this.closest('.club-card');
            const clubTitle = clubCard.querySelector('.club-title').textContent;
            
            // Показываем модальное окно для подтверждения
            showJoinConfirmation(clubTitle, clubCard);
        });
    });

    /**
     * Показать модальное окно подтверждения вступления в клуб
     */
    function showJoinConfirmation(clubTitle, clubCard) {
        // Создаем модальное окно
        const modal = document.createElement('div');
        modal.className = 'modal';
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
        modalContent.style.width = '400px';
        modalContent.style.maxWidth = '90%';
        modalContent.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.15)';
        
        // Заголовок модального окна
        const modalTitle = document.createElement('h3');
        modalTitle.textContent = `Вступить в клуб "${clubTitle}"?`;
        modalTitle.style.fontFamily = 'Inter, sans-serif';
        modalTitle.style.fontWeight = '700';
        modalTitle.style.fontSize = '18px';
        modalTitle.style.marginTop = '0';
        modalTitle.style.marginBottom = '16px';
        
        // Текст модального окна
        const modalText = document.createElement('p');
        modalText.textContent = 'Вы уверены, что хотите подать заявку на вступление в этот клуб?';
        modalText.style.fontFamily = 'Inter, sans-serif';
        modalText.style.fontSize = '16px';
        modalText.style.lineHeight = '1.5';
        modalText.style.marginBottom = '24px';
        
        // Кнопки в футере модального окна
        const modalFooter = document.createElement('div');
        modalFooter.style.display = 'flex';
        modalFooter.style.justifyContent = 'flex-end';
        modalFooter.style.gap = '16px';
        
        // Кнопка "Отмена"
        const cancelButton = document.createElement('button');
        cancelButton.textContent = 'Отмена';
        cancelButton.className = 'button button-secondary';
        cancelButton.style.padding = '10px 20px';
        cancelButton.style.border = 'none';
        cancelButton.style.borderRadius = '24px';
        cancelButton.style.fontFamily = 'Inter, sans-serif';
        cancelButton.style.fontWeight = '700';
        cancelButton.style.fontSize = '14px';
        cancelButton.style.backgroundColor = '#F0F0F4';
        cancelButton.style.color = '#111118';
        cancelButton.style.cursor = 'pointer';
        
        // Кнопка "Подтвердить"
        const confirmButton = document.createElement('button');
        confirmButton.textContent = 'Подтвердить';
        confirmButton.className = 'button button-primary';
        confirmButton.style.padding = '10px 20px';
        confirmButton.style.border = 'none';
        confirmButton.style.borderRadius = '24px';
        confirmButton.style.fontFamily = 'Inter, sans-serif';
        confirmButton.style.fontWeight = '700';
        confirmButton.style.fontSize = '14px';
        confirmButton.style.backgroundColor = '#3B19E6';
        confirmButton.style.color = '#FFFFFF';
        confirmButton.style.cursor = 'pointer';
        
        // Добавляем элементы в модальное окно
        modalFooter.appendChild(cancelButton);
        modalFooter.appendChild(confirmButton);
        modalContent.appendChild(modalTitle);
        modalContent.appendChild(modalText);
        modalContent.appendChild(modalFooter);
        modal.appendChild(modalContent);
        
        // Добавляем модальное окно на страницу
        document.body.appendChild(modal);
        
        // Обработчик для кнопки "Отмена"
        cancelButton.addEventListener('click', function() {
            modal.remove();
        });
        
        // Обработчик для кнопки "Подтвердить"
        confirmButton.addEventListener('click', function() {
            // Имитируем отправку заявки на вступление
            showLoading();
            
            // Имитация задержки API-запроса
            setTimeout(() => {
                hideLoading();
                modal.remove();
                
                // Меняем вид карточки клуба после успешной подачи заявки
                const clubActions = clubCard.querySelector('.club-actions');
                clubActions.innerHTML = '';
                
                // Создаем элемент статуса
                const statusDiv = document.createElement('div');
                statusDiv.className = 'club-status';
                
                const statusSpan = document.createElement('span');
                statusSpan.className = 'status pending';
                statusSpan.textContent = 'Ожидает рассмотрения';
                
                statusDiv.appendChild(statusSpan);
                clubActions.parentNode.replaceChild(statusDiv, clubActions);
                
                // Показываем уведомление об успешной подаче заявки
                showNotification(`Заявка на вступление в клуб "${clubTitle}" отправлена`);
            }, 1000);
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
        const loadingOverlay = document.createElement('div');
        loadingOverlay.className = 'loading-overlay';
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
        
        // Добавляем стили анимации
        const style = document.createElement('style');
        style.textContent = `
            @keyframes spin {
                0% { transform: rotate(0deg); }
                100% { transform: rotate(360deg); }
            }
        `;
        document.head.appendChild(style);
        
        loadingOverlay.appendChild(spinner);
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