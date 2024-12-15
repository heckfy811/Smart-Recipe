<template>
  <section class="hero">
    <div class="hero-content">
      <h1>Вкусные рецепты & Выгодные покупки</h1>
      <p>Выбирайте, чем себя побаловать, а мы поможем сделать с наибольшей выгодой и удобством</p>
      <div class="hero-buttons">
        <router-link to="/recipes" class="btn">Рецепты</router-link>
        <router-link to="/subscription" class="btn-outline">Подписка</router-link>
      </div>
    </div>
    <div class="hero-images">
      <img src="@/assets/images/img1.png" alt="Изображение 1">
      <img src="@/assets/images/img2.png" alt="Изображение 2">
      <img src="@/assets/images/img3.png" alt="Изображение 3">
    </div>
  </section>
</template>

<script>
export default {};
</script>

<style scoped>
.hero {
display: flex;
flex-direction: column;
align-items: center;
}

.hero-content {
text-align: center;
margin-bottom: 20px;
}

.hero-buttons {
display: flex;
gap: 10px;
}

.hero-images {
display: flex;
overflow-x: hidden;
}

.hero-images img {
flex: 0 0 auto;
width: auto;
height: 100%;
object-fit: cover;
}

@media (min-width: 768px) {
.hero {
  flex-direction: row;
  justify-content: space-between;
}

.hero-content {
  text-align: left;
  margin-bottom: 0;
}

.hero-images {
  flex-grow: 1;
}
}
</style>
