<template>
  <div class="login-form">
    <form class="login-form-content" @submit.prevent="handleSubmit">
      <h2>Crear Cuenta</h2>
      <div class="form-group">
        <label for="name" class="form-label">Nombre</label>
        <input type="text" v-model="name" class="form-control" id="name" required />
      </div>
      <div class="form-group">
        <label for="email" class="form-label">Email</label>
        <input type="email" v-model="email" class="form-control" @input="validateEmail" id="email" required />
        <div v-if="emailError" class="error-message">{{ emailError }}</div>
      </div>
      <div class="form-group">
        <label for="password" class="form-label">Contraseña</label>
        <input type="password" v-model="password1" class="form-control" id="password1" @input="validatePasswords"
          required />
      </div>

      <div class="form-group">
        <label for="password" class="form-label">Confirmar Contraseña</label>
        <input type="password" v-model="password2" class="form-control" id="password2" @input="validatePasswords"
          required />
        <p v-if="errorMessage" class="error-message">{{ errorMessage }}</p>
      </div>
      <div v-if="registered" class="registered-user">Ya existe una cuenta para ese correo</div>
      <div class="button"><button type="submit" class="btn btn-dark" :disabled="!buttonEnabled || errorMessage || emailError"
          >Crear Cuenta</button></div>

    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CreateUser',
  data() {
    return {
      name: '',
      email: '',
      password1: '',
      password2: '',
      errorMessage: '',
      buttonEnabled: false,
      registered: false,
      emailError: null
    };
  },
  methods: {
    validateEmail() {
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!this.email.match(emailRegex)) {
        this.emailError = "El email no es válido";
      } else {
        this.emailError = null;
      }
    },
    validatePasswords() {
      if (this.password1 !== this.password2) {
        this.errorMessage = 'Las contraseñas no coinciden';
        this.buttonEnabled = false;
      } else {
        this.errorMessage = '';
        this.buttonEnabled = true;
      }
    },
    async handleSubmit(event) {
      event.preventDefault();
      console.log('Cuenta creada:', this.name, this.email, this.password1, this.password2);
      try {
        const response = await axios.post("http://localhost:3000/create-user", {
          name: this.name,
          email: this.email,
          password: this.password1
        });
        console.log("Respuesta del backend:", response.data);
        // Redireccion a player
        if (response.data.status == 200) {
          this.$router.push('/');
          alert("Inicia sesion con tu nueva cuenta!");
        }
      } catch (error) {
        if (error.response) {
          console.log(error.response)
          console.error("Error en la respuesta del servidor:", error.response.data); //manejar usuario ya registrado
          if (error.response.status == 400) {
            this.registered = true;
          }else{
            this.registered = false;
          }
        } else if (error.request) {
          console.error("No se recibió respuesta del servidor");
        } else {
          console.error("Error al enviar datos:", error.message);
        }
      }
    }
  }
};
</script>

<style scoped>
.login-form {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 50vh;
}

.login-form-content {
  background-color: #a0a0a0;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px;
}

.button {
  display: flex;
  justify-content: center;
}

.btn {
  margin-top: 30px;
  color: white;
}

.error-message {
  margin-top: 20px;
  text-align: center;
}

.registered-user {
  text-align: center;
  margin-top: 20px;
}

.error-message {
  color: red;
  font-size: 0.9em;
  margin-top: 5px;
}
</style>