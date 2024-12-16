<template>
  <footer>
    <div class="footer-content">
      <div class="footer-logo">
        <h3>Smart.Recipe</h3>
        <p>Готовьте вкусно и недорого</p>
      </div>
      <div class="footer-links">
        <ul>
          <li><router-link to="/about">О нас</router-link></li>
          <li><a href="#">Контакты</a></li>
          <li><a href="#">Политика конфиденциальности</a></li>
          <li><a href="#">Условия использования</a></li>
        </ul>
      </div>
      <div class="footer-social">
        <p>Контакты</p>
        <div class="social-icons">
          <a href="#"><img src="@/assets/images/instagram.png" alt="Instagram" /></a>
          <a href="#"><img src="@/assets/images/facebook.png" alt="Facebook" /></a>
          <a href="#"><img src="@/assets/images/twitter.png" alt="Twitter" /></a>
        </div>
        <p>2024 @smartrecipe.com</p>
      </div>
    </div>
    <p>2024 © SmartRecipe</p>
  </footer>
</template>

<script>
export default {};
</script>

<style scoped>
footer {
  background-color: #12A370;
  padding: 40px;
  color: white;
  margin-top: 50px;
}

.footer-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  flex-wrap: wrap; /* Позволяет перенос на следующую строку */
  gap: 20px; /* Отступы между блоками */
}

.footer-logo {
  flex: 1 1 30%; /* Логотип занимает до 30% ширины */
  min-width: 200px; /* Минимальная ширина для логотипа */
}

.footer-logo h3 {
  margin-bottom: 10px;
}

.footer-links {
  flex: 1 1 30%; /* Ссылки занимают до 30% ширины */
  min-width: 200px; /* Минимальная ширина для ссылок */
}

.footer-links ul {
  list-style: none;
  padding: 0;
}

.footer-links a {
  color: white;
  text-decoration: none;
  display: block;
  margin-bottom: 5px;
}

.footer-social {
  flex: 1 1 30%; /* Социальные ссылки занимают до 30% ширины */
  min-width: 200px; /* Минимальная ширина для соцсетей */
}

.footer-social .social-icons {
  display: flex;
  gap: 10px;
  margin-bottom: 10px;
}

.footer-social img {
  width: 30px; /* Размер иконок уменьшен для мобильных */
  height: 30px;
}

.footer-social p {
  font-size: 1em;
  margin-bottom: 10px;
}

footer p {
  text-align: center;
  margin-top: 20px;
  margin-bottom: 30px;
}

/* Адаптивность */
@media (max-width: 768px) {
  .footer-content {
    flex-direction: column; /* Все элементы располагаются вертикально */
    text-align: center; /* Центрируем текст */
    gap: 30px; /* Увеличиваем отступы между секциями */
  }

  .footer-logo,
  .footer-links,
  .footer-social {
    flex: none;
    width: 100%; /* Каждая секция занимает всю ширину */
    min-width: auto; /* Убираем минимальную ширину */
  }

  .footer-social .social-icons {
    justify-content: center; /* Центрируем иконки */
  }

  footer p {
    font-size: 0.9em; /* Уменьшаем шрифт для мобильных */
  }
}
</style>