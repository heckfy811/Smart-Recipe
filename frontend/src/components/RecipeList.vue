<template>
  <div>
    <section class="recipes">
      <div class="top">
        <h2>Наши рецепты</h2>
        <div class="hero-buttons">
          <div class="search-container">
            <div id="search-bar" class="search-bar" :class="{ active: isSearchVisible }">
              <input
                v-model="searchQuery"
                type="text"
                placeholder="Введите запрос..."
                class="search-input"
                @keyup.enter="performSearch"
              />
              <button class="search-btn" @click="performSearch">Найти</button>
            </div>
            <button id="search-btn" class="search-btn" @click="toggleSearchBar">Поиск</button>
            <button id="filter-btn" class="filter-btn" @click="$emit('toggle-sidebar')">Фильтры</button>
          </div>
        </div>
      </div>

      <!-- Сетка рецептов -->
      <div class="recipes-grid">
        <div v-for="(recipe, index) in filteredRecipes" :key="recipe.id" class="recipe-card">
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
  </div>
</template>

<script>
import recipesData from "@/assets/testing-data/testRecipeList.json";

export default {
  name: "RecipeList",
  data() {
    return {
      isSearchVisible: false,
      searchQuery: "",
      recipes: recipesData,
      favourites: JSON.parse(localStorage.getItem("favourites")) || [],
    };
  },
  computed: {
    filteredRecipes() {
      const query = this.searchQuery.trim().toLowerCase();
      return query
        ? this.recipes.filter((recipe) =>
            recipe.title.toLowerCase().includes(query)
          )
        : this.recipes;
    },
  },
  methods: {
    toggleSearchBar() {
      this.isSearchVisible = !this.isSearchVisible;
    },
    performSearch() {
      console.log("Поиск по запросу:", this.searchQuery);
    },
    getStars(rating) {
      const fullStars = Math.floor(rating);
      const halfStar = rating % 1 >= 0.5 ? "★" : "";
      const emptyStars = "☆".repeat(5 - fullStars - (halfStar ? 1 : 0));
      return "★".repeat(fullStars) + halfStar + emptyStars;
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
/* Общие стили */
.recipes {
  padding: 50px;
  background-color: white;
}

.top {
  display: flex; /* Используем Flexbox */
  justify-content: space-between; /* Разделяем заголовок и кнопки по сторонам */
  align-items: center; /* Выравниваем элементы по вертикали */
  margin-bottom: 20px;
}

.top h2{
  margin: 0;
}

.hero-buttons{
  margin-bottom: 5px;
}

/* Общие стили */
.search-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

/* Скрытый поисковый блок */
.search-bar {
  display: none;
  align-items: center;
  gap: 10px;
}
.search-bar.active {
  display: flex; /* Показать при активном состоянии */
}

.search-input {
  padding: 12px 20px;
  border: 2px solid #dbdada;
  border-radius: 25px;
  outline: none;
  width: 300px;
  margin-bottom: 5px;
}

.hero-buttons .search-btn,
.hero-buttons .filter-btn {
  padding: 12px 20px;
  border: none;
  border-radius: 25px;
  text-decoration: none;
  font-size: 1em;
  font-weight: bold;
  cursor: pointer;
  margin-right: 15px;
}

.hero-buttons .search-btn {
  background-color: #12a370;
  color: #fff;
  transition: background-color 0.3s ease, transform 0.3s ease;
}

.hero-buttons .search-btn:hover {
  background-color: #0e8a5d;
  transform: scale(1.05);
}

.hero-buttons .filter-btn {
  background-color: transparent;
  border: 2px solid #12a370;
  color: #12a370;
  transition: color 0.3s ease, background-color 0.3s ease, transform 0.3s ease;
}

.hero-buttons .filter-btn:hover {
  background-color: #12a370;
  color: #fff;
  transform: scale(1.05);
}

.recipes-grid {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  gap: 20px;
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

.recipe-card:hover {
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
}

.recipe-card img {
  width: 100%;
  height: 100%;
  border-radius: 10px;
  margin-bottom: 15px;
}

.recipe-card h3 {
  font-size: 1.2em;
  font-weight: bold;
  margin: 10px 0;
}

.recipe-card p {
  font-size: 2em;
  color: #555;
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

.card-actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 15px;
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
  stroke: #333;
  stroke-width: 2;
  transition: fill 0.3s ease, stroke 0.3s ease;
}

.heart-icon.filled {
  fill: #e63946;
  stroke: #e63946;
}

/* Адаптивность */
@media (max-width: 1024px) {
  .recipe-card {
    flex: 1 1 calc(50% - 20px); /* Две карточки в строке */
    max-width: calc(50% - 20px);
  }
}

@media (max-width: 768px) {
  .top {
    flex-direction: column;
    align-items: flex-start;
  }
  .search-container {
    flex-direction: column;
    width: 100%;
  }
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

/*Animacii*/
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
