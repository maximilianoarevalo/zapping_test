<template>
  <div class="login-form">
    <form class="login-form-content" @submit.prevent="handleLogin">
      <h2>Iniciar Sesión</h2>
      <div class="form-group">
        <label for="email" class="form-label">Email</label>
        <input type="email" v-model="email" class="form-control" id="email" required />
      </div>
      <div class="form-group">
        <label for="password" class="form-label">Contraseña</label>
        <input type="password" v-model="password" class="form-control" id="password" required />
      </div>
      <div v-if="notRegistered" class="not-registered-user">Usuario no registrado</div>
      <div class="button"><button type="submit" class="btn btn-dark">Iniciar Sesión</button></div>
      <div class="register-text"><label class="label">Si no tienes cuenta registrate <router-link to="/crear-cuenta"
            class="nav-link">acá</router-link></label></div>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'LoginPage',
  data() {
    return {
      email: '',
      password: '',
      notRegistered: false
    };
  },
  methods: {
    async handleLogin() {
      try {
        const response = await axios.post("http://localhost:3000/login", {
          email: this.email,
          password: this.password
        });
        console.log("Respuesta del backend:", response.data);
        console.log(response.data.status)
        // Redireccion a player
        if (response.data.status == 200) {
          alert("Inicio de sesión correcto!");
          // Guardar el token en localStorage
          localStorage.setItem("token", response.data.token);
          this.$router.push('/player');
        }
      } catch (error) {
        if (error.response) {
          console.log(error.response)
          console.error("Error en la respuesta del servidor:", error.response.data);
          if (error.response.status == 404) {
            this.notRegistered = true;
          } else if (error.response.status == 400) {
            alert("Contraseña incorrecta!");
            this.notRegistered = false;
          }
          else {
            this.notRegistered = false;
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
  box-shadow: 0 4px 6px rgba(0.2, 0.2, 0.2, 0.2);
  width: 100%;
  max-width: 450px;
}

.button {
  display: flex;
  justify-content: center;
}

.btn {
  margin-top: 30px;
  color: white;
}

.label {
  display: inline;
}

.nav-link {
  display: inline;
  text-decoration: underline;
}

.register-text {
  text-align: center;
  margin-top: 20px;
}

.not-registered-user {
  text-align: center;
  margin-top: 20px;
}
</style>
