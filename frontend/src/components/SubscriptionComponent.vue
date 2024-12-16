<template>
  <section class="subscription">
    <div class="subscription-text">
      <p>
        Приобретите подписку для готовых планов питания, рассчитанных для разных рационов
      </p>
      <a href="#" class="subscription-button">Подробнее</a>
    </div>
  </section>
</template>

<script>
export default {
  name: "SubscriptionComponent",
};
</script>

<style scoped>
/* Основной стиль секции */
.subscription {
  margin-top: 150px;
  height: 454px;
  background-image: url("../src/assets/images/background.png");
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;
  overflow: hidden;
  position: relative;
  width: 100%; /* Всегда занимает ширину экрана */
  animation: cardSlideUp 0.8s ease-out;
}

/* Затемнение фона */
.subscription::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5); /* Затемнение фотографии */
  z-index: 1;
}

/* Текстовый блок */
.subscription-text {
  position: relative;
  z-index: 2; /* Над затемнением */
  background-color: rgba(0, 0, 0, 0.3); /* Прозрачный фон для текста */
  border-radius: 25px;
  padding: 50px;
  text-align: center;
  width: 50%; /* 50% от ширины секции */
}

.subscription-text p {
  color: white !important;
  opacity: 0.8;
  font-size: 1.2em;
}

.subscription-button {
  padding: 12px 20px;
  border: 1px solid #12a370;
  border-radius: 25px;
  cursor: pointer;
  text-decoration: none;
  background-color: #12a370;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px auto 0; /* Центрируем кнопку */
  width: fit-content; /* Размер кнопки соответствует её содержимому */
  transition: background-color 0.3s ease, color 0.3s ease;
}

.subscription-button:hover {
  color: #12a370;
  background-color: white;
}

/*Анимация*/
@keyframes cardSlideUp {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

/* Адаптивность */
@media (max-width: 1024px) {
  .subscription {
    height: 400px; /* Уменьшаем высоту для планшетов */
  }

  .subscription-text {
    width: 70%; /* Увеличиваем ширину текста на узких экранах */
    padding: 40px; /* Уменьшаем отступы */
  }

  .subscription-text p {
    font-size: 1em; /* Уменьшаем шрифт */
  }

  .subscription-button {
    font-size: 1em; /* Поддерживаем читаемый размер текста */
    padding: 10px 20px; /* Сохраняем пропорции */
  }
}

@media (max-width: 768px) {
  .subscription {
    height: 300px; /* Ещё меньше для мобильных устройств */
  }

  .subscription-text {
    width: 90%; /* Текст занимает большую часть ширины */
    padding: 30px; /* Ещё меньше отступы */
  }

  .subscription-text p {
    font-size: 0.9em;
  }

  .subscription-button {
    font-size: 0.9em; /* Уменьшаем шрифт */
    padding: 10px 20px; /* Сохраняем пропорции кнопки */
    margin: 15px auto 0; /* Центрируем кнопку */
  }
}
</style>
