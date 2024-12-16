<template>
    <div>
        <section class="plans">
            <h2>Мои планы питания</h2>
            <div class="plans-grid">
                <div class="plan-card">
                    <img src="../assets/images/lasagna.png" alt="Лазанья">
                    <h3>Рацион 1</h3>
                    <router-link to="/plan" class="btn-small">Посмотреть</router-link>
                </div>
                <div class="plan-card">
                    <img src="../assets/images/beef.png" alt="Тушёная говядина">
                    <h3>Рацион 2</h3>
                    <router-link to="/plan" class="btn-small">Посмотреть</router-link>
                </div>
                <div class="plan-card">
                    <img src="../assets/images/sandwich.png" alt="Сэндвич">
                    <h3>Рацион 3</h3>
                    <router-link to="/plan" class="btn-small">Посмотреть</router-link>
                </div>
            </div>
        </section>
        <section class="subscription">
            <div class="subscription-text">
                <p>Питайтесь вкусно и полезно</p>
            </div>
        </section>
        <div class="container">
            <div class="block">
                <div class="icon"><img src="../assets/images/clock.svg" alt="Часы" width="100" height="100"></div>
                <div class="card">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</div>
            </div>
            <div class="block">
                <div class="icon"><img src="../assets/images/apple.svg" alt="Яблоко" width="100" height="100"></div>
                <div class="card">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</div>
            </div>
            <div class="block">
                <div class="icon"><img src="../assets/images/wallet.svg" alt="Кошелек" width="100" height="100"></div>
                <div class="card">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</div>
            </div>
            <div class="block">
                <div class="icon"><img src="../assets/images/rest.svg" alt="Столовые приборы" width="100" height="100"></div>
                <div class="card">Lorem ipsum dolor sit amet, consectetur adipiscing elit.</div>
            </div>
        </div>
        <button @click="openModal" id="addPlan" class="add-plan">Добавить рацион</button>
    
        <div id="modal" class="modal" v-if="isModalOpen">
            <div class="modal-content">
                <button @click="closeModal" id="closeModalBtn" class="close-btn">✖</button>
                <h4>Заполните форму, чтобы искусственный интеллект смог подобрать вам подходящий рацион питания.</h4>
                <!-- Форма -->
                <form>
                <label>Калории: <input type="number" placeholder="ккал"></label>
                <label>Белки: <input type="number" placeholder="г"></label>
                <label>Жиры: <input type="number" placeholder="г"></label>
                <label>Углеводы: <input type="number" placeholder="г"></label>
                <label>Предпочитаемая кухня:
                    <select>
                    <option value="">Выберите...</option>
                    <option>Итальянская</option>
                    <option>Мексиканская</option>
                    <option>Японская</option>
                    </select>
                </label>
                <label>Исключить продукты:
                    <select>
                    <option value="">Выберите...</option>
                    <option>Молочные</option>
                    <option>Глютен</option>
                    <option>Орехи</option>
                    </select>
                </label>
                <label>Моя цель:
                    <select>
                    <option value="">Выберите...</option>
                    <option>Похудение</option>
                    <option>Набор массы</option>
                    <option>Поддержание веса</option>
                    </select>
                </label>
                <label>Ваши пожелания:
                    <textarea placeholder="Введите текст..."></textarea>
                </label>
                <button type="submit" class="submit-btn">Сгенерировать рацион</button>
                </form>
            </div>
        </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "PlanListComponent.vue",
    data() {
    return {
      isModalOpen: false, // Управляет видимостью модального окна
    };
  },
  methods: {
    openModal() {
        console.log("Открыть модальное окно");
      this.isModalOpen = true; // Открыть модальное окно
    },
    closeModal() {
      this.isModalOpen = false; // Закрыть модальное окно
    },
  },
};
</script>
  