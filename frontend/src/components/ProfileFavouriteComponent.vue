<template>
    <section class="recipes">
      <!-- Заголовок и кнопка "Посмотреть все" -->
      <div class="top">
        <h2>Избранное</h2>
        <router-link
          v-if="favouriteRecipes.length > 3"
          to="/favorite"
          class="view-all"
        >
          Посмотреть все →
        </router-link>
      </div>
  
      <!-- Сетка карточек -->
      <div v-if="displayedRecipes.length > 0" class="recipes-grid">
        <div
          class="recipe-card"
          v-for="(recipe, index) in displayedRecipes"
          :key="recipe.id"
          :style="cardAnimationStyle(index)"
        >
          <!-- Изображение рецепта -->
          <img :src="recipe.image" :alt="recipe.title" />
  
          <!-- Название рецепта -->
          <h3>{{ recipe.title }}</h3>
  
          <!-- Рейтинг -->
          <div v-if="recipe.rating" class="stars">
            <h5>{{ recipe.rating }}</h5>
            <span>{{ getStars(recipe.rating) }}</span>
          </div>
  
          <!-- Кнопка "Посмотреть" и иконка сердечка -->
          <div class="card-actions">
            <router-link :to="'/recipe/' + recipe.id" class="btn-small">Посмотреть</router-link>
            <button class="favourite-btn" @click="removeFromFavourites(recipe.id)">
              <svg
                class="heart-icon filled"
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
  
      <!-- Сообщение, если нет избранных рецептов -->
      <p v-else class="empty-message">У вас нет избранных рецептов.</p>
    </section>
  </template>  
   
  
  <script>
import recipesData from "@/assets/testing-data/testRecipeList.json";

export default {
  name: "ProfileFavouriteComponent",
  data() {
    return {
      recipes: recipesData,
      favouriteIds: JSON.parse(localStorage.getItem("favourites")) || [],
    };
  },
  computed: {
    // Получаем избранные рецепты
    favouriteRecipes() {
      return this.recipes.filter((recipe) =>
        this.favouriteIds.includes(recipe.id)
      );
    },
    // Ограничиваем отображение до 3 рецептов
    displayedRecipes() {
      return this.favouriteRecipes.slice(0, 3);
    },
  },
  methods: {
    getStars(rating) {
      const fullStars = Math.floor(rating);
      const halfStar = rating % 1 >= 0.5 ? "★" : "";
      const emptyStars = "☆".repeat(5 - fullStars - (halfStar ? 1 : 0));
      return "★".repeat(fullStars) + halfStar + emptyStars;
    },
    removeFromFavourites(recipeId) {
      this.favouriteIds = this.favouriteIds.filter((id) => id !== recipeId);
      localStorage.setItem("favourites", JSON.stringify(this.favouriteIds));
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
.recipes {
  padding: 50px;
  background-color: white;
}

.top {
  display: flex;
  justify-content: space-between; /* Текст слева, кнопка справа */
  align-items: center; /* Вертикальное выравнивание по центру */
  margin-bottom: 20px;
}

.top h2 {
  margin: 0; /* Убираем стандартные отступы у заголовка */
  padding: 0;
  font-size: 2em;
}

.view-all {
  text-decoration: none;
  color: #12a370;
  font-weight: bold;
  font-size: 1em;
  margin: 0; /* Убираем все отступы */
  padding-right: 10px;
  transition: color 0.3s ease;
}

.view-all:hover {
  color: #0e8a5d;
}


.empty-message {
  text-align: center;
  font-size: 1.2em;
  margin-top: 50px;
  color: #555;
}

.recipes-grid {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  flex-wrap: wrap;
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
  flex: 1 1 calc(33.333% - 20px);
  max-width: calc(33.333% - 20px);
  position: relative;
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

.favourite-btn:hover {
  transform: scale(1.1);
}

.heart-icon {
  width: 100%;
  height: 100%;
  fill: #e63946;
  stroke: #e63946;
}

.btn-small {
  display: inline-block;
  padding: 8px 15px;
  background-color: #12a370;
  color: white;
  border-radius: 25px;
  text-decoration: none;
  margin: 15px auto 0;
  font-weight: bold;
  transition: background-color 0.3s ease;
  width: 80%;
  max-width: 150px;
  text-align: center;
  box-sizing: border-box;
}

.btn-small:hover {
  background-color: #0e8a5d;
}

/* Адаптивные стили */
@media (max-width: 1024px) {
  .recipe-card {
    flex: 1 1 calc(50% - 20px);
    max-width: calc(50% - 20px);
  }
}

@media (max-width: 768px) {
  .recipe-card {
    flex: 1 1 100%;
    max-width: 100%;
  }

  .btn-small {
    padding: 10px 12px;
    font-size: 0.9em;
    width: 90%;
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
  