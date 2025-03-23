<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Plataforma de streaming</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse d-flex justify-content-between" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item active">
            <router-link to="/" class="nav-link" v-if="!isAuthenticated">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/crear-cuenta" class="nav-link" v-if="!isAuthenticated">Crear Cuenta</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/player" class="nav-link" v-if="isAuthenticated">Player</router-link>
          </li>
        </ul>
        <!-- Boton logout -->
        <div v-if="isAuthenticated">
          <button class="btn btn-light" @click="logout">Cerrar sesión</button>
        </div>
      </div>


    </div>
  </nav>
  <div class="content">
    <router-view></router-view>
  </div>
</template>

<script>

export default {
  name: 'App',
  computed: {
    isAuthenticated() {
      return !!localStorage.getItem("token");
    }
  },
  methods: {
    logout() {
      localStorage.removeItem("token");
      this.isAuthenticated = false;
      this.$router.push("/");
      window.location.reload();
    }
  }
}
</script>

<style>
.navbar-brand {
  font-size: 25px;
  font-weight: bold;
}

.active {
  text-decoration: underline;
}

.navbar {
  height: 60px;
}

.content {
  margin-top: 70px;
  padding: 20px;
}

.navbar-nav {
  display: flex;
  width: 90%;
}

.logout-item {
  margin-left: auto;
}
</style>
