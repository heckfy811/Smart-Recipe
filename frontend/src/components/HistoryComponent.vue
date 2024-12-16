<template>
  <div>
    <section class="recipes">
      <div class="top">
        <h2>Опробованное</h2>
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
          <div class="filter">
            <!-- Кнопка "Фильтры" -->
            <button id="filter-btn" class="filter-btn" @click="$emit('toggle-sidebar')">Фильтры</button>          
          </div>
        </div>
      </div>

      <!-- Сетка рецептов -->
      <div class="recipes-grid">
        <div v-for="recipe in filteredRecipes" :key="recipe.id" class="recipe-card">
          <img :src="recipe.image" :alt="recipe.title" />
          <h3>{{ recipe.title }}</h3>
          <p>{{ recipe.price }}Р</p>
          <div class="stars">
            <h5>{{ recipe.rating }}</h5>
            <span>★★★★☆</span>
          </div>
          <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "HistoryComponent.vue",
  data() {
    return {
      isSearchVisible: false, // Состояние для поиска
      searchQuery: "", // Поле для ввода поиска
      searchInput: "",
      recipes: [ // Список рецептов
        { id: 1, title: "Лазанья из свиной шейки", price: 650, rating: 4.5, image: "../assets/images/lasagna.png" },
        { id: 2, title: "Тушёная говядина", price: 650, rating: 4.5, image: "../assets/images/beef.png" },
        { id: 3, title: "Сэндвич с яйцом пашот", price: 650, rating: 4.5, image: "../assets/images/sandwich.png" },
      ],
    };
  },
  computed: {
    filteredRecipes() {
      if (this.searchInput.trim() === "") {
        return this.recipes; // Показать все рецепты, если строка поиска пустая
      }
      return this.recipes.filter(recipe =>
        recipe.title.toLowerCase().includes(this.searchInput.toLowerCase())
      );
    },
  },
  methods: {
    toggleSearchBar() {
      this.isSearchVisible = !this.isSearchVisible;
    },
    performSearch() {
      // Обновляем searchInput только при нажатии на "Найти"
      this.searchInput = this.searchQuery;
      console.log("Поиск по запросу:", this.searchInput);
    },
  },
};

</script>

<style scoped>
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
  gap: 5px;
}
.search-bar.active {
  display: flex; /* Показать при активном состоянии */
  
}

.search-input {
  border: 1px solid #ddd;
  border-radius: 4px;
  outline: none;
  width: 300px;
}

.submit-btn {
  padding: 8px 12px;
  border: none;
  background-color: #20b57f;
  color: white;
  border-radius: 4px;
  cursor: pointer;
}

.search-btn {
  background: none;
  border: none;
  color: #12A370;
  cursor: pointer;
  font-size: 20px;
}


.search-btn:hover, .filter-btn:hover {
  color:#0c6445;
  background-color: white;
}

.submit-btn:hover{
  background-color: #0c6445;
}
</style>

