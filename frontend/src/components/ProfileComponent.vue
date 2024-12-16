<template>
  <div>
    <section class="profile">
            <h2>Мой профиль</h2>
            <div class="profile-card">
                <div class="profile-avatar">
                    <img src="../assets/images/default-avatar.png" alt="Аватар">
                    <button class="btn edit-avatar-btn">Изменить фото</button>
                </div>
                <div class="profile-info">
                    <div class="profile-field">
                        <label>Имя пользователя</label>
                        <input type="text" value="Nickname" readonly>
                    </div>
                    <div class="profile-field">
                        <label>Номер телефона</label>
                        <input type="text" value="Phone number" readonly>
                    </div>
                    <div class="profile-field">
                        <label>Адрес электронной почты</label>
                        <input type="email" value="Email" readonly>
                    </div>
                    <div class="profile-actions">
                        <button class="btn edit-profile-btn">Редактировать данные профиля</button>
                        <button class="btn logout-btn">Выйти из аккаунта</button>
                    </div>
                </div>
            </div>  
    </section>

    <!-- Секция "Избранное" -->
    <section class="recipes">
      <div class="top">
        <h2>Избранное</h2>
        <router-link to="/favorite" class="view-all">Посмотреть все →</router-link>
      </div>
      <div class="recipes-grid">
        <div 
          class="recipe-card" 
          v-for="(recipe, index) in recipes" 
          :key="index">
          <img :src="recipe.image" :alt="recipe.title" />
          <h3>{{ recipe.title }}</h3>
          <p>{{ recipe.price }}Р</p>
          <div class="stars">
            <h5>{{ recipe.rating }}</h5>
            <span>★★★★☆</span>
          </div>
          <router-link :to="'/recipe' + recipe.id" class="btn-small">Посмотреть</router-link>
        </div>
      </div>
    </section>

    <!-- Секция "История" -->
    <section class="recipes">
      <div class="top">
        <h2>История опробованных рецептов</h2>
        <router-link to="/history" class="view-all">Посмотреть все →</router-link>
      </div>
      <div class="recipes-grid">
        <div 
          class="recipe-card" 
          v-for="(recipe, index) in recipes" 
          :key="index">
          <img :src="recipe.image" :alt="recipe.title" />
          <h3>{{ recipe.title }}</h3>
          <p>{{ recipe.price }}Р</p>
          <div class="stars">
            <h5>{{ recipe.rating }}</h5>
            <span>★★★★☆</span>
          </div>
          <router-link :to="'/history' + recipe.id" class="btn-small">Посмотреть</router-link>
        </div>
      </div>
      <h2>Мои планы питания</h2>
    </section>
  </div>
</template>

<script>
export default {
  name: "ProfileComponent.vue",
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
