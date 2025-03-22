<template>
  <div class="video-container">
    <video ref="videoPlayer" class="video-player" controls autoplay muted></video>
  </div>
</template>

<script>
import Hls from 'hls.js';

export default {
  name: 'HlsPlayer',
  mounted() {
    this.loadVideo();
    setInterval(this.loadVideo, 10000); // se recarga cada 10 segundos
  },
  methods: {
    loadVideo() {
      const videoElement = this.$refs.videoPlayer;

      // Request para obtener m3u8 actualizado
      fetch('http://localhost:3000/livestream')
        .then(response => response.text())
        .then(m3u8Content => {
          console.log('Contenido del archivo M3U8:', m3u8Content);  // Se muestra archivo con segmentos por consola (para validar)
          
          if (Hls.isSupported()) {
            const hls = new Hls();

            // Evento para verificar carga continua de segmentos
            hls.on(Hls.Events.FRAG_LOADED, (event, data) => {
              console.log('Segmento cargado: ', data);
            });

            hls.loadSource(m3u8Url);  // Se carga archivo
            hls.attachMedia(videoElement);  // Se vincula al video
          }
        })
        .catch(error => {
          console.error('Error al cargar el archivo m3u8:', error);
        });
    }
  }
};
</script>

<style scoped>
.video-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f8f9fa;
}

.video-player {
  width: 100%;
  max-width: 800px;
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}
</style>