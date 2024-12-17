<template>
  <section class="popular-recipes">
    <div class="top">
      <h2 class="animated-title">Сейчас популярно</h2>
      <router-link to="/recipes" class="view-all">Посмотреть все →</router-link>
    </div>
    <div class="recipes-grid">
      <div 
        class="recipe-card" 
        v-for="(recipe, index) in sortedRecipes" 
        :key="recipe.id"
        :style="cardAnimationStyle(index)"
      >
        <img :src="recipe.image" :alt="recipe.title" />
        <h3>{{ recipe.title }}</h3>
        <div class="stars">
          <h5>{{ recipe.rating }}</h5>
          <span>{{ getStars(recipe.rating) }}</span>
        </div>
        <div class="card-actions">
          <router-link :to="'/recipe/' + recipe.id" class="btn-small">Посмотреть</router-link>
          <button class="favourite-btn" @click="toggleFavourite(recipe.id)">
            <svg
              class="heart-icon"
              :class="{ filled: isFavourite(recipe.id) }"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M12 21.35l-1.45-1.32C5.4 15.36 2 12.28 2 8.5 2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.78-3.4 6.86-8.55 11.54L12 21.35z"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import popularRecipesData from "@/assets/testing-data/testPopularRecipe.json";

export default {
  data() {
    return {
      recipes: popularRecipesData,
      favourites: JSON.parse(localStorage.getItem("favourites")) || [],
    };
  },
  computed: {
    sortedRecipes() {
      return this.recipes.sort((a, b) => b.popularity - a.popularity);
    },
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
    toggleFavourite(recipeId) {
    if (this.isFavourite(recipeId)) {
      this.favourites = this.favourites.filter((id) => id !== recipeId);
    } else {
      this.favourites.push(recipeId);
    }
    localStorage.setItem("favourites", JSON.stringify(this.favourites));
  },
  isFavourite(recipeId) {
    return this.favourites.includes(recipeId);
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
  display: flex !important;
  justify-content: space-between !important;
  align-items: center !important;
  margin-bottom: 20px !important;
  flex-wrap: wrap !important;
}

.animated-title {
  padding-left: 0 !important;
  font-size: 2em !important;
  font-weight: bold !important;
  margin: 0 !important;
  animation: slideFromLeft 0.8s ease-out;
}

.view-all {
  color: #12a370;
  text-decoration: none;
  font-weight: bold;
  margin: 0 !important;
  margin-bottom: 5.5px !important;
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

.card-actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 15px;
}

.favourite-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 30px;
  width: auto;
  background: none;
  border: none;
  cursor: pointer;
  margin-top: 20px;
  padding: 0;
}

.heart-icon {
  width: 100%;
  height: 100%;
  fill: none;
  stroke: #333; /* Цвет контура сердечка */
  stroke-width: 2;
  transition: fill 0.3s ease, stroke 0.3s ease;
}

.heart-icon.filled {
  fill: #e63946; /* Заполненный цвет сердечка */
  stroke: #e63946; /* Цвет контура совпадает с заливкой */
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
  .view-all {
    margin-top: 10px;
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
  margin-bottom: 10px;
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
