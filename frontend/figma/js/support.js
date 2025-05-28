/**
 * JavaScript для страницы социальной поддержки
 */
document.addEventListener('DOMContentLoaded', function() {
    // Проверяем авторизацию
    if (!window.authApi.isAuthenticated() || localStorage.getItem('userRole') !== 'student') {
        window.location.href = '../login-page.html';
        return;
    }

    // Поиск по таблице поддержки
    const searchInput = document.querySelector('.search-input');
    if (searchInput) {
        searchInput.addEventListener('input', function() {
            const searchTerm = this.value.toLowerCase().trim();
            filterTable(searchTerm);
        });
    }

    /**
     * Фильтрация таблицы по поисковому запросу
     * @param {string} searchTerm - Строка поиска
     */
    function filterTable(searchTerm) {
        const tableRows = document.querySelectorAll('.support-table tbody tr');
        
        tableRows.forEach(row => {
            const supportType = row.cells[0].textContent.toLowerCase();
            const description = row.cells[1].textContent.toLowerCase();
            const deadline = row.cells[2].textContent.toLowerCase();
            const status = row.cells[3].querySelector('.status').textContent.toLowerCase();
            
            const matchesSearch = 
                supportType.includes(searchTerm) || 
                description.includes(searchTerm) || 
                deadline.includes(searchTerm) || 
                status.includes(searchTerm);
            
            row.style.display = matchesSearch ? '' : 'none';
        });

        // Если таблица пуста, показываем сообщение
        const tableBody = document.querySelector('.support-table tbody');
        let visibleRows = 0;
        
        tableRows.forEach(row => {
            if (row.style.display !== 'none') {
                visibleRows++;
            }
        });
        
        // Проверяем, есть ли у таблицы сообщение о пустом результате
        let emptyMessage = document.querySelector('.empty-search-message');
        
        if (visibleRows === 0) {
            // Если нет видимых строк и нет сообщения, добавляем его
            if (!emptyMessage) {
                emptyMessage = document.createElement('tr');
                emptyMessage.className = 'empty-search-message';
                
                const messageCell = document.createElement('td');
                messageCell.colSpan = 4;
                messageCell.textContent = 'По вашему запросу ничего не найдено';
                messageCell.style.textAlign = 'center';
                messageCell.style.padding = '24px';
                messageCell.style.color = '#636388';
                
                emptyMessage.appendChild(messageCell);
                tableBody.appendChild(emptyMessage);
            }
        } else if (emptyMessage) {
            // Если есть видимые строки, но есть и сообщение, удаляем его
            emptyMessage.remove();
        }
    }

    // Обработка загрузки файлов
    const fileInput = document.getElementById('documents');
    const fileNameSpan = document.querySelector('.file-name');
    const fileButton = document.querySelector('.file-button');
    
    if (fileInput && fileNameSpan && fileButton) {
        // При клике на кнопку загрузки файла вызываем клик по скрытому input
        fileButton.addEventListener('click', function() {
            fileInput.click();
        });
        
        // При выборе файла обновляем отображаемое имя
        fileInput.addEventListener('change', function() {
            if (this.files.length > 0) {
                if (this.files.length === 1) {
                    fileNameSpan.textContent = this.files[0].name;
                } else {
                    fileNameSpan.textContent = `Выбрано файлов: ${this.files.length}`;
                }
            } else {
                fileNameSpan.textContent = 'Выберите документ';
            }
        });
    }

    // Обработка отправки формы
    const supportForm = document.getElementById('supportForm');
    if (supportForm) {
        supportForm.addEventListener('submit', function(e) {
            e.preventDefault();
            
            // Валидация формы
            const supportType = document.getElementById('support-type');
            const reason = document.getElementById('reason');
            
            let valid = true;
            let errorMessage = '';
            
            if (supportType.value === '') {
                valid = false;
                errorMessage = 'Пожалуйста, выберите тип поддержки';
                supportType.focus();
            } else if (reason.value.trim() === '') {
                valid = false;
                errorMessage = 'Пожалуйста, укажите причину подачи заявления';
                reason.focus();
            }
            
            if (!valid) {
                showNotification(errorMessage, 'error');
                return;
            }
            
            // Если форма валидна, имитируем отправку
            showLoading();
            
            // Имитируем задержку запроса к API
            setTimeout(() => {
                hideLoading();
                
                // Создаем новую строку в таблице для заявки
                const newRow = createApplicationRow(
                    supportType.options[supportType.selectedIndex].text,
                    `${supportType.options[supportType.selectedIndex].text} для студентов`,
                    getFormattedDeadline(),
                    'на рассмотрении'
                );
                
                // Добавляем строку в таблицу
                const tableBody = document.querySelector('.support-table tbody');
                tableBody.insertBefore(newRow, tableBody.firstChild);
                
                // Анимируем новую строку
                newRow.classList.add('new-row');
                setTimeout(() => {
                    newRow.classList.remove('new-row');
                }, 3000);
                
                // Сбрасываем форму
                supportForm.reset();
                fileNameSpan.textContent = 'Выберите документ';
                
                showNotification('Заявка успешно отправлена', 'success');
            }, 1500);
        });
    }

    /**
     * Создаем новую строку для таблицы заявок
     */
    function createApplicationRow(type, description, deadline, status) {
        const row = document.createElement('tr');
        
        // Ячейка типа
        const typeCell = document.createElement('td');
        typeCell.textContent = type;
        row.appendChild(typeCell);
        
        // Ячейка описания
        const descCell = document.createElement('td');
        descCell.textContent = description;
        row.appendChild(descCell);
        
        // Ячейка дедлайна
        const deadlineCell = document.createElement('td');
        deadlineCell.textContent = deadline;
        row.appendChild(deadlineCell);
        
        // Ячейка статуса
        const statusCell = document.createElement('td');
        const statusDiv = document.createElement('div');
        statusDiv.className = 'status pending';
        statusDiv.textContent = status;
        statusCell.appendChild(statusDiv);
        row.appendChild(statusCell);
        
        // Добавляем стиль анимации
        row.style.backgroundColor = 'rgba(255, 248, 230, 0.3)';
        
        return row;
    }

    /**
     * Получаем отформатированный дедлайн (через 30 дней)
     */
    function getFormattedDeadline() {
        const date = new Date();
        date.setDate(date.getDate() + 30);
        
        const day = String(date.getDate()).padStart(2, '0');
        const month = String(date.getMonth() + 1).padStart(2, '0');
        const year = date.getFullYear();
        
        return `${day}-${month}-${year}`;
    }

    /**
     * Показать уведомление
     */
    function showNotification(message, type = 'info') {
        // Создаем элемент уведомления
        const notification = document.createElement('div');
        notification.className = 'notification';
        notification.textContent = message;
        
        // Стили для уведомления
        notification.style.position = 'fixed';
        notification.style.top = '20px';
        notification.style.right = '20px';
        notification.style.padding = '16px 24px';
        notification.style.borderRadius = '8px';
        notification.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.15)';
        notification.style.zIndex = '9999';
        notification.style.fontFamily = 'Inter, sans-serif';
        notification.style.fontSize = '16px';
        notification.style.fontWeight = '500';
        
        // Применяем стили в зависимости от типа уведомления
        if (type === 'error') {
            notification.style.backgroundColor = '#FFE6E6';
            notification.style.color = '#D90000';
            notification.style.borderLeft = '4px solid #D90000';
        } else if (type === 'success') {
            notification.style.backgroundColor = '#E6F7ED';
            notification.style.color = '#0D8A3E';
            notification.style.borderLeft = '4px solid #0D8A3E';
        } else {
            notification.style.backgroundColor = '#E6F0FF';
            notification.style.color = '#0057D9';
            notification.style.borderLeft = '4px solid #0057D9';
        }
        
        document.body.appendChild(notification);
        
        // Анимация появления
        notification.style.opacity = '0';
        notification.style.transform = 'translateX(20px)';
        notification.style.transition = 'opacity 0.3s, transform 0.3s';
        
        setTimeout(() => {
            notification.style.opacity = '1';
            notification.style.transform = 'translateX(0)';
        }, 10);
        
        // Удаление уведомления через 5 секунд
        setTimeout(() => {
            notification.style.opacity = '0';
            notification.style.transform = 'translateX(20px)';
            
            setTimeout(() => {
                notification.remove();
            }, 300);
        }, 5000);
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
        spinner.style.width = '64px';
        spinner.style.height = '64px';
        spinner.style.border = '6px solid #f3f3f3';
        spinner.style.borderTop = '6px solid #3B19E6';
        spinner.style.borderRadius = '50%';
        spinner.style.animation = 'spin 1s linear infinite';
        
        // Добавляем стили анимации, если их еще нет
        if (!document.getElementById('loading-spinner-style')) {
            const style = document.createElement('style');
            style.id = 'loading-spinner-style';
            style.textContent = `
                @keyframes spin {
                    0% { transform: rotate(0deg); }
                    100% { transform: rotate(360deg); }
                }
            `;
            document.head.appendChild(style);
        }
        
        loadingOverlay.appendChild(spinner);
        document.body.appendChild(loadingOverlay);
        
        // Анимация появления
        loadingOverlay.style.opacity = '0';
        loadingOverlay.style.transition = 'opacity 0.3s';
        
        setTimeout(() => {
            loadingOverlay.style.opacity = '1';
        }, 10);
    }

    /**
     * Скрыть индикатор загрузки
     */
    function hideLoading() {
        const loadingOverlay = document.querySelector('.loading-overlay');
        if (loadingOverlay) {
            loadingOverlay.style.opacity = '0';
            
            // Анимация исчезновения
            setTimeout(() => {
                loadingOverlay.remove();
            }, 300);
        }
    }
}); 