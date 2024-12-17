<template>
  <div class="auth-wrapper">
    <div class="auth-card">
      <h1>Добро пожаловать в SmartRecipe</h1>
      <form @submit.prevent="handleLogin">
        <input
          type="text"
          v-model="loginData.identifier"
          placeholder="Введите ID, номер телефона или email"
          required
        />
        <input
          type="password"
          v-model="loginData.password"
          placeholder="Введите пароль"
          required
        />
        <button type="submit">Войти</button>
      </form>
      <p>
        Нет аккаунта? <router-link to="/register">Зарегистрироваться</router-link>
      </p>
    </div>
  </div>
</template>

<script>
import usersData from "@/assets/testing-data/testAuth.json";

export default {
  name: "AuthComponent",
  data() {
    return {
      loginData: {
        identifier: "",
        password: "",
      },
      users: usersData,
    };
  },
  methods: {
    handleLogin() {
      const user = this.users.find(
        (u) =>
          (u.id === this.loginData.identifier ||
            u.email === this.loginData.identifier ||
            u.phone === this.loginData.identifier) &&
          u.password === this.loginData.password
      );

      if (user) {
        alert(`Добро пожаловать, ${user.id}!`);
        localStorage.setItem("isAuthenticated", "true");
        localStorage.setItem("user", JSON.stringify(user));
        this.$router.push("/");
      } else {
        alert("Неверный ID, email, номер телефона или пароль.");
      }
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
