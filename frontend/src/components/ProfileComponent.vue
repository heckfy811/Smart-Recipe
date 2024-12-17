<template>
  <section class="profile">
    <h2>Мой профиль</h2>
    <div class="profile-card">
      <div class="profile-avatar">
        <img src="../assets/images/default-avatar.png" alt="Аватар" />
        <button class="btn edit-avatar-btn">Изменить фото</button>
      </div>
      <div class="profile-info">
        <div class="profile-field">
          <label>Имя пользователя</label>
          <input type="text" :value="user.name" readonly />
        </div>
        <div class="profile-field">
          <label>Номер телефона</label>
          <input type="text" :value="user.phone" readonly />
        </div>
        <div class="profile-field">
          <label>Адрес электронной почты</label>
          <input type="email" :value="user.email" readonly />
        </div>
        <div class="profile-actions">
          <button class="btn edit-profile-btn">Редактировать данные профиля</button>
          <button class="btn logout-btn" @click="logout">Выйти из аккаунта</button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  name: "ProfileComponent",
  data() {
    return {
      user: JSON.parse(localStorage.getItem("user")) || {
        name: "Неизвестный",
        phone: "Неизвестно",
        email: "Неизвестно",
      },
    };
  },
  methods: {
    logout() {
      localStorage.removeItem("isAuthenticated");
      localStorage.removeItem("user");
      this.$router.push("/auth");
    },
  },
};
</script>

<style scoped>
.profile {
  padding: 20px;
  background-color: #f9f9f9;
}

.profile-card {
  display: flex;
  align-items: center;
  gap: 20px;
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.profile-avatar img {
  width: 120px;
  height: 120px;
  border-radius: 50%;
}

.profile-avatar .btn {
  margin-top: 10px;
}

.profile-field label {
  font-weight: bold;
}

.profile-field input {
  width: 100%;
  padding: 5px;
  margin-top: 5px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background-color: #f5f5f5;
}

.profile-actions {
  margin-top: 20px;
}

.profile-actions .btn {
  padding: 10px 15px;
  margin-right: 10px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.edit-profile-btn {
  background-color: #12a370;
  color: white;
}

.logout-btn {
  background-color: #dc3545;
  color: white;
}
</style>
