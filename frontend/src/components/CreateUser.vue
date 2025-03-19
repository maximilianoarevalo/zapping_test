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
          <input type="email" v-model="email" class="form-control" id="email" required />
      </div>
      <div class="form-group">
        <label for="password" class="form-label">Contraseña</label>
          <input type="password" v-model="password" class="form-control" id="password1" required />
      </div>

      <div class="form-group">
        <label for="password" class="form-label">Confirmar Contraseña</label>
          <input type="password" v-model="password" class="form-control" id="password2" required />
      </div>
      <div class="button"><button type="submit" class="btn btn-dark">Crear Cuenta</button></div>
      
    </form>
  </div>
  </template>
  
  <script>
  export default {
    name: 'CreateUser',
    data() {
      return {
        name: '',
        email: '',
        password1: '',
        password2: ''
      };
    },
    methods: {
      //TODO: validar doble password
      handleSubmit() {
        //console.log('Cuenta creada:', this.name, this.email, this.password2);
      }
    }
  };
  </script>
  
  <style scoped>
.login-form {
  display: flex;
  justify-content: center;  
  align-items: center;      
  height: 60vh;        
}

.login-form-content {
  background-color: #a0a0a0;  
  padding: 30px;              
  border-radius: 8px;        
  box-shadow: 0 4px 6px rgba(0.2, 0.2, 0.2, 0.2);  
  width: 100%;
  max-width: 400px;         
}

.button{
  display: flex;
  justify-content: center;
}

.btn{
  margin-top: 30px;
  color: white;
}
  </style>
  