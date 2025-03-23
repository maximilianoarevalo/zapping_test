<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand">Plataforma de streaming</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto">
          <li class="nav-item">
            <router-link to="/" class="nav-link" v-if="!isAuthenticated">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/crear-cuenta" class="nav-link" v-if="!isAuthenticated">Crear Cuenta</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/player" class="nav-link" v-if="isAuthenticated">Player</router-link>
          </li>
        </ul>
        <ul class="navbar-nav">
          <li class="nav-item" v-if="isAuthenticated" id="logoutButton">
            <button class="btn btn-light ms-lg-2" @click="logout">Cerrar sesión</button>
          </li>
        </ul>
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
@import url('https://fonts.googleapis.com/css2?family=Winky+Sans:ital,wght@0,300..900;1,300..900&display=swap');

* {
  font-family: 'Winky Sans', sans-serif;
}

.navbar-brand {
  font-size: 25px;
  font-weight: bold;
}

.active {
  text-decoration: underline;
}

.navbar {
  height: 7vh;
}

.navbar-nav {
  display: flex;
  width: 100%;
}

#logoutButton{
  margin-left: 80%;
}

.ms-lg-2 {
  margin-left: 10px;
}
</style>
