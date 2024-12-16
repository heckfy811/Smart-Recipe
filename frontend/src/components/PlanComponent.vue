<template>
  <div>
    <section class="ration">
        <h1>Рацион 1</h1>
        <!-- Блок с калорийностью -->
        <div class="nutrition-info">
            <div>Калории <br><strong>1800 кк</strong></div>
            <div>Белки <br><strong>85 г</strong></div>
            <div>Жиры <br><strong>44 г</strong></div>
            <div>Углеводы <br><strong>119 г</strong></div>
        </div>

        <!-- Завтрак -->
        <div class="meal-block">
            <div class="recipe-top">
                <h2>Завтрак</h2>
                <p>КБЖУ: 400, 30, 11, 55</p>
            </div>
            <div class="recipes-grid">
                <div class="recipe-card">
                    <img src="../assets/images/lasagna.png" alt="Лазанья">
                    <h3>Блюдо 1</h3>         
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
                <div class="recipe-card">
                    <img src="../assets/images/beef.png" alt="Тушёная говядина">
                    <h3>Блюдо 2</h3>
               
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
            </div>
        </div>

        <!-- Обед -->
        <div class="meal-block">
            <div class="recipe-top">
                <h2>Обед</h2>
                <p>КБЖУ: 800, 30, 11, 55</p>
            </div>
            <div class="recipes-grid">
                <div class="recipe-card">
                    <img src="../assets/images/lasagna.png" alt="Лазанья">
                    <h3>Блюдо 1</h3>         
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
                <div class="recipe-card">
                    <img src="../assets/images/beef.png" alt="Тушёная говядина">
                    <h3>Блюдо 2</h3>
               
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
                <div class="recipe-card">
                    <img src="../assets/images/sandwich.png" alt="Сэндвич">
                    <h3>Блюдо 3</h3>
               
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
            </div>
        </div>

        <!-- Ужин -->
        <div class="meal-block">
            <div class="recipe-top">
                <h2>Ужин</h2>
                <p>КБЖУ: 500, 30, 11, 55</p>
            </div>
            <div class="recipes-grid">
                <div class="recipe-card">
                    <img src="../assets/images/lasagna.png" alt="Лазанья">
                    <h3>Блюдо 1</h3>         
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
                <div class="recipe-card">
                    <img src="../assets/images/beef.png" alt="Тушёная говядина">
                    <h3>Блюдо 2</h3>
               
                    <router-link to="/recipe" class="btn-small">Посмотреть</router-link>
                  </div>
            </div>
        </div>

        <!-- Пагинация -->
        <div class="pagination">
            <a href="#">← Предыдущий рацион</a>
            <a href="#">Следующий рацион →</a>
        </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "PlanComponent.vue",
};
</script>
