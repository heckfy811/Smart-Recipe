<template>
  <section class="profile">
    <h2>Мой профиль</h2>
    <div class="profile-card">
      <div class="profile-avatar">
        <img :src="user.userphoto" alt="Аватар" />
        <input
          type="file"
          ref="fileInput"
          style="display: none"
          @change="uploadPhoto"
        />
        <button class="btn edit-avatar-btn" @click="choosePhoto">Изменить фото</button>
      </div>
      <div class="profile-info">
        <div class="profile-field">
          <label>Имя пользователя</label>
          <input
            type="text"
            v-model="userEdit.id"
            :readonly="!isEditing"
          />
        </div>
        <div class="profile-field">
          <label>Номер телефона</label>
          <input
            type="text"
            v-model="userEdit.phone"
            :readonly="!isEditing"
          />
        </div>
        <div class="profile-field">
          <label>Адрес электронной почты</label>
          <input
            type="email"
            v-model="userEdit.email"
            :readonly="!isEditing"
          />
        </div>
        <div class="profile-actions">
          <button class="btn edit-profile-btn" @click="toggleEdit">
            {{ isEditing ? "Сохранить" : "Редактировать данные профиля" }}
          </button>
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
        id: "Неизвестный",
        phone: "Неизвестно",
        email: "Неизвестно",
        userphoto: "/assets/images/default-avatar.png",
      },
      userEdit: {}, // Для редактирования данных
      isEditing: false,
    };
  },
  methods: {
    toggleEdit() {
      if (this.isEditing) {
        // Сохранение данных в localStorage
        this.user = { ...this.userEdit };
        localStorage.setItem("user", JSON.stringify(this.user));
        this.isEditing = false;
      } else {
        // Включить режим редактирования
        this.userEdit = { ...this.user };
        this.isEditing = true;
      }
    },
    choosePhoto() {
      // Открытие проводника
      this.$refs.fileInput.click();
    },
    uploadPhoto(event) {
      const file = event.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
          // Сохранение фото в localStorage
          this.user.userphoto = e.target.result; // Base64 URL
          localStorage.setItem("user", JSON.stringify(this.user));
        };
        reader.readAsDataURL(file);
      }
    },
    logout() {
      localStorage.removeItem("isAuthenticated");
      localStorage.removeItem("user");
      this.$router.push("/auth");
    },
  },
  mounted() {
    // Убедимся, что данные синхронизированы с localStorage
    this.user = JSON.parse(localStorage.getItem("user")) || this.user;
  },
};
</script>


<style scoped>
.profile {
  padding: 20px;
  background-color: #f9f9f9;
  display: block;
  justify-content: center;
}

.profile-card {
  display: flex;
  flex-direction: space-between; 
  align-items: center;
  gap: 20px;
  background-color: #fff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  width: 90%; 
  margin: 0 auto; 
}

.profile-avatar img {
  width: 120px;
  height: 120px;
  border-radius: 50%;
}

.profile-avatar .btn {
  margin-top: 10px;
}

.profile-field {
  width: 100%;
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
  display: flex;
  justify-content: space-between; /* Кнопки в одну строку */
  flex-wrap: wrap; /* Позволяет кнопкам переноситься */
}

.profile-actions .btn {
  padding: 10px 15px;
  margin: 5px;
  cursor: pointer;
  text-align: center;
}

.edit-profile-btn {
  color: #ffffff !important; 
  background-color: #12a370;
  border: 1px solid #12a370;
  border-radius: 20px;
  padding: 8px 15px;
  text-decoration: none;
  transition: color 0.3s ease, background-color 0.3s ease;
}

.edit-profile-btn:hover {
  color: #12a370 !important; 
  background-color: #ffffff;
} 

.logout-btn {
  color: #12a370 !important;
  border: 1px solid #12a370;
  border-radius: 20px;
  padding: 8px 15px;
  text-decoration: none;
  transition: color 0.3s ease, background-color 0.3s ease;
}

.logout-btn:hover {
  color: white !important;
  background-color: #12a370; 
}

.edit-avatar-btn {
  color: #ffffff !important; 
  background-color: #12a370;
  border: 1px solid #12a370;
  border-radius: 20px;
  padding: 8px 15px;
  text-decoration: none;
  transition: color 0.3s ease, background-color 0.3s ease;
}

.edit-avatar-btn:hover {
  color: #12a370 !important; 
  background-color: #ffffff;
}

/* Адаптивность */
@media (max-width: 768px) {
  .profile-card {
    flex-direction: column;
    gap: 10px; 
  }

  .profile-avatar,
  .profile-actions {
    align-items: center;
    text-align: center;
  }



  .profile-actions {
    flex-direction: column; 
    align-items: center;
  }

  .profile-actions .btn {
    width: 100%; 
    margin: 5px 0; 
  }
}
</style>