<template>
  <div>
    <section class="map">
        <h2>Рассчет полной корзины в самом выгодном магазине</h2>
        <div class="stars">
            <h3>"Пятерочка"</h3>
            <h5>4.5</h5>
            <span>★★★★☆</span>
        </div>
        <p>Адрес: улица Академика Анохина, 26к5, 15 мин пешком 🚶‍♂️</p>
        <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2250.533463646628!2d37.46843346992128!3d55.66232302469724!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x46b54d9f2a8eead9%3A0xf7979b37e8853f4f!2z0J_Rj9GC0ZHRgNC-0YfQutCw!5e0!3m2!1sru!2suk!4v1734185876539!5m2!1sru!2suk" width="800" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>    </section>

    <section class="price">
        <h2>Цены на ингредиенты</h2>
        <h3>Ингредиенты для вьетнамской тушёной говядины (Bo Kho):</h3>
        <ul class="ingredients">
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>            
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
            <li><p>Название ингредиента, фирма, граммовка</p> <span>Цена, руб</span></li>
        </ul>

        <h3 class="total-cost"><strong>Общая стоимость:</strong> 1500 руб</h3>
        <router-link to="/near" class="highlight-link">Посмотреть ближайший магазин</router-link>
    </section>
  </div>
</template>

<script>
export default {
  name: "Cheap.vue",
};
</script>

<style scoped>
.price h3 {
  font-size: 18px;
  font-weight: bold;
  margin-bottom: 10px;
}

.ingredients {
  list-style-type: none;
  padding: 0;
  margin: 0;
  border-top: 1px solid #ccc;
}

.ingredients li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #ccc;

}

.ingredients p {
  margin: 0;
  color: #666666c4;
  font-weight: 500;
}

.ingredients span {
  color: #666666c4;
  font-weight: bold;
  font-size: medium;
}

/* Итоговая стоимость */
.total-cost {
  font-size: 20px;
  font-weight: bold;
  margin-top: 20px;
}

.highlight-link {
  display: inline-block;
  margin-top: 10px;
  text-decoration: none;
  color: #12A370;
  font-weight: bold;
}

.highlight-link:hover {
  text-decoration: underline;
  color: #0c6445;
}
</style>