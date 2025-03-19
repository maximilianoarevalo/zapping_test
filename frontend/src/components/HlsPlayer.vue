<template>
  <div class="video-container">
    <video ref="videoPlayer" class="video-player" controls></video>
  </div>
</template>

<script>
//import Hls from 'hls.js';

export default {
  name: 'HlsPlayer',
  data() {
    return {
    };
  },
  mounted() { }
  ,
  methods: {
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