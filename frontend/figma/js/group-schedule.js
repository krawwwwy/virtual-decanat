// Данные расписаний для двух групп (фиктивные)
const scheduleData = {
  'ИДБ-22-11': {
    'Monday': [
      {
        title: 'Дискретная математика',
        time: '9:00 AM - 10:30 AM',
        icon: '../../assets/images/subjects/math-icon.svg'
      },
      {
        title: 'Начертательная геометрия',
        time: '11:00 AM - 12:30 PM',
        icon: '../../assets/images/subjects/geometry-icon.svg'
      },
      {
        title: 'Машинное обучение',
        time: '1:00 PM - 2:30 PM',
        icon: '../../assets/images/subjects/ml-icon.svg'
      },
      {
        title: 'Лабораторная работа Физика',
        time: '3:00 PM - 4:30 PM',
        icon: '../../assets/images/subjects/physics-icon.svg'
      }
    ],
    'Tuesday': [
      {
        title: 'Проектирование информационных систем',
        time: '9:00 AM - 10:30 AM',
        icon: '../../assets/images/subjects/systems-icon.svg'
      },
      {
        title: 'Основы Web-разработки',
        time: '11:00 AM - 12:30 PM',
        icon: '../../assets/images/subjects/web-icon.svg'
      },
      {
        title: 'Управление интеллектуальными активами',
        time: '1:00 PM - 2:30 PM',
        icon: '../../assets/images/subjects/assets-icon.svg'
      },
      {
        title: 'Информационные системы и технологии',
        time: '3:00 PM - 4:30 PM',
        icon: '../../assets/images/subjects/it-icon.svg'
      }
    ]
  },
  'ИДБ-22-10': {
    'Monday': [
      {
        title: 'Информационные системы и технологии',
        time: '9:00 AM - 10:30 AM',
        icon: '../../assets/images/subjects/it-icon.svg'
      },
      {
        title: 'Машинное обучение',
        time: '11:00 AM - 12:30 PM',
        icon: '../../assets/images/subjects/ml-icon.svg'
      },
      {
        title: 'Лабораторная работа Физика',
        time: '1:00 PM - 2:30 PM',
        icon: '../../assets/images/subjects/physics-icon.svg'
      },
      {
        title: 'Проектирование информационных систем',
        time: '3:00 PM - 4:30 PM',
        icon: '../../assets/images/subjects/systems-icon.svg'
      }
    ],
    'Tuesday': [
      {
        title: 'Дискретная математика',
        time: '9:00 AM - 10:30 AM',
        icon: '../../assets/images/subjects/math-icon.svg'
      },
      {
        title: 'Управление интеллектуальными активами',
        time: '11:00 AM - 12:30 PM',
        icon: '../../assets/images/subjects/assets-icon.svg'
      },
      {
        title: 'Основы Web-разработки',
        time: '1:00 PM - 2:30 PM',
        icon: '../../assets/images/subjects/web-icon.svg'
      },
      {
        title: 'Начертательная геометрия',
        time: '3:00 PM - 4:30 PM',
        icon: '../../assets/images/subjects/geometry-icon.svg'
      }
    ]
  }
};

// Текущая выбранная группа и день недели
let currentGroup = 'ИДБ-22-11';
let editMode = false;

// DOM элементы
document.addEventListener('DOMContentLoaded', function() {
  const groupInput = document.getElementById('groupInput');
  const scheduleItems = document.querySelectorAll('.schedule-item');
  const editButtons = document.querySelectorAll('.action-button:not(.today-button):not(.manage-button)');
  const manageButton = document.querySelector('.manage-button');
  const editLinks = document.querySelectorAll('.edit-link');
  const scheduleTitle = document.querySelector('.schedule-title');
  
  // Инициализация значения поля ввода
  groupInput.value = currentGroup;
  
  // Обработчик изменения группы
  groupInput.addEventListener('input', function(e) {
    const input = e.target.value.toUpperCase();
    
    // Автодополнение для двух групп
    if (input === 'И' || input === 'ИД' || input === 'ИДБ' || input === 'ИДБ-' || 
        input === 'ИДБ-2' || input === 'ИДБ-22' || input === 'ИДБ-22-') {
      // Показать подсказки
      showGroupSuggestions(input);
    } else if (input === 'ИДБ-22-11' || input === 'ИДБ-22-10') {
      // Применить выбор группы
      currentGroup = input;
      updateSchedule();
      hideGroupSuggestions();
    }
  });
  
  // Добавим блок подсказок для групп (если его еще нет)
  if (!document.getElementById('groupSuggestions')) {
    const suggestionsDiv = document.createElement('div');
    suggestionsDiv.id = 'groupSuggestions';
    suggestionsDiv.className = 'group-suggestions';
    suggestionsDiv.style.display = 'none';
    suggestionsDiv.style.position = 'absolute';
    suggestionsDiv.style.backgroundColor = 'white';
    suggestionsDiv.style.border = '1px solid #D4D0E7';
    suggestionsDiv.style.borderRadius = '8px';
    suggestionsDiv.style.padding = '8px 0';
    suggestionsDiv.style.zIndex = '100';
    suggestionsDiv.style.width = '576px';
    suggestionsDiv.style.boxShadow = '0 4px 8px rgba(0,0,0,0.1)';
    
    const group1 = document.createElement('div');
    group1.className = 'group-suggestion';
    group1.textContent = 'ИДБ-22-11';
    group1.style.padding = '8px 16px';
    group1.style.cursor = 'pointer';
    group1.style.fontSize = '18px';
    group1.style.fontFamily = 'Inter, sans-serif';
    
    const group2 = document.createElement('div');
    group2.className = 'group-suggestion';
    group2.textContent = 'ИДБ-22-10';
    group2.style.padding = '8px 16px';
    group2.style.cursor = 'pointer';
    group2.style.fontSize = '18px';
    group2.style.fontFamily = 'Inter, sans-serif';
    
    group1.addEventListener('click', function() {
      groupInput.value = 'ИДБ-22-11';
      currentGroup = 'ИДБ-22-11';
      updateSchedule();
      hideGroupSuggestions();
    });
    
    group2.addEventListener('click', function() {
      groupInput.value = 'ИДБ-22-10';
      currentGroup = 'ИДБ-22-10';
      updateSchedule();
      hideGroupSuggestions();
    });
    
    suggestionsDiv.appendChild(group1);
    suggestionsDiv.appendChild(group2);
    
    // Добавляем подсказки после блока с инпутом
    document.querySelector('.group-selection').appendChild(suggestionsDiv);
  }
  
  // Обработчики для кнопок редактирования
  editButtons.forEach(button => {
    button.addEventListener('click', function() {
      if (this.textContent === 'Редактировать') {
        this.textContent = 'Сохранить';
        enableEditMode(this.closest('.content').querySelector('.schedule-list'));
      } else {
        this.textContent = 'Редактировать';
        disableEditMode(this.closest('.content').querySelector('.schedule-list'));
        // Показываем уведомление об успешном сохранении
        showNotification('Расписание успешно сохранено!');
      }
    });
  });
  
  // Обработчик для кнопки управления расписанием
  manageButton.addEventListener('click', function() {
    showNotification('Вы перешли в режим управления расписанием группы ' + currentGroup);
  });
  
  // Обработчики для ссылок редактирования
  editLinks.forEach(link => {
    link.addEventListener('click', function(e) {
      e.preventDefault();
      const scheduleItem = this.closest('.schedule-item');
      const subjectTitle = scheduleItem.querySelector('.subject-title');
      const subjectTime = scheduleItem.querySelector('.subject-time');
      
      // Создаем форму редактирования
      const originalTitle = subjectTitle.textContent;
      const originalTime = subjectTime.textContent;
      
      const editForm = document.createElement('div');
      editForm.className = 'edit-form';
      editForm.style.backgroundColor = '#F0F0F5';
      editForm.style.padding = '15px';
      editForm.style.borderRadius = '8px';
      editForm.style.marginTop = '10px';
      
      editForm.innerHTML = `
        <div style="margin-bottom: 10px;">
          <label style="display: block; margin-bottom: 5px; font-weight: 500;">Название предмета:</label>
          <input type="text" value="${originalTitle}" class="edit-subject-title" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
        </div>
        <div style="margin-bottom: 10px;">
          <label style="display: block; margin-bottom: 5px; font-weight: 500;">Время:</label>
          <input type="text" value="${originalTime}" class="edit-subject-time" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
        </div>
        <div style="display: flex; gap: 10px; justify-content: flex-end;">
          <button class="cancel-edit" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #E7EDF3;">Отмена</button>
          <button class="save-edit" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #3B19E6; color: white;">Сохранить</button>
        </div>
      `;
      
      // Вставляем форму после элемента расписания
      scheduleItem.insertAdjacentElement('afterend', editForm);
      
      // Скрываем элемент расписания
      scheduleItem.style.display = 'none';
      
      // Обработчики для кнопок формы
      editForm.querySelector('.cancel-edit').addEventListener('click', function() {
        scheduleItem.style.display = 'flex';
        editForm.remove();
      });
      
      editForm.querySelector('.save-edit').addEventListener('click', function() {
        const newTitle = editForm.querySelector('.edit-subject-title').value;
        const newTime = editForm.querySelector('.edit-subject-time').value;
        
        subjectTitle.textContent = newTitle;
        subjectTime.textContent = newTime;
        
        scheduleItem.style.display = 'flex';
        editForm.remove();
        
        showNotification('Изменения сохранены!');
      });
    });
  });
  
  // Инициализация расписания
  updateSchedule();
});

// Функция отображения подсказок групп
function showGroupSuggestions(input) {
  const suggestionsDiv = document.getElementById('groupSuggestions');
  if (suggestionsDiv) {
    suggestionsDiv.style.display = 'block';
    
    // Позиционирование
    const groupSelection = document.querySelector('.group-selection');
    const rect = groupSelection.getBoundingClientRect();
    suggestionsDiv.style.top = (groupSelection.offsetHeight) + 'px';
    
    // Подсветка подходящих групп
    const suggestions = suggestionsDiv.querySelectorAll('.group-suggestion');
    suggestions.forEach(suggestion => {
      if (suggestion.textContent.startsWith(input)) {
        suggestion.style.backgroundColor = '#F0F0F5';
      } else {
        suggestion.style.backgroundColor = '';
      }
    });
  }
}

// Функция скрытия подсказок групп
function hideGroupSuggestions() {
  const suggestionsDiv = document.getElementById('groupSuggestions');
  if (suggestionsDiv) {
    suggestionsDiv.style.display = 'none';
  }
}

// Функция обновления расписания
function updateSchedule() {
  const scheduleTitle = document.querySelector('.schedule-title');
  scheduleTitle.textContent = `Расписание учебной группы ${currentGroup}`;
  
  // Обновляем расписание для понедельника
  updateDaySchedule('Monday', 0);
  
  // Обновляем расписание для вторника
  updateDaySchedule('Tuesday', 1);
}

// Обновление расписания для конкретного дня
function updateDaySchedule(day, dayIndex) {
  const scheduleList = document.querySelectorAll('.schedule-list')[dayIndex];
  scheduleList.innerHTML = '';
  
  const dayData = scheduleData[currentGroup][day];
  dayData.forEach(subject => {
    const scheduleItem = document.createElement('div');
    scheduleItem.className = 'schedule-item';
    
    scheduleItem.innerHTML = `
      <div class="schedule-item-left">
        <div class="subject-icon">
          <img src="${subject.icon}" alt="${subject.title}">
        </div>
        <div class="subject-info">
          <h3 class="subject-title">${subject.title}</h3>
          <span class="subject-time">${subject.time}</span>
        </div>
      </div>
      <div class="schedule-item-actions" style="display: flex; gap: 10px;">
        <a href="#" class="edit-link">редактировать</a>
        <a href="#" class="delete-link" style="color: #ff3b30;">удалить</a>
      </div>
    `;
    
    scheduleList.appendChild(scheduleItem);
    
    // Добавляем обработчик для новой ссылки редактирования
    const editLink = scheduleItem.querySelector('.edit-link');
    editLink.addEventListener('click', function(e) {
      e.preventDefault();
      const subjectTitle = scheduleItem.querySelector('.subject-title');
      const subjectTime = scheduleItem.querySelector('.subject-time');
      
      // Создаем форму редактирования
      const originalTitle = subjectTitle.textContent;
      const originalTime = subjectTime.textContent;
      
      const editForm = document.createElement('div');
      editForm.className = 'edit-form';
      editForm.style.backgroundColor = '#F0F0F5';
      editForm.style.padding = '15px';
      editForm.style.borderRadius = '8px';
      editForm.style.marginTop = '10px';
      
      editForm.innerHTML = `
        <div style="margin-bottom: 10px;">
          <label style="display: block; margin-bottom: 5px; font-weight: 500;">Название предмета:</label>
          <input type="text" value="${originalTitle}" class="edit-subject-title" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
        </div>
        <div style="margin-bottom: 10px;">
          <label style="display: block; margin-bottom: 5px; font-weight: 500;">Время:</label>
          <input type="text" value="${originalTime}" class="edit-subject-time" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
        </div>
        <div style="display: flex; gap: 10px; justify-content: flex-end;">
          <button class="cancel-edit" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #E7EDF3;">Отмена</button>
          <button class="save-edit" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #3B19E6; color: white;">Сохранить</button>
        </div>
      `;
      
      // Вставляем форму после элемента расписания
      scheduleItem.insertAdjacentElement('afterend', editForm);
      
      // Скрываем элемент расписания
      scheduleItem.style.display = 'none';
      
      // Обработчики для кнопок формы
      editForm.querySelector('.cancel-edit').addEventListener('click', function() {
        scheduleItem.style.display = 'flex';
        editForm.remove();
      });
      
      editForm.querySelector('.save-edit').addEventListener('click', function() {
        const newTitle = editForm.querySelector('.edit-subject-title').value;
        const newTime = editForm.querySelector('.edit-subject-time').value;
        
        subjectTitle.textContent = newTitle;
        subjectTime.textContent = newTime;
        
        scheduleItem.style.display = 'flex';
        editForm.remove();
        
        showNotification('Изменения сохранены!');
      });
    });
    
    // Добавляем обработчик для удаления предмета
    const deleteLink = scheduleItem.querySelector('.delete-link');
    deleteLink.addEventListener('click', function(e) {
      e.preventDefault();
      if (confirm('Вы уверены, что хотите удалить этот предмет из расписания?')) {
        scheduleItem.style.height = scheduleItem.offsetHeight + 'px';
        scheduleItem.style.transition = 'all 0.3s ease-in-out';
        setTimeout(() => {
          scheduleItem.style.height = '0px';
          scheduleItem.style.opacity = '0';
          scheduleItem.style.padding = '0';
          scheduleItem.style.margin = '0';
          scheduleItem.style.overflow = 'hidden';
        }, 10);
        
        setTimeout(() => {
          scheduleItem.remove();
          showNotification('Предмет удален из расписания');
          
          // Удаляем из данных
          const index = dayData.findIndex(item => 
            item.title === scheduleItem.querySelector('.subject-title').textContent &&
            item.time === scheduleItem.querySelector('.subject-time').textContent
          );
          
          if (index !== -1) {
            dayData.splice(index, 1);
          }
        }, 300);
      }
    });
  });
  
  // Добавляем кнопку "Добавить предмет" в конец списка
  const addButton = document.createElement('div');
  addButton.className = 'add-subject-button';
  addButton.style.padding = '15px';
  addButton.style.marginTop = '10px';
  addButton.style.backgroundColor = '#F0F0F5';
  addButton.style.borderRadius = '8px';
  addButton.style.textAlign = 'center';
  addButton.style.cursor = 'pointer';
  addButton.style.display = 'flex';
  addButton.style.alignItems = 'center';
  addButton.style.justifyContent = 'center';
  addButton.innerHTML = `
    <div style="width: 24px; height: 24px; background-color: #3B19E6; border-radius: 50%; color: white; display: flex; align-items: center; justify-content: center; margin-right: 10px; font-weight: bold;">+</div>
    <span style="font-family: 'Inter', sans-serif; font-weight: 500; font-size: 18px;">Добавить предмет</span>
  `;
  
  scheduleList.appendChild(addButton);
  
  // Обработчик для кнопки добавления
  addButton.addEventListener('click', function() {
    // Создаем форму для добавления нового предмета
    const addForm = document.createElement('div');
    addForm.className = 'add-form';
    addForm.style.backgroundColor = '#F0F0F5';
    addForm.style.padding = '15px';
    addForm.style.borderRadius = '8px';
    addForm.style.marginTop = '10px';
    
    // Иконки предметов для выбора
    const icons = [
      { name: 'Математика', value: '../../assets/images/subjects/math-icon.svg' },
      { name: 'Геометрия', value: '../../assets/images/subjects/geometry-icon.svg' },
      { name: 'Машинное обучение', value: '../../assets/images/subjects/ml-icon.svg' },
      { name: 'Физика', value: '../../assets/images/subjects/physics-icon.svg' },
      { name: 'Информационные системы', value: '../../assets/images/subjects/systems-icon.svg' },
      { name: 'Web-разработка', value: '../../assets/images/subjects/web-icon.svg' },
      { name: 'Интеллектуальные активы', value: '../../assets/images/subjects/assets-icon.svg' },
      { name: 'ИТ', value: '../../assets/images/subjects/it-icon.svg' }
    ];
    
    // Генерируем опции для выбора иконки
    const iconOptions = icons.map(icon => 
      `<option value="${icon.value}">${icon.name}</option>`
    ).join('');
    
    // Форма для добавления предмета
    addForm.innerHTML = `
      <h3 style="margin-top: 0; font-family: 'Inter', sans-serif; font-weight: 700; font-size: 20px;">Добавить новый предмет</h3>
      <div style="margin-bottom: 10px;">
        <label style="display: block; margin-bottom: 5px; font-weight: 500;">Название предмета:</label>
        <input type="text" class="new-subject-title" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
      </div>
      <div style="margin-bottom: 10px;">
        <label style="display: block; margin-bottom: 5px; font-weight: 500;">Время:</label>
        <input type="text" class="new-subject-time" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
      </div>
      <div style="margin-bottom: 15px;">
        <label style="display: block; margin-bottom: 5px; font-weight: 500;">Иконка предмета:</label>
        <select class="new-subject-icon" style="width: 100%; padding: 8px; border: 1px solid #D4D0E7; border-radius: 4px;">
          ${iconOptions}
        </select>
      </div>
      <div style="display: flex; gap: 10px; justify-content: flex-end;">
        <button class="cancel-add" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #E7EDF3;">Отмена</button>
        <button class="save-add" style="padding: 8px 16px; border: none; border-radius: 4px; cursor: pointer; background-color: #3B19E6; color: white;">Добавить</button>
      </div>
    `;
    
    // Вставляем форму после кнопки добавления
    addButton.insertAdjacentElement('afterend', addForm);
    
    // Скрываем кнопку добавления
    addButton.style.display = 'none';
    
    // Обработчики для кнопок формы
    addForm.querySelector('.cancel-add').addEventListener('click', function() {
      addButton.style.display = 'flex';
      addForm.remove();
    });
    
    addForm.querySelector('.save-add').addEventListener('click', function() {
      const newTitle = addForm.querySelector('.new-subject-title').value;
      const newTime = addForm.querySelector('.new-subject-time').value;
      const newIcon = addForm.querySelector('.new-subject-icon').value;
      
      if (newTitle && newTime) {
        // Добавляем новый предмет в данные
        dayData.push({
          title: newTitle,
          time: newTime,
          icon: newIcon || '../../assets/images/subjects/it-icon.svg' // Значение по умолчанию
        });
        
        // Перерисовываем расписание
        updateDaySchedule(day, dayIndex);
        
        showNotification('Новый предмет добавлен в расписание');
      } else {
        alert('Пожалуйста, заполните все поля!');
      }
    });
  });
}

// Включение режима редактирования
function enableEditMode(scheduleList) {
  editMode = true;
  const items = scheduleList.querySelectorAll('.schedule-item');
  
  items.forEach(item => {
    item.style.backgroundColor = '#F9F8FC';
    item.style.border = '1px dashed #3B19E6';
    item.style.cursor = 'pointer';
  });
}

// Выключение режима редактирования
function disableEditMode(scheduleList) {
  editMode = false;
  const items = scheduleList.querySelectorAll('.schedule-item');
  
  items.forEach(item => {
    item.style.backgroundColor = '#FFFFFF';
    item.style.border = 'none';
    item.style.cursor = 'default';
  });
}

// Функция для отображения уведомлений
function showNotification(message) {
  // Создаем элемент для уведомления, если его еще нет
  if (!document.getElementById('notification')) {
    const notification = document.createElement('div');
    notification.id = 'notification';
    notification.style.position = 'fixed';
    notification.style.top = '20px';
    notification.style.right = '20px';
    notification.style.padding = '15px 20px';
    notification.style.backgroundColor = '#3B19E6';
    notification.style.color = 'white';
    notification.style.borderRadius = '8px';
    notification.style.boxShadow = '0 4px 8px rgba(0,0,0,0.2)';
    notification.style.opacity = '0';
    notification.style.transition = 'opacity 0.3s ease-in-out';
    notification.style.zIndex = '1000';
    notification.style.fontFamily = 'Lexend, sans-serif';
    document.body.appendChild(notification);
  }
  
  const notification = document.getElementById('notification');
  notification.textContent = message;
  notification.style.opacity = '1';
  
  // Скрываем уведомление через 3 секунды
  setTimeout(() => {
    notification.style.opacity = '0';
  }, 3000);
} 