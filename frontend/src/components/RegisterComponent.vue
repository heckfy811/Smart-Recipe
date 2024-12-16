<template>
    <div class="auth-wrapper">
      <div class="auth-card">
        <h1>Регистрация в SmartRecipe</h1>
        <form @submit.prevent="handleRegister">
          <input
            type="text"
            v-model="formData.id"
            placeholder="Введите ID"
            required
          />
          <input
            type="email"
            v-model="formData.email"
            placeholder="Введите email"
            required
          />
          <input
            type="password"
            v-model="formData.password"
            placeholder="Введите пароль"
            required
          />
          <input
            type="password"
            v-model="formData.passwordRepeat"
            placeholder="Повторите пароль"
            required
          />
          <button type="submit">Зарегистрироваться</button>
        </form>
        <p>
          Уже есть аккаунт? <router-link to="/auth">Войти</router-link>
        </p>
      </div>
    </div>
</template>
  
<script>
export default {
  name: "RegisterComponent",
  data() {
    return {
      formData: {
        id: "",
        phoneNumber: "",
        email: "",
        password: "",
        passwordRepeat: "",
      },
      phoneError: "",
    };
  },
  methods: {
    handleRegister() {
      if (this.formData.password !== this.formData.passwordRepeat) {
        alert("Пароли не совпадают!");
        return;
      }
      // Имитация успешной регистрации
      alert("Регистрация успешна!");
      this.$router.push("/"); // Перенаправление на главную страницу
    },
  },
};
</script>
  
<style scoped>
    * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    }
    body {
    font-family: Arial, sans-serif;
    background: url('/src/assets/images/auth-back.jpg') no-repeat center center fixed;
    background-size: 135%;
    margin: 0;
    line-height: 1.6;
    position: relative;
    }
    body::before {
    content: "";
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    z-index: 1;
    }
    a {
    text-decoration: none;
    color: #28a745;
    }
    .auth-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        min-height: 100vh;
        position: relative;
        z-index: 2;
    }
    .auth-card {
        background: #fff;
        color: #333;
        padding: 40px;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        text-align: center;
        max-width: 400px;
        width: 90%;
    }
    .auth-card h1 {
        margin-bottom: 20px;
    }
    .auth-card input {
        width: 100%;
        padding: 10px;
        margin-bottom: 20px;
        border: 1px solid #ddd;
        border-radius: 4px;
    }
    .auth-card button {
        width: 100%;
        padding: 10px;
        background-color: #28a745;
        border: none;
        color: #fff;
        border-radius: 4px;
        cursor: pointer;
        font-size: 16px;
    }
    .auth-card button:hover {
        background-color: #218838;
    }
</style>
  