<template>
  <div>
    <section class="recipes">
      <div class="top">
        <h2>Наши рецепты</h2>
        <div class="search-container">
          <div id="search-bar" class="search-bar" :class="{ active: isSearchVisible }">
            <input
              v-model="searchQuery"
              type="text"
              placeholder="Введите запрос..."
              class="search-input"
              @keyup.enter="performSearch"
            />
            <button class="submit-btn" @click="performSearch">Найти</button>
          </div>
          <button id="search-btn" class="search-btn" @click="toggleSearchBar">Поиск</button>
          <button id="filter-btn" class="filter-btn" @click="$emit('toggle-sidebar')">Фильтры</button>
        </div>
      </div>

      <!-- Сетка рецептов -->
      <div class="recipes-grid">
        <div v-for="recipe in filteredRecipes" :key="recipe.id" class="recipe-card">
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
  display: flex !important;
  justify-content: space-between !important;
  align-items: center !important;
  margin-bottom: 20px !important;
  flex-wrap: wrap !important;
}

.search-container {
  display: flex;
  align-items: center;
  gap: 10px;
}

.search-bar {
  display: none;
}
.search-bar.active {
  display: flex;
  gap: 5px;
}
.search-input {
  width: 200px;
  padding: 5px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.search-btn,
.filter-btn {
  padding: 8px 12px;
  border: none;
  background-color: #12a370;
  color: white;
  border-radius: 4px;
  cursor: pointer;
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
