<template>
    <section class="best-recipe">
      <h2>Лучший рецепт</h2>
      <div class="best-recipe-card">
        <img src="@/assets/images/dumplings.png" alt="Домашние пельмени" class="best-recipe-card-img" />
        <div class="additional-images">
          <img src="@/assets/images/dumplings1.png" alt="Шаг 1" />
          <img src="@/assets/images/dumplings2.png" alt="Шаг 2" />
          <img src="@/assets/images/dumplings3.png" alt="Шаг 3" />
        </div>
        <div class="recipe-info">
          <h3>Домашние Пельмени</h3>
          <p>Время готовки 30 минут</p>
          <p>
            Домашние пельмени – это идеальное сочетание тонкого теста и сочной, ароматной
            начинки. Их можно приготовить с разнообразными начинками, от классического
            мясного фарша до необычных вкусов с грибами или овощами. Пельмени не только
            вкусные, но и питательные, их легко готовить для всей семьи. Их можно
            подавать с любыми соусами: сметаной, уксусом или бульоном, что делает их
            универсальным блюдом на любой вкус. Приготовление пельменей своими руками –
            это отличный способ создать домашний уют и порадовать близких.
          </p>
          <router-link to="/recipe/1" class="btn-small">Попробовать рецепт</router-link>
        </div>
      </div>
    </section>
  </template>
  
  <script>
  export default {};
  </script>
  
  <style scoped>
  /* Стили для этого компонента, если они необходимы */
  </style>  