<template>
    <header>
      <nav class="navbar">
        <router-link to="/" class="logo">
          <div class="smart">Smart</div>.Recipe
        </router-link>
        <ul class="nav-links">
          <li><router-link to="/" class="head click">Главная</router-link></li>
          <li><router-link to="/recipes" class="head">Рецепты</router-link></li>
          <li><router-link to="/plans" class="head">Планы питания</router-link></li>
          <li><router-link to="/about" class="head">О нас</router-link></li>
          <li><router-link to="/profile" class="profile-btn">Профиль</router-link></li>
        </ul>
      </nav>
    </header>
  </template>
  
  <script>
  export default {};
  </script>
  
  <style scoped>
  /* Добавьте любые специфические стили, если необходимо */
  </style>
  