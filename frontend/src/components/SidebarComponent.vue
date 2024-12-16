<template>
  <div>
    <div id="overlay" class="overlay" @click="$emit('close-sidebar')"></div>
    <aside v-if="isVisible" class="filter-sidebar">
        <div class="filter-header">
          <h2>Фильтры</h2>
          <button id="close-btn" class="close-btn" @click="$emit('close-sidebar')">✕</button>        
        </div>
        <div class="filter-content">
          <!-- Цена -->
          <div class="filter-section">
            <p>Цена, руб.</p>
            <div class="input-group">
              <input type="number" placeholder="От">
              <input type="number" placeholder="До">
            </div>
          </div>

            <!-- Время приготовления -->
            <div class="filter-section">
                <p>Время приготовления, мин.</p>
                <div class="input-group">
                    <input type="number" placeholder="От">
                    <input type="number" placeholder="До">
                </div>
            </div>

            <!-- Категории -->
            <div class="filter-section">
                <p>Категория</p>
                <ul>
                    <li><input type="checkbox" id="category1"> <label for="category1">Первые блюда</label></li>
                    <li><input type="checkbox" id="category2"> <label for="category2">Вторые блюда</label></li>
                    <li><input type="checkbox" id="category3"> <label for="category3">Десерты</label></li>
                    <li><input type="checkbox" id="category3"> <label for="category3">Закуски</label></li>
                    <li><input type="checkbox" id="category3"> <label for="category3">Напитки</label></li>
                </ul>
            </div>

            <!-- Кухни -->
            <div class="filter-section">
                <p>Кухня</p>
                <ul>
                    <li><input type="checkbox" id="cuisine1"> <label for="cuisine1">Итальянская</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Японская</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Русская</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Мексиканская</label></li>
                </ul>
            </div>

            <div class="filter-section">
                <p>Способ приготовления</p>
                <ul>
                    <li><input type="checkbox" id="cuisine1"> <label for="cuisine1">Варка</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Жарка</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Гриль</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">Запекание</label></li>
                </ul>
            </div>

            <div class="filter-section">
                <p>Калорийность</p>
                <ul>
                    <li><input type="checkbox" id="cuisine1"> <label for="cuisine1">До 300 ккал</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">От 301 до 500 ккал</label></li>
                    <li><input type="checkbox" id="cuisine2"> <label for="cuisine2">От 501 ккал</label></li>
                </ul>
            </div>

            <!-- Применить -->
            <button class="apply-btn" @click="closeSidebar">Применить</button>
          </div>
    </aside>
  </div>
</template>

<script>
export default {
  name: "SidebarComponent.vue",
  props: {
    isVisible: {
      type: Boolean,
      required: true,
    },
  },
};
</script>

