<template>
  <section class="popular-recipes">
    <div class="top">
      <h2 class="animated-title">Сейчас популярно</h2>
      <router-link to="/recipes" class="view-all">Посмотреть все →</router-link>
    </div>
    <div class="recipes-grid">
      <div 
        class="recipe-card" 
        v-for="(recipe, index) in recipes" 
        :key="index"
        :style="cardAnimationStyle(index)"
      >
        <img :src="recipe.image" :alt="recipe.title" />
        <h3>{{ recipe.title }}</h3>
        <div class="stars">
          <h5>{{ recipe.rating }}</h5>
          <span>{{ getStars(recipe.rating) }}</span>
        </div>
        <router-link :to="'/recipe/' + recipe.id" class="btn-small">Посмотреть</router-link>
      </div>
    </div>
  </section>
</template>

<script>
import lasagna from "@/assets/images/lasagna.png";
import beef from "@/assets/images/beef.png";
import sandwich from "@/assets/images/sandwich.png";

export default {
  data() {
    return {
      recipes: [
        {
          id: 1,
          title: "Лазанья из свиной шейки",
          rating: 4.5,
          image: lasagna,
        },
        {
          id: 2,
          title: "Тушёная говядина",
          rating: 4.2,
          image: beef,
        },
        {
          id: 3,
          title: "Сэндвич с яйцом пашот",
          rating: 3.8,
          image: sandwich,
        },
      ],
    };
  },
  methods: {
    getStars(rating) {
      const fullStars = Math.floor(rating);
      const halfStar = rating % 1 >= 0.5 ? "★" : "";
      const emptyStars = "☆".repeat(5 - fullStars - (halfStar ? 1 : 0));
      return "★".repeat(fullStars) + halfStar + emptyStars;
    },
    cardAnimationStyle(index) {
      return {
        animation: `cardSlideUp 0.5s ease-out forwards`,
        "animation-delay": `${index * 0.1}s`,
      };
    },
  },
};
</script>

<style scoped>
.popular-recipes {
  padding: 50px;
  background-color: white;
}

.top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap; /* Позволяет переносить элементы на следующую строку */
}

.top .view-all {
  margin-top: 10px; /* Отступ для переноса на новую строку */
}

.animated-title {
  font-size: 2em;
  font-weight: bold;
  margin: 0;
  animation: slideFromLeft 0.8s ease-out;
}

.view-all {
  color: #12a370;
  text-decoration: none;
  font-weight: bold;
  animation: slideFromRight 0.8s ease-out;
}

.recipes-grid {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  flex-wrap: wrap; /* Позволяет карточкам переноситься на следующую строку */
  margin-top: 30px;
  animation: cardSlideUp 0.8s ease-out;
}

.recipe-card {
  background-color: white;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  padding: 20px;
  text-align: center;
  transform: translateY(50px);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  flex: 1 1 calc(33.333% - 20px); /* Три карточки в строке */
  max-width: calc(33.333% - 20px);
}

@media (max-width: 1024px) {
  .recipe-card {
    flex: 1 1 calc(50% - 20px); /* Две карточки в строке */
    max-width: calc(50% - 20px);
  }
}

@media (max-width: 768px) {
  .recipe-card {
    flex: 1 1 100%; /* Одна карточка в строке */
    max-width: 100%;
  }

  .btn-small {
    padding: 10px 12px; /* Пропорционально уменьшаем внутренние отступы */
    font-size: 0.9em;
    width: 90%; /* Увеличиваем ширину кнопки на маленьких экранах */
  }
}

.recipe-card:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
}

.recipe-card img {
  width: 100%;
  border-radius: 10px;
  margin-bottom: 15px;
}

.recipe-card h3 {
  font-size: 1.2em;
  font-weight: bold;
  margin: 10px 0;
}

.stars {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin: 10px 0;
}

.stars h5 {
  font-size: 1em;
  margin: 0;
  line-height: 1;
}

.stars span {
  font-size: 1.2em;
  color: #ffa500;
  line-height: 1;
  margin-top: 7%;
}

.btn-small {
  display: inline-block;
  padding: 8px 15px;
  background-color: #12a370;
  color: white;
  border-radius: 25px;
  text-decoration: none;
  margin: 15px auto 0; /* Центрируем кнопку внутри карточки */
  font-weight: bold;
  transition: background-color 0.3s ease;
  width: 80%; /* Адаптивная ширина кнопки относительно карточки */
  max-width: 150px; /* Максимальная ширина кнопки */
  text-align: center; /* Центрируем текст внутри кнопки */
  box-sizing: border-box; /* Учитываем padding в ширине */
}

.btn-small:hover {
  background-color: #0e8a5d;
}

/* Анимация */
@keyframes slideFromLeft {
  from {
    transform: translateX(-100px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

@keyframes slideFromRight {
  from {
    transform: translateX(100px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

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
</style>
