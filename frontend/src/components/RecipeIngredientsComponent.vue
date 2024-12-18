<template>
    <div v-if="recipe" class="ingredients">
      <h2>Рецепт и ингредиенты</h2>
      <br />
      <p>Ингредиенты:</p>
      <ul>
        <li v-for="(ingredient, index) in recipe.ingredients" :key="index">
          {{ ingredient }}
        </li>
      </ul>
      <br />
      <p>Рецепт:</p>
      <ol>
        <li v-for="(step, index) in recipe.steps" :key="index">
          {{ step }}
        </li>
      </ol>
    </div>
  </template>
  
  <script>
  export default {
    name: "RecipeIngredientsComponent",
    props: ["recipe"],
  };
  </script>
  
  <style scoped>
  .ingredients {
    margin: 20px;
    padding: 20px;
    background: #f9f9f9;
    border-radius: 10px;
    animation: slideFromLeft 1s ease-out;
  }
  
  .ingredients h2 {
    margin-bottom: 10px;
  }

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
  </style>
  