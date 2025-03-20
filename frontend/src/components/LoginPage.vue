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
      <div class="button"><button type="submit" class="btn btn-dark">Iniciar Sesión</button></div>
      <div class="register-text"><label class="label">Si no tienes cuenta registrate <router-link to="/crear-cuenta" class="nav-link">acá</router-link></label></div>
    </form>
  </div>
</template>

<script>
export default {
  name: 'LoginPage',
  data() {
    return {
      email: '',
      password: ''
    };
  },
  methods: {
    handleLogin() {
      console.log('Iniciando sesión:', this.email, this.password);
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

.button{
  display: flex;
  justify-content: center;
}

.btn{
  margin-top: 30px;
  color: white;
}

.label{
  display: inline;
}

.nav-link{
  display: inline;
  text-decoration: underline;
}

.register-text{
  text-align: center;
  margin-top: 20px;
}

</style>
