<template>
  <div>
    <section class="map">
        <h2>Рассчет полной корзины в ближайшем магазине от вас</h2>
        <div class="stars">
            <h3>"Перекрёсток"</h3>
            <h5>4.5</h5>
            <span>★★★★☆</span>
        </div>
        <p>Адрес: пр. Вернадского, 86, 7 мин пешком 🚶‍♂️</p>
        <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d2250.558126720553!2d37.47671419999999!3d55.661894100000005!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x46b54d9feb106343%3A0x4a859012783558cf!2z0L_RgC4g0JLQtdGA0L3QsNC00YHQutC-0LPQviwgODYsINCc0L7RgdC60LLQsCwg0KDQvtGB0YHQuNGPLCAxMTk1NzE!5e0!3m2!1sru!2suk!4v1734185165508!5m2!1sru!2suk" width="800" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>        
    </section>

    
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
        <router-link to="/cheap" class="highlight-link">Посмотреть самый выгодный магазин</router-link>
    </section>
  </div>
</template>

<script>
export default {
  name: "NearComponent.vue",
};
</script>
