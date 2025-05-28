/**
 * JavaScript для функциональности страниц абитуриента
 */
document.addEventListener('DOMContentLoaded', function() {
    // Обработка загрузки файлов на странице подачи заявления
    setupFileUpload();
    
    // Обработка отправки формы заявления
    setupApplicationForm();
    
    // Обработка формы проверки статуса
    setupStatusForm();
});

/**
 * Настройка обработчика загрузки файлов
 */
function setupFileUpload() {
    const fileInput = document.getElementById('documents');
    const fileButton = document.querySelector('.file-button');
    const fileName = document.querySelector('.file-name');
    const fileList = document.getElementById('file-list');
    
    if (!fileInput || !fileButton || !fileName) return;
    
    // Обработка клика на кнопку загрузки файла
    fileButton.addEventListener('click', function() {
        fileInput.click();
    });
    
    // Обработка выбора файла
    fileInput.addEventListener('change', function() {
        if (this.files.length > 0) {
            fileName.textContent = `Выбрано файлов: ${this.files.length}`;
            
            // Очистка предыдущего списка файлов
            if (fileList) {
                fileList.innerHTML = '';
                
                // Добавление каждого файла в список
                Array.from(this.files).forEach(file => {
                    const fileItem = document.createElement('div');
                    fileItem.className = 'file-item';
                    
                    const fileItemName = document.createElement('span');
                    fileItemName.className = 'file-item-name';
                    fileItemName.textContent = file.name;
                    
                    const fileRemove = document.createElement('span');
                    fileRemove.className = 'file-remove';
                    fileRemove.textContent = '✕';
                    fileRemove.addEventListener('click', function() {
                        fileItem.remove();
                        
                        // Пересчитываем количество файлов
                        const remainingFiles = document.querySelectorAll('.file-item').length;
                        fileName.textContent = remainingFiles > 0 ? 
                            `Выбрано файлов: ${remainingFiles}` : 
                            'Загрузить документы';
                    });
                    
                    fileItem.appendChild(fileItemName);
                    fileItem.appendChild(fileRemove);
                    fileList.appendChild(fileItem);
                });
            }
        } else {
            fileName.textContent = 'Загрузить документы';
            if (fileList) {
                fileList.innerHTML = '';
            }
        }
    });
}

/**
 * Настройка обработчика отправки формы заявления
 */
function setupApplicationForm() {
    const applicationForm = document.getElementById('applicationForm');
    
    if (!applicationForm) return;
    
    applicationForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        // Показываем загрузку
        showLoading();
        
        // Имитируем отправку данных (задержка для демонстрации)
        setTimeout(() => {
            hideLoading();
            
            // Сохраняем данные в localStorage для демонстрации
            const formData = new FormData(applicationForm);
            const applicationData = {};
            
            formData.forEach((value, key) => {
                applicationData[key] = value;
            });
            
            localStorage.setItem('applicationData', JSON.stringify(applicationData));
            
            // Показываем уведомление об успешной отправке
            showNotification('Заявление успешно отправлено', 'success');
            
            // Перенаправляем на страницу статуса
            setTimeout(() => {
                window.location.href = 'status.html';
            }, 2000);
        }, 1500);
    });
}

/**
 * Настройка обработчика формы проверки статуса
 */
function setupStatusForm() {
    const statusForm = document.getElementById('statusForm');
    const statusResult = document.getElementById('statusResult');
    
    if (!statusForm || !statusResult) return;
    
    // Заполняем форму данными из localStorage, если они есть
    const applicationData = localStorage.getItem('applicationData');
    
    if (applicationData) {
        const parsedData = JSON.parse(applicationData);
        
        // Заполняем поля формы
        const fullnameInput = document.getElementById('fullname');
        const birthdateInput = document.getElementById('birthdate');
        const passportInput = document.getElementById('passport');
        
        if (fullnameInput && parsedData.fullname) {
            fullnameInput.value = parsedData.fullname;
        }
        
        if (birthdateInput && parsedData.birthdate) {
            birthdateInput.value = parsedData.birthdate;
        }
        
        if (passportInput && parsedData.passport) {
            passportInput.value = parsedData.passport;
        }
    }
    
    statusForm.addEventListener('submit', function(e) {
        e.preventDefault();
        
        // Показываем загрузку
        showLoading();
        
        // Имитируем проверку статуса
        setTimeout(() => {
            hideLoading();
            
            // Показываем результат
            statusResult.classList.remove('hidden');
            
            // Прокручиваем страницу к результату
            statusResult.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            });
        }, 1200);
    });
}

/**
 * Показать уведомление
 */
function showNotification(message, type = 'info') {
    // Если уже есть уведомление, удаляем его
    const existingNotification = document.querySelector('.notification');
    if (existingNotification) {
        existingNotification.remove();
    }
    
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
    
    // Блокировка прокрутки страницы
    document.body.style.overflow = 'hidden';
    
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
        
        // Возвращаем прокрутку страницы
        document.body.style.overflow = '';
        
        // Анимация исчезновения
        setTimeout(() => {
            loadingOverlay.remove();
        }, 300);
    }
} 