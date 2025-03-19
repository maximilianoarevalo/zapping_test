<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Plataforma de streaming</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item active">
            <router-link to="/" class="nav-link">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/crear-cuenta" class="nav-link">Crear Cuenta</router-link>
          </li>
          <li class="nav-item">
            <router-link to="/player" class="nav-link">Player</router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
  <router-view></router-view>
</template>

<script>

export default {
  name: 'App',
  components: {
  }
}
</script>

<style>
.navbar-brand{
  font-size: 25px;
  font-weight: bold;
}
.active{
  text-decoration: underline;
}

</style>
