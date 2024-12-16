<template>
    <section class="popular-recipes">
      <div class="top">
        <h2>Сейчас популярно</h2>
        <router-link to="/recipes" class="view-all">Посмотреть все →</router-link>
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
  </template>
  
  <script>
  export default {
    data() {
      return {
        recipes: [
          {
            id: 1,
            title: "Лазанья из свиной шейки",
            price: 650,
            rating: 4.5,
            image: "@/assets/images/lasagna.png",
          },
          {
            id: 2,
            title: "Тушёная говядина",
            price: 650,
            rating: 4.5,
            image: "@/assets/images/beef.png",
          },
          {
            id: 3,
            title: "Сэндвич с яйцом пашот",
            price: 650,
            rating: 4.5,
            image: "@/assets/images/sandwich.png",
          },
        ],
      };
    },
  };
  </script>
  
  <style scoped>
  /* Перенесите стили сюда */
  </style>  